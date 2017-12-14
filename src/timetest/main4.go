package main

import (
	"fmt"
	"time"
)

const (
	DateFormat  = "2006-01-02"
	beginString = "2017-10-19"
)

func main() {
	location, _ := time.LoadLocation("Local")
	begin, err := time.ParseInLocation(DateFormat, beginString, location)
	if err != nil {
		return
	}

	fmt.Println("begintest", begin)
}
