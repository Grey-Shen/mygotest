package main

import (
	"fmt"
	"os"
	"os/signal"
)

var ()

// func main() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, syscall.SIGUSR1)
// 	s := <-c
// 	log.Printf("Got signal: %s\n", s.String())

// 	// go test(c)
// 	// <-time.After(5 * time.Minute)
// }

// func test(c <-chan os.Signal) {
// 	// for {
// 	// 	log.Println("===test=====")
// 	// 	select {
// 	// 	case <-c:
// 	// 		log.Println("===over=====")
// 	// 		return
// 	// 	default:
// 	// 		time.Sleep(1 * time.Second)
// 	// 	}
// 	// }
// 	log.Println("===wait====")
// 	for s := range c {
// 		log.Printf("Got signal: %s\n", s.String())
// 	}
// }

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal==:", s)
}
