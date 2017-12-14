package main

import (
	"fmt"
	"reflect"
	"time"
)

// find a fermormance method for slice append
type Person struct {
	Name string
	Age  int
}

// func main() {
// 	var (
// 		persons = make([]*Person, 0, 1000000000)
// 		a       = make([]*Person, 0, 1000)
// 		a1      = make([]*Person, 1000)
// 	)

// 	fmt.Println("=========== perpare data ============")
// 	for i := 0; i < 1000000; i++ {
// 		persons = append(persons, &Person{
// 			Name: "Jack",
// 			Age:  i,
// 		})
// 	}
// 	fmt.Println("====== prepare ok")

// 	fmt.Println("======test1 start========")
// 	now := time.Now()
// 	for j := 0; j < 100; j++ {
// 		for _, p := range persons {
// 			a = append(a, p)
// 			if len(a) == cap(a) {
// 				a = a[0:0]
// 			}
// 		}
// 	}
// 	fmt.Println("====== spent ", time.Now().Sub(now))

// 	fmt.Println("======test2 start========")

// 	for j := 0; j < 100; j++ {
// 		i := 0
// 		for _, p := range persons {
// 			if a1[i] == nil {
// 				a1[i] = new(Person)
// 			}
// 			a1[i] = p
// 			i++
// 			if i == 1000 {
// 				i = 0
// 			}
// 		}
// 	}

// 	fmt.Println("====== spent ", time.Now().Sub(now))
// }

func copyInsert(slice interface{}, pos int, value interface{}) interface{} {
	v := reflect.ValueOf(slice)
	v = reflect.Append(v, reflect.ValueOf(value))
	reflect.Copy(v.Slice(pos+1, v.Len()), v.Slice(pos, v.Len()))
	v.Index(pos).Set(reflect.ValueOf(value))
	return v.Interface()
}

func Insert(slice interface{}, pos int, value interface{}) interface{} {

	v := reflect.ValueOf(slice)

	ne := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 1, 1)

	ne.Index(0).Set(reflect.ValueOf(value))
	v = reflect.AppendSlice(v.Slice(0, pos), reflect.AppendSlice(ne, v.Slice(pos, v.Len())))

	return v.Interface()
}
func main() {
	slice := []int{1, 2}
	slice2 := []int{1, 2}
	slice3 := []int{1, 2}
	slice4 := []int{1, 2}

	t0 := time.Now()
	for i := 1; i < 10000; i++ {
		slice = append(slice[:1], append([]int{i}, slice[1:]...)...)
	}
	t1 := time.Now()
	for i := 1; i < 10000; i++ {
		slice2 = Insert(slice2, 1, i).([]int)
	}

	t2 := time.Now()
	for i := 1; i < 10000; i++ {
		slice3 = copyInsert(slice3, 1, i).([]int)
		//  fmt.Println(slice3)
	}

	t3 := time.Now()

	//元素检测
	for i := 0; i < 10000; i++ {
		if slice[i] != slice2[i] || slice2[i] != slice3[i] {
			fmt.Println("error")
		}
	}

	fmt.Println("reflect append insert:", t2.Sub(t1), "append insert: ", t1.Sub(t0), "copy Insert: ", t3.Sub(t2))
}
