package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Pass a context with a timeout to tell a blocking function that it
// 	// should abandon its work after the timeout elapses
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	go printHello(ctx)

// 	// select {
// 	// case <-time.After(1 * time.Second):
// 	// 	fmt.Println("overslept")
// 	// case <-ctx.Done():
// 	// 	fmt.Println(ctx.Err()) // prints "context deadline exceeded"
// 	// }
// 	<-ctx.Done()
// 	go printHello1()
// 	for {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("hello3")
// 	}
// 	fmt.Println("haha", ctx.Err())
// }

// func printHello(ctx context.Context) error {
// 	// time.Sleep(10 * time.Second)
// 	for {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("hello")
// 	}
// 	fmt.Println("i am ok")
// 	return nil
// }

// func printHello1() {
// 	for {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("hello1")
// 	}
// }
///////////////////////////////////////////////////////

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
