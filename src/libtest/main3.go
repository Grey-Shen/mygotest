package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	var numList = []int{1, 11, 1692, 2}
	begin := time.Now()
	wg.Add(1)
	go execute(numList, 50000)
	wg.Add(1)
	go execute(numList, 50000)
	wg.Wait()
	fmt.Println("Benchmark: ", time.Since(begin))
}

func execute(numList []int, num int) {
	fmt.Println("begin---", numList)
	for _, target := range numList {
		for i := 0; i < num; i++ {
			end := 0
			b := 1
			for {
				result := target * b
				b++
				s := strconv.Itoa(result)
				for _, num := range s {
					n, _ := strconv.Atoi(string(num))
					if n == 0 {
						end = end | 1
					} else {
						d := 1 << uint(n)
						end = end | d
					}
				}

				if end == 1023 {
					break
				}
			}
		}
	}
	fmt.Println("done---", numList)
	wg.Done()
}
