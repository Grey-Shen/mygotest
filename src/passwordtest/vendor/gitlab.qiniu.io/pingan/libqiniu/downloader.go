package libqiniu

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	ctx "golang.org/x/net/context"
)

type (
	DownloadError struct {
		StatusCode   int
		Header       http.Header
		ResponseBody []byte
	}
	Downloader struct {
		client *http.Client
	}
)

func (err DownloadError) Error() string {
	return fmt.Sprintf("Failed Response: Code: %d, Body: %s", err.StatusCode, err.ResponseBody)
}

func NewDownloader(client *http.Client) Downloader {
	return Downloader{client: client}
}

func (downloader Downloader) Exist(context ctx.Context, url string) (bool, error) {
	var (
		request  *http.Request
		response *http.Response
		body     []byte
		err      error
	)
	if request, err = http.NewRequest("HEAD", url, http.NoBody); err != nil {
		return false, err
	} else {
		if context != nil {
			request = request.WithContext(context)
		}
		if response, err = downloader.client.Do(request); err != nil {
			return false, err
		} else {
			defer response.Body.Close()
			switch response.StatusCode {
			case http.StatusOK:
				return true, nil
			case http.StatusNotFound:
				return false, nil
			case http.StatusBadRequest:
				return false, os.ErrInvalid
			case http.StatusUnauthorized:
				return false, os.ErrPermission
			default:
				if body, err = ioutil.ReadAll(response.Body); err != nil {
					return false, err
				} else {
					return false, DownloadError{
						StatusCode:   response.StatusCode,
						Header:       response.Header,
						ResponseBody: body,
					}
				}
			}
		}
	}
}

func (downloader Downloader) Download(context ctx.Context, url string) (io.ReadCloser, error) {
	var (
		request  *http.Request
		response *http.Response
		body     []byte
		err      error
	)
	if request, err = http.NewRequest("GET", url, http.NoBody); err != nil {
		return nil, err
	} else {
		if context != nil {
			request = request.WithContext(context)
		}
		if response, err = downloader.client.Do(request); err != nil {
			return nil, err
		} else {
			switch response.StatusCode {
			case http.StatusOK:
				return response.Body, nil
			case http.StatusBadRequest:
				response.Body.Close()
				return nil, os.ErrInvalid
			case http.StatusUnauthorized:
				response.Body.Close()
				return nil, os.ErrPermission
			case http.StatusNotFound:
				response.Body.Close()
				return nil, os.ErrNotExist
			default:
				body, err = ioutil.ReadAll(response.Body)
				response.Body.Close()
				if err != nil {
					return nil, err
				} else {
					return nil, DownloadError{
						StatusCode:   response.StatusCode,
						Header:       response.Header,
						ResponseBody: body,
					}
				}
			}
		}
	}
}

func (downloader Downloader) DownloadTo(context ctx.Context, url string, writter io.Writer) (int64, error) {
	if body, err := downloader.Download(context, url); err != nil {
		return 0, err
	} else {
		return io.Copy(writter, body)
	}
}

func (downloader Downloader) DownloadToFile(context ctx.Context, url string, path string) error {
	if file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
		return err
	} else {
		defer file.Close()
		_, err = downloader.DownloadTo(context, url, file)
		return err
	}
}
