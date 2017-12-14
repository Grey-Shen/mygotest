package main

import "fmt"

const (
	begin = 1
	end   = 4
)

type person struct {
	address []string
}

func main() {
	// year, month, day := time.Now().Date()
	// beginTime := time.Date(year, month, day, begin, 0, 0, 0, time.Local)
	// fmt.Println("beginTime", beginTime)
	// endTime := time.Date(year, month, day, end, 0, 0, 0, nil)
	// nextBeginTime := time.Now().Add(30 * time.Second)
	// fmt.Println("nextBegintime", nextBeginTime)
	// tmp := nextBeginTime.Sub(time.Now())
	// time.Sleep(tmp)
	// fmt.Println("tmp", tmp.Nanoseconds())

	// const intSize = 32 << (^uint(0) >> 63)
	// fmt.Println(intSize)

	li := new(person)
	li.address = []string{"beijing", "shanghai"}
	fmt.Println(li)
}

// func test(a int, test ...bool) {
// 	if len(test) > 0 {
// 		if test[0] {
// 			fmt.Println("hello", a)
// 		}
// 	} else {

// 		fmt.Println("world", a)
// 	}

// 	// s
// 	loc, err := time.LoadLocation("Asia/Shanghai")
// 	if err != nil {
// 		fmt.Println("zoneerr", err)
// 	}
// 	fmt.Println(time.Now().In(loc))
// }
