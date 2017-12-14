package main

import "fmt"

const (
	AppId     = "bbd90991"
	AppKey    = "8f2236e2e5d1f2bbd351165e45eefb49b82d81c8c09e33e6f9baec5c496455e6"
	AccessKey = "fakeaccesskey"
	SecretKey = "fakesecretkey"
)

func Authorization(method, requestURI string) string {
	return CalcAuthorization(method, requestURI, "APPKEY", AppId, AppKey)
}

func main() {
    SetHeader(gofight.H{"Authorization": test.Authorization("post", "/documents/reserve_upload/"+test.AppId+"?pack=yes")}).
    const JsonStream =
    `{
        "title":   "faketitle",
        "bizdata": gofight.D{"bizkey1": "bizval1", "bizkey2": "bizval2"},
            "pages": []gofight.D{
            {"bizdata": gofight.D{"bizkey3": "bizval3", "bizkey4": "bizval4"}},
            {"bizdata": gofight.D{"bizkey5": "bizval5", "bizkey6": "bizval6"}},
            {},
        },
        "created_by": "bachue",
    }`

}