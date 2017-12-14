package libqiniu_test

import (
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	. "gitlab.qiniu.io/pingan/libqiniu"
	"gitlab.qiniu.io/pingan/libqiniu/mock/mock_http"
	ctx "golang.org/x/net/context"
)

var _ = Describe("Upload", func() {
	var (
		ctrl             *gomock.Controller
		token            Uptoken
		uploader         FormUploader
		mockRoundTripper *mock_http.MockRoundTripper
		fakeResponse     *http.Response
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	BeforeEach(func() {
		mockRoundTripper = mock_http.NewMockRoundTripper(ctrl)
		postman := form_postman.NewPostman(&http.Client{
			Transport: mockRoundTripper,
		})
		uploader = NewFormUploader(postman, zone)
	})
	BeforeEach(func() {
		var err error
		deadline := time.Unix(1504448207, 0)
		putPolicy := NewPutPolicyForBucket("testbucket").SetDeadline(deadline)
		token, err = aksk.Uptoken(putPolicy)
		Expect(err).NotTo(HaveOccurred())
	})
	BeforeEach(func() {
		fakeResponse = new(http.Response)
		fakeResponse.Status = "200 OK"
		fakeResponse.StatusCode = 200
		fakeResponse.Proto = "HTTP/1.1"
		fakeResponse.ProtoMajor = 1
		fakeResponse.ProtoMinor = 1
		fakeResponse.Header = make(http.Header)
		fakeResponse.Header.Add("Content-Type", "application/json")
		fakeResponse.ContentLength = int64(2)
		fakeResponse.Body = ioutil.NopCloser(strings.NewReader("{}"))
	})

	Context(".UploadReader", func() {
		It("should upload the reader content", func() {
			const filename = "hw.txt"
			const content = "hello world"
			fakeResponse.Body = ioutil.NopCloser(strings.NewReader(`{"key":"hw.txt"}`))
			fakeResponse.ContentLength = int64(len(`{"key":"hw.txt"}`))
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.Header.Get("Authorization")).To(BeEmpty())
				mediaType, params, err := mime.ParseMediaType(request.Header.Get("Content-Type"))
				Expect(err).NotTo(HaveOccurred())
				Expect(mediaType).To(Equal("multipart/form-data"))

				reader := multipart.NewReader(request.Body, params["boundary"])
				part, err := reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				Expect(part.FormName()).To(Equal("token"))
				token, err := ioutil.ReadAll(part)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(token)).To(Equal("yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:6HhGjzZOzt6j4duX-j7SrowhPZE=:eyJzY29wZSI6InRlc3RidWNrZXQiLCJkZWFkbGluZSI6MTUwNDQ0ODIwN30="))
				Expect(part.Close()).NotTo(HaveOccurred())
				part, err = reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				Expect(part.FormName()).To(Equal("key"))
				key, err := ioutil.ReadAll(part)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(key)).To(Equal(filename))
				Expect(part.Close()).NotTo(HaveOccurred())
				part, err = reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				Expect(part.FormName()).To(Equal("x:author"))
				author, err := ioutil.ReadAll(part)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(author)).To(Equal("Rong Zhou"))
				Expect(part.Close()).NotTo(HaveOccurred())
				part, err = reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				Expect(part.FormName()).To(Equal("file"))
				Expect(part.FileName()).To(Equal(filename))
				file, err := ioutil.ReadAll(part)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(file)).To(Equal(content))
				Expect(part.Close()).NotTo(HaveOccurred())
				part, err = reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				Expect(part.FormName()).To(Equal("crc32"))
				hash, err := ioutil.ReadAll(part)
				Expect(err).NotTo(HaveOccurred())
				Expect(hash).To(Equal([]byte(fmt.Sprintf("%010d", crc32.ChecksumIEEE([]byte(content))))))
				Expect(part.Close()).NotTo(HaveOccurred())
				_, err = reader.NextPart()
				Expect(err).To(Equal(io.EOF))
			}).Return(fakeResponse, nil)
			reader := strings.NewReader(content)
			ret, err := uploader.UploadReader(ctx.Background(), reader, int64(reader.Len()), filename, token, Params{
				Key:        filename,
				MIMEType:   "text/plain",
				XVariables: map[string]string{"x:author": "Rong Zhou"},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(ret.Key).To(Equal(filename))
		})
	})
})
