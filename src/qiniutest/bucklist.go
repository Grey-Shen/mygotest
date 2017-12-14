package main

import (
	"fmt"
	"io"

	"qiniupkg.com/api.v7/kodo"
)

const (
	ACCESS_KEY = "C2QcWBwdgB5XIPgexk9LoAJp4uJoW0yaHywpqP0q"
	SECRET_KEY = "v-RAMVOIuLKBGCFemR1omfSAGlc8NgGxqTudEnPQ"
)

var childPath string

func main() {
	kodo.SetMac(ACCESS_KEY, SECRET_KEY)
	client := kodo.New(0, nil)
	bucket := client.Bucket("shenyajun-bucket")
	items, prefix, markout, err := bucket.List(nil, "", ",", "", -1)
	if err != nil {
		if err == io.EOF {
			fmt.Println("list all")
		} else {
			fmt.Println("failed to list all keys: %s", err)
		}
	}
	for _, items := range items {
		fmt.Println("items", items)
	}
	fmt.Println("prefix", prefix)
	fmt.Println("markout", markout)
}

// type ListItem struct {
// 	Key      string `json:"key"`
// 	Hash     string `json:"hash"`
// 	Fsize    int64  `json:"fsize"`
// 	PutTime  int64  `json:"putTime"`
// 	MimeType string `json:"mimeType"`
// 	EndUser  string `json:"endUser"`
// }

// 首次请求，请将 marker 设置为 ""。
// 无论 err 值如何，均应该先看 entries 是否有内容。
// 如果后续没有更多数据，err 返回 EOF，markerOut 返回 ""（但不通过该特征来判断是否结束）。

// func (p Bucket) List(
// 	ctx Context, prefix, delimiter, marker string, limit int) (entries []ListItem, commonPrefixes []string, markerOut string, err error) {

// 	listUrl := p.makeListURL(prefix, delimiter, marker, limit)

// 	var listRet struct {
// 		Marker   string     `json:"marker"`
// 		Items    []ListItem `json:"items"`
// 		Prefixes []string   `json:"commonPrefixes"`
// 	}
// 	err = p.Conn.Call(ctx, &listRet, "POST", listUrl)
// 	if err != nil {
// 		return
// 	}
// 	if listRet.Marker == "" {
// 		return listRet.Items, listRet.Prefixes, "", io.EOF
// 	}
// 	return listRet.Items, listRet.Prefixes, listRet.Marker, nil
// }
