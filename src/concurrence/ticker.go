package main

import (
	"fmt"
	"time"
)

func main() {
	var ticker = time.NewTicker(3 * time.Second)
	ticks := ticker.C
	go func() {
	      		for _ = range ticks {
			_, ok := <-ticks
		if ok {
				fmt.Println(ticker.C)
			}
		}
	}()
}
func name() {

}
