package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	jobBegin = make(chan int)
	jobStop  = make(chan os.Signal)
)

func Init() {
	signal.Notify(jobStop, syscall.SIGUSR1)
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func(signal <-chan time.Time) {
		for s := range signal {
			if s.Second() == 5 {
				fmt.Println("=== second ===", s.Second())
				jobBegin <- s.Second()
			}
		}
	}(ticker.C)

	for {
		select {
		case <-jobBegin:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			go jobStart(ctx)
		case <-jobStop:
			fmt.Println("==get===")
			return
		default:
			fmt.Println("hah")
		}
	}
}

func jobStart(ctx context.Context) {

	for {
		log.Println("doing=====")
		select {
		case <-ctx.Done():
			fmt.Println("done=====")
			return
		default:
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	Init()
	log.Println("=====I am here====")
	<-time.After(5 * time.Minute)
}
