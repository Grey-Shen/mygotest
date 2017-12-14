package libqiniu_test

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "gitlab.qiniu.io/pingan/libqiniu"
	"gitlab.qiniu.io/pingan/libqiniu/mock/mock_http"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aksk", func() {
	Context(".Sign", func() {
		deadline := time.Unix(1504443724, 0)
		It("Should sign the resource url", func() {
			u, err := url.Parse("http://www.qiniutest.com/image.jpg")
			Expect(err).NotTo(HaveOccurred())
			authURL := aksk.NewURL(u)
			Expect(authURL.PublicURL()).To(
				Equal("http://www.qiniutest.com/image.jpg"))
			Expect(authURL.SetDeadline(deadline).PrivateURL()).To(
				Equal("http://www.qiniutest.com/image.jpg?e=1504443724&token=yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:396a5z100nciFfG8LCV6-6eNqeU="))
		})

		It("Should sign the resource url with fop", func() {
			u, err := url.Parse("http://www.qiniutest.com/image.jpg")
			Expect(err).NotTo(HaveOccurred())
			authURL := aksk.NewURL(u)
			authURL.SetFop(NewFopCommand(imageslim{}))
			Expect(authURL.PublicURL()).To(
				Equal("http://www.qiniutest.com/image.jpg?imageslim"))
			Expect(authURL.SetDeadline(deadline).PrivateURL()).To(
				Equal("http://www.qiniutest.com/image.jpg?imageslim&e=1504443724&token=yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:4pLoS5PUPLQ-i1MAR9gPkSRJ6ak="))
		})

		It("Should sign the resource url with https", func() {
			u, err := url.Parse("http://www.qiniutest.com/image.jpg")
			Expect(err).NotTo(HaveOccurred())
			authURL := aksk.NewURL(u).HTTPs()
			Expect(authURL.PublicURL()).To(
				Equal("https://www.qiniutest.com/image.jpg"))
			Expect(authURL.SetDeadline(deadline).PrivateURL()).To(
				Equal("https://www.qiniutest.com/image.jpg?e=1504443724&token=yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:QjKNZ6fghkwwK2Na4u_LEdqpQNA="))
		})
	})

	Context(".Uptoken", func() {
		It("Should generate uptoken from put policy", func() {
			const token = "yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:Lq8HNxTDeB14tRmB1wt6EdC5vK0=:eyJzY29wZSI6InRlc3Q6aW1hZ2UuanBnIiwiZGVhZGxpbmUiOjE1MDQ0NDgyMDcsInNhdmVLZXkiOiJpbWFnZS5qcGcifQ=="
			deadline := time.Unix(1504448207, 0)
			putPolicy := NewPutPolicy("test", "image.jpg").SetDeadline(deadline)
			generatedUptoken, err := aksk.Uptoken(&putPolicy)
			Expect(err).NotTo(HaveOccurred())
			Expect(generatedUptoken.String()).To(Equal(token))
			generatedUptokenAgain, err := aksk.Uptoken(generatedUptoken)
			Expect(generatedUptokenAgain.String()).To(Equal(token))
		})
	})

	Context(".SignedTransport", func() {
		var ctrl *gomock.Controller

		BeforeEach(func() {
			ctrl = gomock.NewController(GinkgoT())
		})
		AfterEach(func() {
			ctrl.Finish()
		})

		It("Should create a signed transport to sign json request", func() {
			mockRoundTripper := mock_http.NewMockRoundTripper(ctrl)
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.Header.Get("Authorization")).To(Equal(
					"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:XUcJn4nlJnmOP5zFKpZlUXiPd2s="))
				buf := new(bytes.Buffer)
				_, err := io.Copy(buf, request.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(buf.String()).To(Equal(`{"name":"test"}`))
			}).Return(nil, nil)
			roundTripper := aksk.SignedTransport(mockRoundTripper)

			request, err := http.NewRequest("POST", "http://www.qiniutest.com/allkeys", strings.NewReader(`{"name":"test"}`))
			Expect(err).NotTo(HaveOccurred())
			request.Header.Set("Content-Type", "application/json")
			_, err = roundTripper.RoundTrip(request)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should create a signed transport to sign post form request", func() {
			mockRoundTripper := mock_http.NewMockRoundTripper(ctrl)
			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
				Expect(request.Header.Get("Authorization")).To(Equal(
					"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:lHxr7aWcrdnDCRsiskJUO9WGN2Y="))
				buf := new(bytes.Buffer)
				_, err := io.Copy(buf, request.Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(buf.String()).To(Equal(`name=test&age=20`))
			}).Return(nil, nil)
			roundTripper := aksk.SignedTransport(mockRoundTripper)

			request, err := http.NewRequest("POST", "http://www.qiniutest.com/allkeys", strings.NewReader(`name=test&age=20`))
			Expect(err).NotTo(HaveOccurred())
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			_, err = roundTripper.RoundTrip(request)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context(".ValidateCallback", func() {
		It("should validate json request", func() {
			request, err := http.NewRequest("POST", "http://www.qiniutest.com/allkeys", strings.NewReader(`{"name":"test"}`))
			Expect(err).NotTo(HaveOccurred())
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Authorization", "QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:XUcJn4nlJnmOP5zFKpZlUXiPd2s=")
			ok, err := aksk.ValidateCallback(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeTrue())
		})

		It("should validate post form request", func() {
			request, err := http.NewRequest("POST", "http://www.qiniutest.com/allkeys", strings.NewReader(`name=test&age=20`))
			Expect(err).NotTo(HaveOccurred())
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request.Header.Set("Authorization", "QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:1InzZpKZljVxlHdjAI_c1LZD_HY=")
			ok, err := aksk.ValidateCallback(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeFalse())

			request2, err := http.NewRequest("POST", "http://www.qiniutest.com/allkeys", strings.NewReader(`name=test&age=20`))
			Expect(err).NotTo(HaveOccurred())
			request2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			request2.Header.Set("Authorization", "QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:lHxr7aWcrdnDCRsiskJUO9WGN2Y=")
			ok, err = aksk.ValidateCallback(request2)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeTrue())
		})
	})
})
