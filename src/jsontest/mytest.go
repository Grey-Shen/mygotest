package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type myCookieJar struct {
	cookies   []*http.Cookie
	cookieUrl string
	expireAt  time.Time
}

func (m *myCookieJar) SetCookies() {

}

const (
	ACCESS_KEY         = "C2QcWBwdgB5XIPgexk9LoAJp4uJoW0yaHywpqP0q"
	SECRET_KEY         = "v-RAMVOIuLKBGCFemR1omfSAGlc8NgGxqTudEnPQ"
	authURL            = "http://115.231.180.112:3000/api/session"
	createNamespaceURL = "http://115.231.180.112:3000/api/namespace"
)

type PasswordCredentialRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (request *PasswordCredentialRequest) MarshalReader() (io.Reader, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Marshal PasswordCredentialRequest failed: %s", err)
	}
	return bytes.NewReader(b), nil
}

type NamespaceCreateParams struct {
	Name string `json:"name"`
}

func NewNamespaceCreateParams(username string) *NamespaceCreateParams {
	return &NamespaceCreateParams{
		Name: username,
	}
}

func (params *NamespaceCreateParams) MarshalReader() (io.Reader, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("Marshal UserGetNamespaceParams failed: %s", err)
	}
	return bytes.NewReader(b), nil
}

func main() {

	credential := PasswordCredentialRequest{
		Username: "admin1@qiniu.com",
		Password: "hao123456",
	}

	reader, err := credential.MarshalReader()
	if err != nil {
		return
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=====login and get cookies========")
	client := &http.Client{Jar: jar}
	request, _ := http.NewRequest("PUT", authURL, reader)
	resp, _ := client.Do(request)
	fmt.Println("code:", resp.StatusCode)
	fmt.Println("cookies:", resp.Cookies())
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	resp.Body.Close()

	fmt.Println("===========create namespace=======")
	namespaceParams := NewNamespaceCreateParams("test10003")
	namespaceReader, err := namespaceParams.MarshalReader()
	if err != nil {
		return
	}
	request1, _ := http.NewRequest("PUT", createNamespaceURL, namespaceReader)
	resp1, _ := client.Do(request1)
	fmt.Println(resp1.StatusCode)
	fmt.Println("cookies1:", resp1.Cookies())
	b1, _ := ioutil.ReadAll(resp1.Body)
	myurl, _ := url.Parse(createNamespaceURL)
	fmt.Println("cookies", client.Jar.Cookies(myurl))
	fmt.Println(string(b1))
	resp1.Body.Close()
}
