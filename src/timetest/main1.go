package main

import (
	"context"
	"fmt"
	"time"
)

type mytest struct {
	b func(a int)
}

var a map[string]func(b int)

// func main() {
// 	log.Println("11==", time.Now().Day())
// 	y, m, d := time.Now().Date()
// 	log.Println("22==", time.Date(y, m, d, 3, 0, 0, 0, time.UTC))
// 	log.Println('a')

// 	mytimer := time.NewTimer(5 * time.Second)

// 	go func(signal <-chan time.Time) {
// 		select {
// 		case <-signal:
// 			// for {
// 			fmt.Println("hahaha")
// 			// time.Sleep(1 * time.Second)
// 			// }
// 		}
// 	}(mytimer.C)

// 	mytimer.Reset(5 * time.Second)

// 	<-time.After(50 * time.Second)
// }

func main() {
	// ticker := time.NewTicker(2 * time.Second)

	// for {
	// 	<-ticker.C
	// 	fmt.Println("haha")
	// }

	// <-time.After(30 * time.Second)

	f := func(a int) {
		fmt.Println("====", a)
	}

	test := mytest{b: f}

	test.b(4)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			fmt.Println("hah")
			select {
			case <-ctx.Done():
				fmt.Println("time out")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("22")
			}
		}
	}(ctx)

	<-time.After(10 * time.Second)
}
