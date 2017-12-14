package main

import (
	"fmt"
	"net/url"
)

func main() {
	// if result, err := base64.URLEncoding.DecodeString("8ehryqviSaMIjkVQDGeDcKRZ6qc="); err != nil {
	// 	fmt.Println("decode failed", err)
	// } else {
	// 	fmt.Println("result", string(result))
	// }
	desc := url.QueryEscape("这是一个测试")
	fmt.Println(desc)
	if descUnEscape, err := url.QueryUnescape("%E8%BF%99%E6%98%AF%E4%B8%80%E4%B8%AA%E6%B5%8B%E8%AF%95"); err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(descUnEscape)
	}
}
