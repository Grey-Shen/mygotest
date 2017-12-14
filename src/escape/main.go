package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println(url.QueryEscape("/documents/reserve_append/:docId/pages/:pageNo"))
}
