package main

import "fmt"

type structtest struct {
	a int
	b int
}

func (s *structtest) SetA(value int) {
	s.a = value
}

func main() {
	t := structtest{
		a: 1,
		b: 2,
	}

	t.SetA(5)
	fmt.Println("testtest", t)
}
