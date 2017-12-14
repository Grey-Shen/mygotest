package libqiniu

import (
	"fmt"
	"hash/crc32"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"strings"
	"sync"

	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	ctx "golang.org/x/net/context"
	buffer "gopkg.in/djherbis/buffer.v1"
	nio "gopkg.in/djherbis/nio.v2"
)

type (
	Params struct {
		Key        string
		MIMEType   string
		XVariables map[string]string
	}

	UploadRet struct {
		Hash         string `json:"hash"`
		PersistentID string `json:"persistentId,omitempty"`
		Key          string `json:"key"`
		Error        string `json:"error,omitempty"`
	}

	FormUploader struct {
		*Zone
		postman form_postman.Postman
	}
)

var NoParams = Params{}

func NewFormUploader(postman form_postman.Postman, zone *Zone) FormUploader {
	return FormUploader{Zone: zone, postman: postman}
}

func (uploader FormUploader) Upload(context ctx.Context, file io.ReadSeeker, filename string, uptoken Uptoken, params Params) (*UploadRet, error) {
	if size, err := file.Seek(0, os.SEEK_END); err != nil {
		return nil, err
	} else if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	} else {
		return uploader.UploadReader(context, file, size, filename, uptoken, params)
	}
}

func (uploader FormUploader) UploadFile(context ctx.Context, filepath string, uptoken Uptoken, params Params) (*UploadRet, error) {
	if file, err := os.Open(filepath); err != nil {
		return nil, err
	} else {
		defer file.Close()
		if fileInfo, err := file.Stat(); err != nil {
			return nil, err
		} else {
			return uploader.UploadReader(context, file, fileInfo.Size(), path.Base(filepath), uptoken, params)
		}
	}
}

func (uploader FormUploader) UploadReader(context ctx.Context, reader io.Reader, size int64, filename string, uptoken Uptoken, params Params) (*UploadRet, error) {
	pipeR, pipeW := nio.Pipe(buffer.New(4 * (1 << 20)))

	doneChan := make(chan struct{})
	defer close(doneChan)
	retChan := make(chan *UploadRet)
	defer close(retChan)
	errChan := make(chan error, 2)
	defer close(errChan)

	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()

	writer := multipart.NewWriter(pipeW)
	go func(writer *multipart.Writer, doneChan chan<- struct{}, errChan chan<- error) {
		defer wg.Done()
		defer pipeW.Close()

		var err error
		if err = writeFields(writer, uptoken, params); err != nil {
			errChan <- err
			return
		}

		hasher := crc32.NewIEEE()
		reader = io.TeeReader(reader, hasher)

		if err = writeFormFile(writer, filename, params.MIMEType); err != nil {
			errChan <- err
			return
		}
		if _, err = io.CopyN(pipeW, reader, size); err != nil {
			errChan <- err
			return
		}
		if err = writer.WriteField("crc32", fmt.Sprintf("%010d", hasher.Sum32())); err != nil {
			errChan <- err
			return
		}
		if err = writer.Close(); err != nil {
			errChan <- err
			return
		}
		doneChan <- struct{}{}
	}(writer, doneChan, errChan)

	go func(context ctx.Context, contentType string, retChan chan<- *UploadRet, errChan chan<- error) {
		defer wg.Done()
		defer pipeR.Close()

		request, err := http.NewRequest("POST", uploader.UpHosts[0], pipeR)
		if err != nil {
			errChan <- err
			return
		}
		if context == nil {
			context = ctx.Background()
		}
		context = ctx.WithValue(context, "donot_sign_request", true)
		request = request.WithContext(context)

		request.Header.Set("Content-Type", contentType)
		request.Header.Set("Accept", "application/json")

		var ret UploadRet
		if _, err := uploader.postman.SendRequest(request, &ret); err != nil {
			errChan <- err
		} else {
			retChan <- &ret
		}
	}(context, writer.FormDataContentType(), retChan, errChan)

	select {
	case <-doneChan:
	case err := <-errChan:
		return nil, err
	}

	select {
	case ret := <-retChan:
		return ret, nil
	case err := <-errChan:
		return nil, err
	}

}

func writeFields(writer *multipart.Writer, uptoken Uptoken, params Params) error {
	var err error
	if err = writer.WriteField("token", uptoken.String()); err != nil {
		return err
	}
	if params.Key != "" {
		if err = writer.WriteField("key", params.Key); err != nil {
			return err
		}
	}
	if params.XVariables != nil {
		for k, v := range params.XVariables {
			if strings.HasPrefix(k, "x:") && v != "" {
				if err = writer.WriteField(k, v); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func writeFormFile(writer *multipart.Writer, filename, mimeType string) error {
	part := make(textproto.MIMEHeader)
	part.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, escapeQuotes(filename)))
	if mimeType != "" {
		part.Set("Content-Type", mimeType)
	}
	_, err := writer.CreatePart(part)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
