package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var serverDone = make(chan struct{})
var serverDone1 = make(chan struct{})
var serverDone2 = make(chan struct{})
var serverDone3 = make(chan struct{})
var serverDone4 = make(chan struct{})
var serverDone5 = make(chan struct{})

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000; i++ {
		go messageLoop()
	}
	<-time.After(10 * time.Second)
	close(serverDone)
	fmt.Println("finished")
}

func messageLoop() {
	var ticker = time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	var counter = 0
	for {
		select {
		case <-serverDone:
			return
		case <-serverDone1:
			return
		// case <-serverDone2:
		//  return
		// case <-serverDone3:
		//  return
		// case <-serverDone4:
		//  return
		// case <-serverDone5:
		//  return
		case <-ticker.C:
			counter += 1
		}
	}
}
