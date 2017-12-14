package main

import (
	"fmt"
	"strings"
	"time"
)

// import (
// 	"fmt"
// 	"time"
// )

// const dateFormat = "20060102150304"

// func main() {
// 	ticker := time.NewTicker(time.Second * 1)
// 	go getticker(ticker.C)
// 	time.Sleep(time.Second * 10)
// }

// func getticker(c <-chan time.Time) {
// 	for s := range c {
// 		fmt.Println("ticker....", s.Format("2006-01-02 15:04:05"))
// 	}
// }

// func main() {
// 	cookie := http.Cookie{
// 		Name:     "ECLOUD_SESSION",
// 		Value:    "MFRIOEk1Uk9BNlZZM1lZUVg3U0FJTTA2VkY4R0ZaMUEsMTQ5MjA2Njk0ODQ5MjQ1MDM1MSw0Njg0NGEwZmE3ZGNkNTMyYjI0YTNkZTU2Njk4MTM0OTdkZTgxYzg5",
// 		Path:     "/",
// 		HttpOnly: true,
// 	}

// 	fmt.Println(cookie.String())
// }

const (
	keyRefreshInterval = 1 * time.Second
)

func init() {
	ticker := time.NewTicker(keyRefreshInterval)
	go func(signalChannel <-chan time.Time) {
		fmt.Println("22222")
		for s := range signalChannel {
			fmt.Println(s)
		}
		fmt.Println("333")
	}(ticker.C)
}

func main() {
	// var dates [2]time.Time
	// for i, date := range dates {
	// 	fmt.Println("i", i, "date", date)
	// }

	// ticker := time.NewTicker(time.Second * keyRefreshInterval)
	// // go func(signalChannel <-chan time.Time) {
	// for s := range ticker.C {
	// 	fmt.Println("222%s", s.Format("2006-01-02 15:04:05"))
	// }
	// // }(ticker.C)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("1111")
	// 	}
	// }()
	// wg.Wait()
	// timestring := time.Now().Format(time.RFC3339)
	// mytime, err := time.Parse(time.RFC3339, timestring)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mytime)
	aa := []int{1, 3, 4, 5}
	for _, k := range aa {
		fmt.Println(k)
	}

	s := []string{"foo"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

	type HAHA struct {
		haha []string
	}
	tt := HAHA{}
	fmt.Println(len(tt.haha))

	bb := []string{}
	cc := []string{"a", "b"}
	for _, k := range cc {
		bb = append(bb, k)
	}
	for _, k := range bb {
		fmt.Println(k)
	}

	var test = ""
	fmt.Println("lal", strings.Split(test, ","))
}
