// package main

// import "fmt"

// func main() {
// 	appId := "c1c81a8d"
// 	like := fmt.Sprintf(`/^UDMP\-PACKAGES\-[0-9a-f]{42}\-%s\-\d{14}$/`, appId)
// 	fmt.Println("like is", like)
// 	a := []string{"a", "b", "c"}
// 	b := []string{"d", "e", "f"}
// 	a = append(a, b...)
// 	fmt.Println("a is ", a)
// 	fmt.Println("rune", rune('0'))

// 	str := "Hello, 世界"

// 	// r, size := utf8.DecodeRune(str[0])
// 	fmt.Println(rune(str[0]))

// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	// const body = "Go is a general-purpose language designed with systems programming in mind."
	const body = ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Wed, 19 Jul 1972 19:00:00 GMT")
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%qhah", dump)

}

// 9328c9d5

// /^UDMP\-PACKAGES\-[0-9a-f]{42}\-c1c81a8d\-\d{14}$/
