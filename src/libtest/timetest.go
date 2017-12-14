package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	var count int
// 	for {
// 		select {
// 		case <-ticker.C:
// 			fmt.Println("ticker..")
// 			count++
// 		default:
// 		}

// 		if count == 10 {
// 			ticker.Stop()
// 			fmt.Println("Stop the ticker")
// 			count++
// 		}
// 	}

// }

func main() {
	myTimer := time.NewTimer(time.Second * 3)
	exitTimer := time.NewTimer(time.Second * 10)
	var count int
LOOP:
	for {

		select {
		case <-myTimer.C:
			fmt.Println("Timer..")
			count++
			if ok := myTimer.Reset(time.Second); ok {
				fmt.Println("Can not reset myTimer")
			}
		case <-exitTimer.C:
			fmt.Println("Exit...")
			break LOOP
		default:
		}

		if count == 5 {
			myTimer.Stop()
			fmt.Println("Stop the timer")
			count++
		}
	}

}
