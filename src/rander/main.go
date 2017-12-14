package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	// vals := []int{10, 12, 14, 16, 18, 20}
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// for _, i := range r.Perm(len(vals)) {
	// 	val := vals[i]
	// 	fmt.Println(val)
	// }

	// Shuffle1(vals)
	// fmt.Println(vals)
	// vals1 := [3]int{3, 4, 5}
	// Shuffle1(vals1)
	// fmt.Println(vals1)
	// Shuffle2(vals1)
	// fmt.Println(vals1)

	taskIds := []string{"aaa", "bbb", "ccc"}
	fmt.Println("taskIdtest", shuffleTaskIds(taskIds))
}

func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]int, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func Shuffle1(slice interface{}) {
	rand.Seed(time.Now().UnixNano())
	rv := reflect.ValueOf(slice)
	if rv.Kind() == reflect.Slice {
		swap := reflect.Swapper(slice)
		length := rv.Len()
		for i := length - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			swap(i, j)
		}
	}
}

func shuffleTaskIds(taskIds []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(taskIds))
	perm := r.Perm(len(taskIds))
	for i, randIndex := range perm {
		ret[i] = taskIds[randIndex]
	}
	return ret
}

// func Shuffle2(s interface{}) {
// 	arr := reflect.ValueOf(s)
// 	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
// 		r := rand.New(rand.NewSource(time.Now().Unix()))
// 		for i, v := range r.Perm(arr.Len()) {
// 			arr[i] := vals[i]
// 		}
// 	}
// }
func Shuffle2(array []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

// func ToInterfaces(s interface{}) []interface{} {
// 	arr := reflect.ValueOf(s)

// 	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
// 		interfaces := make([]interface{}, arr.Len())
// 		for i := 0; i < arr.Len(); i++ {
// 			interfaces[i] = arr.Index(i).Interface()
// 		}
// 		return interfaces
// 	}
// 	return nil
// }
