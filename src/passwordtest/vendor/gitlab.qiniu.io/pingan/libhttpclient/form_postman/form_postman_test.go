package form_postman_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	"gitlab.qiniu.io/pingan/libhttpclient/form_postman/mock_http"
	ctx "golang.org/x/net/context"
)

var _ = Describe("FormPostman", func() {
	var ctrl *gomock.Controller
	var mockRoundTripper *mock_http.MockRoundTripper
	var postman Postman
	var fakeResponse *http.Response

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})
	BeforeEach(func() {
		mockRoundTripper = mock_http.NewMockRoundTripper(ctrl)
		postman = NewPostman(&http.Client{
			Transport: mockRoundTripper,
		})
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
		const body = `[{"code":200,"data":{"fsize":15508,"hash":"FlsRbe0-0Cr3S6kqkGKl_-WDfbxF","mimeType":"text/css"` +
			`,"putTime":15040850589759659,"type":0}},{"code":200,"data":{"fsize":158215,"hash":"FsALPkde4eIWUOod0VIYbdUlxx92"` +
			`,"mimeType":"text/javascript","putTime":15040850590777631,"type":0}}]`
		fakeResponse.ContentLength = int64(len(body))
		fakeResponse.Body = ioutil.NopCloser(strings.NewReader(body))
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	type BatchOpRet struct {
		Code int `json:"code,omitempty"`
		Data struct {
			Hash     string `json:"hash"`
			Fsize    int64  `json:"fsize"`
			PutTime  int64  `json:"putTime"`
			MimeType string `json:"mimeType"`
			Type     int    `json:"type"`
			Error    string `json:"error"`
		} `json:"data,omitempty"`
	}

	Context(".SendRequest", func() {
		It("should send the request", func() {
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				buf := new(bytes.Buffer)
				_, err := io.Copy(buf, request.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(buf.String()).To(Equal(`op=%2Fstat%2FdGVzdDprZXkx&op=%2Fstat%2FdGVzdDprZXky`))
			}).Return(fakeResponse, nil)

			request, err := http.NewRequest("POST", "http://rs.qiniu.com/batch", strings.NewReader(`op=%2Fstat%2FdGVzdDprZXkx&op=%2Fstat%2FdGVzdDprZXky`))
			Expect(err).NotTo(HaveOccurred())
			ops := make([]BatchOpRet, 0, 2)
			_, err = postman.SendRequest(request, &ops)
			Expect(err).NotTo(HaveOccurred())
			Expect(ops).To(HaveLen(2))
			Expect(ops[0].Code).To(Equal(200))
			Expect(ops[0].Data.Hash).To(Equal("FlsRbe0-0Cr3S6kqkGKl_-WDfbxF"))
			Expect(ops[0].Data.Fsize).To(Equal(int64(15508)))
			Expect(ops[0].Data.PutTime).To(Equal(int64(15040850589759659)))
			Expect(ops[0].Data.MimeType).To(Equal("text/css"))
			Expect(ops[0].Data.Type).To(Equal(0))
			Expect(ops[1].Code).To(Equal(200))
			Expect(ops[1].Data.Hash).To(Equal("FsALPkde4eIWUOod0VIYbdUlxx92"))
			Expect(ops[1].Data.Fsize).To(Equal(int64(158215)))
			Expect(ops[1].Data.PutTime).To(Equal(int64(15040850590777631)))
			Expect(ops[1].Data.MimeType).To(Equal("text/javascript"))
			Expect(ops[1].Data.Type).To(Equal(0))
		})
	})

	Context(".Send", func() {
		It("should send the request", func() {
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				contextValue := request.Context().Value("key")
				Expect(contextValue).NotTo(BeNil())
				Expect(contextValue.(string)).To(Equal("value"))
				Expect(request.Header.Get("Content-Type")).To(Equal("application/x-www-form-urlencoded"))
				Expect(request.Header.Get("Accept")).To(Equal("application/json"))
				Expect(request.Header.Get("Authorization")).To(Equal("auth"))
				buf := new(bytes.Buffer)
				_, err := io.Copy(buf, request.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(buf.String()).To(Equal(`op=%2Fstat%2FdGVzdDprZXkx&op=%2Fstat%2FdGVzdDprZXky`))
			}).Return(fakeResponse, nil)

			ops := make([]BatchOpRet, 0, 2)
			_, err := postman.Send(
				ctx.WithValue(ctx.Background(), "key", "value"),
				"POST", "http://rs.qiniu.com/batch", "auth",
				strings.NewReader(`op=%2Fstat%2FdGVzdDprZXkx&op=%2Fstat%2FdGVzdDprZXky`),
				&ops)
			Expect(err).NotTo(HaveOccurred())
			Expect(ops).To(HaveLen(2))
			Expect(ops[0].Code).To(Equal(200))
			Expect(ops[0].Data.Hash).To(Equal("FlsRbe0-0Cr3S6kqkGKl_-WDfbxF"))
			Expect(ops[0].Data.Fsize).To(Equal(int64(15508)))
			Expect(ops[0].Data.PutTime).To(Equal(int64(15040850589759659)))
			Expect(ops[0].Data.MimeType).To(Equal("text/css"))
			Expect(ops[0].Data.Type).To(Equal(0))
			Expect(ops[1].Code).To(Equal(200))
			Expect(ops[1].Data.Hash).To(Equal("FsALPkde4eIWUOod0VIYbdUlxx92"))
			Expect(ops[1].Data.Fsize).To(Equal(int64(158215)))
			Expect(ops[1].Data.PutTime).To(Equal(int64(15040850590777631)))
			Expect(ops[1].Data.MimeType).To(Equal("text/javascript"))
			Expect(ops[1].Data.Type).To(Equal(0))
		})
	})

	Context(".PostForm", func() {
		It("should send the form request", func() {
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.Method).To(Equal("POST"))
				contextValue := request.Context().Value("key")
				Expect(contextValue).NotTo(BeNil())
				Expect(contextValue.(string)).To(Equal("value"))
				Expect(request.Header.Get("Content-Type")).To(Equal("application/x-www-form-urlencoded"))
				Expect(request.Header.Get("Accept")).To(Equal("application/json"))
				Expect(request.Header.Get("Authorization")).To(Equal("auth"))
				buf := new(bytes.Buffer)
				_, err := io.Copy(buf, request.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(buf.String()).To(Equal(`op=%2Fstat%2FdGVzdDprZXkx&op=%2Fstat%2FdGVzdDprZXky`))
			}).Return(fakeResponse, nil)

			values := make(url.Values)
			values.Add("op", "/stat/dGVzdDprZXkx")
			values.Add("op", "/stat/dGVzdDprZXky")
			ops := make([]BatchOpRet, 0, 2)
			err := postman.PostForm(
				ctx.WithValue(ctx.Background(), "key", "value"),
				"http://rs.qiniu.com/batch", "auth", values, &ops)
			Expect(err).NotTo(HaveOccurred())
			Expect(ops).To(HaveLen(2))
			Expect(ops[0].Code).To(Equal(200))
			Expect(ops[0].Data.Hash).To(Equal("FlsRbe0-0Cr3S6kqkGKl_-WDfbxF"))
			Expect(ops[0].Data.Fsize).To(Equal(int64(15508)))
			Expect(ops[0].Data.PutTime).To(Equal(int64(15040850589759659)))
			Expect(ops[0].Data.MimeType).To(Equal("text/css"))
			Expect(ops[0].Data.Type).To(Equal(0))
			Expect(ops[1].Code).To(Equal(200))
			Expect(ops[1].Data.Hash).To(Equal("FsALPkde4eIWUOod0VIYbdUlxx92"))
			Expect(ops[1].Data.Fsize).To(Equal(int64(158215)))
			Expect(ops[1].Data.PutTime).To(Equal(int64(15040850590777631)))
			Expect(ops[1].Data.MimeType).To(Equal("text/javascript"))
			Expect(ops[1].Data.Type).To(Equal(0))
		})
	})
})
