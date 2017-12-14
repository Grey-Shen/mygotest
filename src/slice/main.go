package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	haha(nil)
	var timetest time.Time
	fmt.Println("timetest", timetest)

	var tt [2]time.Time
	if tt[0].IsZero() {
		fmt.Println("aa")
	}
}

func haha(a []byte) {
	if len(a) == 0 {
		log.Println("test1") // print test1
	}
}
