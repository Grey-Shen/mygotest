package main

import (
	"fmt"
	_ "net/http/pprof"
	"strconv"
	"time"
)

// func main() {
// 	// go http.ListenAndServe(":8989", nil)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		cancel()
// 	}()
// 	A(ctx)

// 	// // no context
// 	// a := make(chan bool)
// 	// go A_a(a)
// 	// time.Sleep(time.Second * 5)
// 	// a <- true
// 	select {}
// }

// func C(ctx context.Context) {
// 	log.Println("Running in C")
// OuterLoop:
// 	for {
// 		log.Println("C.....")
// 		time.Sleep(time.Second * 1)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("C Done")
// 			break OuterLoop
// 		default:
// 		}
// 	}
// }

// func B(ctx context.Context) {
// 	log.Println("Running in B")
// 	ctx, _ = context.WithCancel(ctx)
// 	go C(ctx)
// OuterLoop:
// 	for {
// 		time.Sleep(time.Second * 1)
// 		log.Println("B....")
// 		select {
// 		case <-ctx.Done():
// 			log.Println("B Done")
// 			break OuterLoop
// 		default:
// 		}
// 	}
// }

// func A(ctx context.Context) {
// 	log.Println("Running in A")
// 	go B(ctx)
// OuterLoop:
// 	for {
// 		log.Println("A....")
// 		time.Sleep(time.Second * 1)
// 		select {
// 		case <-ctx.Done():
// 			log.Println("A Done")
// 			break OuterLoop
// 		default:
// 		}
// 	}
// }

// func C_c() {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		log.Println("C_c ......")
// 	}
// }

// func B_b() {
// 	go C_c()
// 	for {
// 		time.Sleep(time.Second * 1)
// 		log.Println("B_b ......")
// 	}
// }

// func A_a(a chan bool) {
// 	go B_b()
// OuterLoop:
// 	for {
// 		time.Sleep(time.Second * 1)
// 		log.Println("A_a ......")
// 		select {
// 		case <-a:
// 			log.Println("A_a is over")
// 			break OuterLoop
// 		default:
// 		}
// 	}
// }

// func main() {
// 	var a = map[string]interface{}{}
// 	var b = map[string]interface{}{}
// 	// var c = map[string]interface{}{}
// 	b["abc"] = "def"
// 	a["a"] = b
// 	b = map[string]interface{}{}
// 	// b = map[string]interface{}{}
// 	a["a"].(map[string]interface{})["mm"] = "bb"
// 	fmt.Println(a["a"])

// 	// a1["test"] = test
// }

// func main() {
// 	// Pass a context with a timeout to tell a blocking function that it
// 	// should abandon its work after the timeout elapses.
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	go func(ctx context.Context) {
// 		for {
// 			time.Sleep(time.Second * 1)
// 			fmt.Println("11111111")
// 		}
// 	}(ctx)

// 	select {
// 	case <-time.After(20 * time.Second):
// 		fmt.Println("overslept")
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
// 	}

// }

// func main() {
// 	// gen generates integers in a separate goroutine and
// 	// sends them to the returned channel.
// 	// The callers of gen need to cancel the context once
// 	// they are done consuming generated integers not to leak
// 	// the internal goroutine started by gen.
// 	gen := func(ctx context.Context) <-chan int {
// 		dst := make(chan int)
// 		n := 1
// 		go func() {
// 			for {
// 				fmt.Println("loop")
// 				select {
// 				case <-ctx.Done():
// 					return // returning not to leak the goroutine
// 				case dst <- n:
// 					n++
// 				}
// 			}
// 		}()
// 		return dst
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel() // cancel when we are finished consuming integers

// 	// test := gen(ctx)

// 	// for n := range test {
// 	// 	fmt.Println(n)
// 	// 	if n == 5 {
// 	// 		break
// 	// 	}
// 	// }
// 	gen(ctx)
// }

// type Foo struct {
// 	bar string
// }

// func main() {
// 	list := []Foo{
// 		{"A"},
// 		{"B"},
// 		{"C"},
// 	}

// 	list2 := make([]*Foo, len(list))

// 	//错误的例子
// 	for i, value := range list {
// 		list2[i] = &value
// 	}

// 	list3 := make([]Foo, len(list))

// 	for i, value := range list {
// 		list3[i] = value
// 		}

// 	//正确的例子
// 	//for i, _ := range list {
// 	//	list2[i] = &list[i]
// 	//}

// 	fmt.Println(list[0], list[1], list[2])
// 	fmt.Println(list2[0], list2[1], list2[2])
// 	fmt.Println(list3[0], list3[1], list3[2])
// }

func main() {
	var numList = []int{1, 2, 11, 1692}
	begin := time.Now()
	for i := 0; i < 1000000; i++ {
		for _, target := range numList {
			end := 0
			b := 1
			for {
				result := target * b
				b++
				s := strconv.Itoa(result)

				for _, num := range s {
					// if _, ok := allNums[num]; !ok {
					// 	allNums[num] = true
					// }
					n, _ := strconv.Atoi(string(num))
					if n == 0 {
						end = end | 1
					} else {
						d := 1 << uint(n)
						end = end | d
					}
				}

				// fmt.Printf("end is %b\n", end)
				if end == 1023 {
					// fmt.Println("s is", s)
					break
				}
				// if len(allNums) == 10 {
				// 	fmt.Println("s is ", s)
				// 	break
				// }

			}
		}
	}

	fmt.Println("benchmark: ", time.Since(begin))

}

// func main() {
// 	var result int
// 	a := 123
// 	b := strconv.Itoa(a)
// 	fmt.Println("b is", b)
// 	for _, num := range b {
// 		n, _ := strconv.Atoi(string(num))
// 		fmt.Println("num is", n)
// 		d := 1 << uint(n-1)
// 		// fmt.Println("d is ", d)
// 		fmt.Printf("d is %b\n", d)
// 		result = result | d
// 		// fmt.Println("result is", result)
// 		fmt.Printf("result is %b\n", result)
// 	}
// }
