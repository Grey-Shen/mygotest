package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	var myTransport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		c := http.Client{
			Transport: myTransport,
			Timeout:   60 * time.Second,
		}

		rb := bytes.NewBufferString("hello")

		req, _ := http.NewRequest("GET", "http://localhost:8889/hijack", rb)

		if resp, err := c.Do(req); err != nil {
			panic(err)
		} else {
			fmt.Println("====length=====:", resp.ContentLength)
			bs := make([]byte, 0, resp.ContentLength)
			buf := make([]byte, 1)
			for {
				if _, err := resp.Body.Read(buf); err != nil {
					if err == io.EOF {
						break
					} else {
						panic(err)
					}
				} else {
					fmt.Println("====1=", string(buf))
					bs = append(bs, buf...)
				}
			}
			fmt.Println("========", string(buf))
		}
	}
}
