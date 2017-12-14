package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"qiniupkg.com/api.v7/kodo"
)

const (
	ACCESS_KEY = "C2QcWBwdgB5XIPgexk9LoAJp4uJoW0yaHywpqP0q"
	SECRET_KEY = "v-RAMVOIuLKBGCFemR1omfSAGlc8NgGxqTudEnPQ"
)

const chunk = 65535

func main() {
	kodo.SetMac(ACCESS_KEY, SECRET_KEY)
	client := kodo.New(0, nil)
	buff := bytes.NewBuffer([]byte{})
	request, _ := http.NewRequest("PUT", "http://ofmufa5nc.bkt.clouddn.com/gopl-zh.pdf", bytes.NewReader(bs))

	request.Header.Set("Range", value)
	resp, _ := client.Do(request)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	b, _ := ioutil.ReadAll(resp.Body)
	_, err := buff.Write(b)
	if err != nil {
		return err
	}
	// defer resp.Body.Close()
	fmt.Println(string(b))
}

POST http://rsf.pabosuat.sdb.com.cn/list?bucket=udmp-storage
HTTP/1.1 1 1 map[User-Agent:[QiniuGo/7.0.5 (linux; amd64; Tassadar) go1.8]
Authorization:[QBox LtqZYkembgXMbUAt-dSLwafEOTcg2r3VQtb9z2RS:4H4xb35dnAi2PjyiuhzawfuyL8E=]
Content-Type:[application/x-www-form-urlencoded]]
<nil> <nil> 0 [] false rsf.pabosuat.sdb.com.cn map[] map[] <nil> map[]   <nil> <nil> <nil> <nil>}