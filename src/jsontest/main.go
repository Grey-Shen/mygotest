package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	type ColorGroup struct {
		Name     string `json:"username"`
		Password string `json:"password"`
	}
	group := ColorGroup{
		Name:     "admin1@qiniu.com",
		Password: "hao123456",
	}

	bs, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	// params := make(url.Values)
	// params.Add("username", "admin1@qiniu.com")
	// params.Add("password", "hao123456")

	// reader1 := strings.NewReader(params.Encode())

	client := &http.Client{}
	request, _ := http.NewRequest("PUT", "http://115.231.180.112:3000/api/session", bytes.NewReader(bs))
	resp, _ := client.Do(request)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Cookies())
	b, _ := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()
	fmt.Println(string(b))
}
