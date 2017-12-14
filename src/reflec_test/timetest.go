package main

import (
	"fmt"
	"time"
)

var tt = "2h5m"

type AppParams struct {
	AppName  string `json:"app_name" binding:"required,min=4"`
	FullName string `json:"full_name"`
}

func main() {
	dd, err := time.ParseDuration(tt)
	if err != nil {
		fmt.Println("Parse err:", err)
	}

	fmt.Println("dd is", dd)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	mytime := time.Now()
	ttjson, _ := mytime.MarshalBinary()
	fmt.Println(string(ttjson))

}
