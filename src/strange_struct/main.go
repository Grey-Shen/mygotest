package main

import "fmt"

// strange struct

func main() {
	a := &struct{}{}
	b := &struct{}{}
	fmt.Println(a == b)
	fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
}
