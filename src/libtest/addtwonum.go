package main

import (
	"fmt"
	"math"
)

// type myNode struct {
// 	first  list.List
// 	second list.list
// }

// func main() {
// }

// func (m *myNode) New() {
// 	m.first = list.New()
// 	m.second = list.New()
// 	return m
// }

// func (m *myNode) initFirst() {

// }

// func addTwoNums(l1 *list.List, l2 *list.List) *list.List {
//     var (
//         e1 = l1.Front()
//         e2 = l2.Front()
//         )
// 	for ; e1 != nil; e1 = e1.Next() {
// 		tmp := e1.value.(int) + e2.Value.(int)
//         if e2.next()
// 	}
// }

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	l := list.New()
// 	p1 := person{name: "jack", age: 11}
// 	p2 := person{name: "lucy", age: 22}
// 	p3 := person{name: "john", age: 24}
// 	l.PushBack(p1)
// 	l.PushBack(p2)
// 	l.PushBack(p3)
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Println("name :", e.Value.(person).name)
// 		fmt.Println("age :", e.Value.(person).age)
// 	}
// }

func main() {
	// s := "abcd"
	// fmt.Println(strconv.QuoteToASCII(s))
	// a := s[1]
	// b := []string{"a", "b"}
	// fmt.Println(b[0])
	// fmt.Println(strconv.QuoteRuneToASCII(rune(a)))
	// fmt.Println(strconv.QuoteRuneToASCII(s[1]))
	a := 128
	b := 129
	fmt.Println(math.Max(float64(a), float64(b)))

}
