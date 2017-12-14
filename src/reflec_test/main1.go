package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Address []string
}

func main() {
	var a, b []Person
	a = []Person{
		Person{
			Name:    "a",
			Address: []string{"A", "B"},
		},
		Person{
			Name:    "b",
			Address: []string{"C", "D"},
		},
	}
	b = []Person{
		Person{},
		// Person{},
	}
	// var a, b []string
	// a = []string{"a", "b"}
	// b = []string{"", ""}

	reflect.Copy(reflect.ValueOf(b), reflect.ValueOf(a))
	fmt.Printf("a.Address: %p, b.address: %p\n", a[0].Address, b[0].Address)
	fmt.Printf("%#v\n", b)
	a[0].Address = []string{"E", "F"}
	fmt.Printf("a.Address: %p, b.address: %p\n", a[0].Address, b[0].Address)
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("b: %#v\n", b)

}
