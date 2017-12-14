package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var numList = []int{1, 2, 11, 1692}
	begin := time.Now()
	for _, target := range numList {
		wg.Add(1)
		go func(target int) {
			for i := 0; i < 1000000; i++ {
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
			wg.Done()
		}(target)
	}
	wg.Wait()
	fmt.Println("Benchmark: ", time.Since(begin))
}
