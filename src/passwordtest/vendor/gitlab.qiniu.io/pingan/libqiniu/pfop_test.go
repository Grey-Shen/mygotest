package libqiniu_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.qiniu.io/pingan/libhttpclient/form_postman"
	. "gitlab.qiniu.io/pingan/libqiniu"
	"gitlab.qiniu.io/pingan/libqiniu/mock/mock_http"
	"gitlab.qiniu.io/pingan/libqiniu/op"
)

var _ = Describe("Upload", func() {
	var (
		ctrl             *gomock.Controller
		pfopClient       PfopClient
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
		pfopClient = NewPfopClient(postman, aksk, zone)
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

	Context(".Pfop", func() {
		It("Should execute pfop command and save as another key", func() {
			fakeResponse.Body = ioutil.NopCloser(strings.NewReader(`{"persistentID":"testid"}`))
			fakeResponse.ContentLength = int64(len(`{"persistentID":"testid"}`))
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.Header.Get("Authorization")).To(Equal("QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:MruOx-KTgjea8NIDh0nKe1JzrVs="))
				body, err := ioutil.ReadAll(request.Body)
				Expect(err).NotTo(HaveOccurred())
				values, err := url.ParseQuery(string(body))
				Expect(err).NotTo(HaveOccurred())
				Expect(values.Get("bucket")).To(Equal("testbucket"))
				Expect(values.Get("key")).To(Equal("fakekey"))
				Expect(values.Get("fops")).To(Equal("imageslim|saveas/dGVzdGJ1Y2tldDphbm90aGVya2V5"))
				Expect(values.Get("notifyURL")).To(Equal("http://fake.callback.com/pfop/callback"))
				Expect(values.Get("force")).To(Equal("1"))
			}).Return(fakeResponse, nil)

			persistentID, err := pfopClient.Pfop(
				context.Background(), op.Entry{Bucket: "testbucket", Key: "fakekey"}, PfopParams{
					Fops:      NewFopCommand(imageslim{}).ToPipeline().SaveAs("testbucket", "anotherkey"),
					NotifyURL: "http://fake.callback.com/pfop/callback",
					Force:     true,
				})
			Expect(err).NotTo(HaveOccurred())
			Expect(persistentID.String()).To(Equal("testid"))
		})
	})

	Context(".Prefop", func() {
		It("Should query pfop status", func() {
			const body = `{"code":0,"desc":"The fop was completed successfully","id":"testid","inputBucket":"testbucket","inputKey":"testkey","items":[{"cmd":"imageslim|saveas/dGVzdGJ1Y2tldDphbm90aGVya2V5","code":0,"desc":"The fop was completed successfully","hash":"lpK9FqMroCoT_qTocTwXhwf7tDKb","key":"anotherkey","returnOld":0}],"pipeline":"default.libqiniu-mps","reqid":"YSYAACkgwJvNs80U"}`
			fakeResponse.Body = ioutil.NopCloser(strings.NewReader(body))
			fakeResponse.ContentLength = int64(len(body))

			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.URL.String()).To(Equal("http://status.qiniu.com/status/get/prefop?id=testid"))
				Expect(request.Header.Get("Authorization")).To(BeEmpty())
			}).Return(fakeResponse, nil)

			status, err := pfopClient.Prefop(context.Background(), "testid")
			Expect(err).NotTo(HaveOccurred())
			Expect(status.Id).To(Equal("testid"))
			Expect(status.Code).To(Equal(PFOP_SUCCESSFUL))
			Expect(status.Desc).To(Equal("The fop was completed successfully"))
			Expect(status.InputBucket).To(Equal("testbucket"))
			Expect(status.InputKey).To(Equal("testkey"))
			Expect(status.Items).To(HaveLen(1))
			Expect(status.Items[0].Code).To(Equal(PFOP_SUCCESSFUL))
			Expect(status.Items[0].Error).To(BeEmpty())
			Expect(status.Items[0].Key).To(Equal("anotherkey"))
			Expect(status.Items[0].ReturnOld).To(Equal(op.False))
			Expect(status.String()).To(Equal(body))
		})
	})
})
