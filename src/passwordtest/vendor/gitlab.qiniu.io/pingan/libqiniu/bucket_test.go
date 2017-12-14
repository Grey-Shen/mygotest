package libqiniu_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "gitlab.qiniu.io/pingan/libqiniu"
	"gitlab.qiniu.io/pingan/libqiniu/mock/mock_http"
	"golang.org/x/net/context"
)

var _ = Describe("Bucket", func() {
	var (
		ctrl             *gomock.Controller
		mockRoundTripper *mock_http.MockRoundTripper
		bucket           *Bucket
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
		client := NewClient(&http.Client{
			Transport: mockRoundTripper,
		})
		authorizedClient := client.Authorize(aksk)
		zonedClient := authorizedClient.Zone(zone)
		bucket = zonedClient.Bucket("testbucket")
	})

	Context(".Entry", func() {
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

		Context(".Op", func() {
			It("should stat key", func() {
				const body = `{"fsize":15508,"hash":"FlsRbe0-0Cr3S6kqkGKl_-WDfbxF","mimeType":"text/css","putTime":15040850589759659,"type":0}`
				fakeResponse.ContentLength = int64(len(body))
				fakeResponse.Body = ioutil.NopCloser(strings.NewReader(body))

				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/stat/dGVzdGJ1Y2tldDprZXkucG5n"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:mGzJkwkBXpKsYuDEtIh-B2NPB7I="))
				}).Return(fakeResponse, nil)

				fileInfo, err := bucket.Entry("key.png").Stat(context.Background())
				Expect(err).NotTo(HaveOccurred())
				Expect(fileInfo.Fsize).To(Equal(int64(15508)))
				Expect(fileInfo.Hash).To(Equal("FlsRbe0-0Cr3S6kqkGKl_-WDfbxF"))
				Expect(fileInfo.MimeType).To(Equal("text/css"))
				Expect(fileInfo.PutTime).To(Equal(int64(15040850589759659)))
				Expect(fileInfo.Type).To(Equal(TypeNormal))
			})

			It("should move key to another bucket", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/move/dGVzdGJ1Y2tldDprZXkucG5n/dGVzdGJ1Y2tldDI6a2V5Mi5wbmc=/force/false"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:DgdJMY9CaE8jb-xApcrLst1rgjQ="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").Move(context.Background(), "testbucket2", "key2.png", false)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should copy key to another bucket", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/copy/dGVzdGJ1Y2tldDprZXkucG5n/dGVzdGJ1Y2tldDI6a2V5Mi5wbmc=/force/true"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:Hwg3LjZzr5K31ccdznXjzEUCPA8="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").Copy(context.Background(), "testbucket2", "key2.png", true)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should change mime type", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/chgm/dGVzdGJ1Y2tldDprZXkucG5n/mime/dGV4dC9jc3M="))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:eY401v06ekM08AiHvUc0zN28-rg="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").ChangeMime(context.Background(), "text/css")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should change type", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/chtype/dGVzdGJ1Y2tldDprZXkucG5n/type/1"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:ADF9oAgxYxKSZWSxzS53DaI2_Ug="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").ChangeType(context.Background(), TypeInfrequent)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should delete the key", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/delete/dGVzdGJ1Y2tldDprZXkucG5n"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:1k0kPhe9guMCCumL9O004d_n1_M="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").Delete(context.Background())
				Expect(err).NotTo(HaveOccurred())
			})

			It("should delete the key after several days", func() {
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/deleteAfterDays/dGVzdGJ1Y2tldDprZXkucG5n/30"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:f2oPKVnx9gICPtTxpG6czLyp2b4="))
				}).Return(fakeResponse, nil)
				err := bucket.Entry("key.png").DeleteAfterDays(context.Background(), 30)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context(".Batch", func() {
			It("should send batch operations as one request", func() {
				const body = `[{"code":200,"data":{"fsize":15508,"hash":"FlsRbe0-0Cr3S6kqkGKl_-WDfbxF","mimeType":` +
					`"text/css","putTime":15040850589759659,"type":0}},{"code":200},{"code":200},{"code":200}]`
				fakeResponse.ContentLength = int64(len(body))
				fakeResponse.Body = ioutil.NopCloser(strings.NewReader(body))

				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(Equal(
						"http://rs.qiniu.com/batch"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:lXWlhnGwmP3ufdjG3v1VVSSsrzQ="))
					buf := new(bytes.Buffer)
					_, err := io.Copy(buf, request.Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(buf.String()).To(Equal("op=%2Fstat%2FdGVzdGJ1Y2tldDprZXkucG5n" +
						"&op=%2Fcopy%2FdGVzdGJ1Y2tldDprZXkucG5n%2FdGVzdDprZXkyLmpwZw%3D%3D%2Fforce%2Ftrue" +
						"&op=%2Fmove%2FdGVzdGJ1Y2tldDprZXkucG5n%2FdGVzdDI6a2V5Mi5qcGc%3D%2Fforce%2Ffalse" +
						"&op=%2Fdelete%2FdGVzdGJ1Y2tldDprZXkucG5n"))
				}).Return(fakeResponse, nil)

				ret, err := bucket.Batch().
					Stat("key.png").
					Copy("key.png", "test", "key2.jpg", true).
					Move("key.png", "test2", "key2.jpg", false).
					Delete("key.png").
					Do(context.Background())
				Expect(err).NotTo(HaveOccurred())
				Expect(ret).To(HaveLen(4))
				Expect(ret[0].Code).To(Equal(200))
				Expect(ret[0].Data.Hash).To(Equal("FlsRbe0-0Cr3S6kqkGKl_-WDfbxF"))
				Expect(ret[0].Data.Fsize).To(Equal(int64(15508)))
				Expect(ret[0].Data.MimeType).To(Equal("text/css"))
				Expect(ret[0].Data.PutTime).To(Equal(int64(15040850589759659)))
				Expect(ret[0].Data.Type).To(Equal(TypeNormal))
				Expect(ret[0].Data.Error).To(BeEmpty())
				Expect(ret[1].Code).To(Equal(200))
				Expect(ret[1].Data.Error).To(BeEmpty())
				Expect(ret[2].Code).To(Equal(200))
				Expect(ret[2].Data.Error).To(BeEmpty())
				Expect(ret[3].Code).To(Equal(200))
				Expect(ret[3].Data.Error).To(BeEmpty())
			})
		})
	})

	Context(".List", func() {
		It("should list all files", func() {
			var responseBodies = []string{
				`{"items":[{"key":"COD Modern Warfare Remastered Gameplay Trailer at E3 2016.mp4","hash":"lsDGZQy57xC2i4s7f-0lnUcHpvZy"` +
					`,"fsize":37835047,"mimeType":"video/mp4","putTime":14690946759277271,"type":0},{"key":"CentOS-6.5-x86_64-minimal.iso",` +
					`"hash":"lpuMwRRJmx5oVabCLaL0CGJlhnD0","fsize":417333248,"mimeType":"application/x-iso9660-image","putTime":14726988281881666,"type":0}],"marker":"abc"}`,
				`{"items":[{"key":"CentOS-6.8-x86_64-LiveDVD.iso","hash":"liIxIozpaxMIR18GlNFVETbeUroR","fsize":2012217344,"mimeType":"application/x-iso9660-image",` +
					`"putTime":14721434436923118,"type":0},{"key":"CentOS-7-x86_64-DVD.iso","hash":"ls-jhj3JnsSPeM7TtBWi9vsoaHn4","fsize":4335861760,` +
					`"mimeType":"application/x-iso9660-image","putTime":14713170025868914,"type":0}],"marker":"def"}`,
				`{"items":[{"key":"Mathworks.Matlab.R2012a.WINDOWS.iso","hash":"lpVWCWFGkKDDEgWQP0wn5Q0jUP6s","fsize":4803622912,"mimeType":` +
					`"application/x-iso9660-image","putTime":14832582923394153,"type":0}]}`,
			}

			responses := make([]http.Response, len(responseBodies))
			for i := 0; i < len(responseBodies); i++ {
				responses[i].Status = "200 OK"
				responses[i].StatusCode = 200
				responses[i].Proto = "HTTP/1.1"
				responses[i].ProtoMajor = 1
				responses[i].ProtoMinor = 1
				responses[i].Header = make(http.Header)
				responses[i].Header.Add("Content-Type", "application/json")
				responses[i].ContentLength = int64(len(responseBodies[i]))
				responses[i].Body = ioutil.NopCloser(strings.NewReader(responseBodies[i]))
			}

			gomock.InOrder(
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(HavePrefix(
						"http://rsf.qiniu.com/list?"))
					values := request.URL.Query()
					Expect(values).To(HaveLen(2))
					Expect(values.Get("bucket")).To(Equal("testbucket"))
					Expect(values.Get("limit")).To(Equal("1000"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:CbmJ80hbTpKZzu63XKdEgr1ckGc="))
				}).Return(&responses[0], nil),
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(HavePrefix(
						"http://rsf.qiniu.com/list?"))
					values := request.URL.Query()
					Expect(values).To(HaveLen(3))
					Expect(values.Get("bucket")).To(Equal("testbucket"))
					Expect(values.Get("limit")).To(Equal("1000"))
					Expect(values.Get("marker")).To(Equal("abc"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:ev_jfQvNOkZY2xqUbZJaUPVDIio="))
				}).Return(&responses[1], nil),
				mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).Do(func(request *http.Request) {
					Expect(request.URL.String()).To(HavePrefix(
						"http://rsf.qiniu.com/list?"))
					values := request.URL.Query()
					Expect(values).To(HaveLen(3))
					Expect(values.Get("bucket")).To(Equal("testbucket"))
					Expect(values.Get("limit")).To(Equal("1000"))
					Expect(values.Get("marker")).To(Equal("def"))
					Expect(request.Header.Get("Authorization")).To(Equal(
						"QBox yBq1NRj639CViLiXfKxsYFH6Rfw55QKVGZcsqOu_:r1DKjtqbIBUxsnjY5yOKAE2UcQ0="))
				}).Return(&responses[2], nil),
			)

			var file File

			iter := bucket.List(context.Background(), "")
			ok := iter.Next(&file)
			Expect(ok).To(BeTrue())
			Expect(file.Key).To(Equal("COD Modern Warfare Remastered Gameplay Trailer at E3 2016.mp4"))
			Expect(file.Hash).To(Equal("lsDGZQy57xC2i4s7f-0lnUcHpvZy"))
			Expect(file.Fsize).To(Equal(int64(37835047)))
			Expect(file.PutTime).To(Equal(int64(14690946759277271)))
			Expect(file.MimeType).To(Equal("video/mp4"))
			Expect(file.Type).To(Equal(TypeNormal))

			ok = iter.Next(&file)
			Expect(ok).To(BeTrue())
			Expect(file.Key).To(Equal("CentOS-6.5-x86_64-minimal.iso"))
			Expect(file.Hash).To(Equal("lpuMwRRJmx5oVabCLaL0CGJlhnD0"))
			Expect(file.Fsize).To(Equal(int64(417333248)))
			Expect(file.PutTime).To(Equal(int64(14726988281881666)))
			Expect(file.MimeType).To(Equal("application/x-iso9660-image"))
			Expect(file.Type).To(Equal(TypeNormal))

			ok = iter.Next(&file)
			Expect(ok).To(BeTrue())
			Expect(file.Key).To(Equal("CentOS-6.8-x86_64-LiveDVD.iso"))
			Expect(file.Hash).To(Equal("liIxIozpaxMIR18GlNFVETbeUroR"))
			Expect(file.Fsize).To(Equal(int64(2012217344)))
			Expect(file.PutTime).To(Equal(int64(14721434436923118)))
			Expect(file.MimeType).To(Equal("application/x-iso9660-image"))
			Expect(file.Type).To(Equal(TypeNormal))

			ok = iter.Next(&file)
			Expect(ok).To(BeTrue())
			Expect(file.Key).To(Equal("CentOS-7-x86_64-DVD.iso"))
			Expect(file.Hash).To(Equal("ls-jhj3JnsSPeM7TtBWi9vsoaHn4"))
			Expect(file.Fsize).To(Equal(int64(4335861760)))
			Expect(file.PutTime).To(Equal(int64(14713170025868914)))
			Expect(file.MimeType).To(Equal("application/x-iso9660-image"))
			Expect(file.Type).To(Equal(TypeNormal))

			ok = iter.Next(&file)
			Expect(ok).To(BeTrue())
			Expect(file.Key).To(Equal("Mathworks.Matlab.R2012a.WINDOWS.iso"))
			Expect(file.Hash).To(Equal("lpVWCWFGkKDDEgWQP0wn5Q0jUP6s"))
			Expect(file.Fsize).To(Equal(int64(4803622912)))
			Expect(file.PutTime).To(Equal(int64(14832582923394153)))
			Expect(file.MimeType).To(Equal("application/x-iso9660-image"))
			Expect(file.Type).To(Equal(TypeNormal))

			ok = iter.Next(&file)
			Expect(ok).To(BeFalse())
		})
	})
})
