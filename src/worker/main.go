package main

import (
	"fmt"
	"time"
)

func main() {
	var taskPool = make(chan int)
	go producer(taskPool)
	for i := 0; i < 5; i++ {
		go consumer(taskPool, i)
	}
	time.Sleep(100 * time.Second)
}

func producer(taskPool chan<- int) {
	var count int
	for {
		if len(taskPool) > 0 {
			count++
			taskPool <- count
			fmt.Println("poolsize", len(taskPool))
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("producer: I am full")
		}
	}
}

func consumer(taskPool <-chan int, i int) {
	for count := range taskPool {
		fmt.Println("consumer: ", i, "count: ", count)
		time.Sleep(10 * time.Second)
		fmt.Println("consumer: ", i, "I am sleeping")
	}
}
