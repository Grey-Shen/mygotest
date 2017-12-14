package main

import (
	"fmt"
	"log"
	"reflect"
)

type TestHandler interface {
	printinfo()
}

type Ball struct {
	Color string `species:"gopher" color:"blue"`
	Count int
}

func (b *Ball) printinfo() {
	fmt.Println(b.Color, b.Color)
}

type TaskInfo struct {
	TaskId string
	Name   string
}

func (t *TaskInfo) printinfo() {
	fmt.Println(t.TaskId, t.Name)
}

func (ball Ball) getColor() string {
	return ball.Color
}

func main() {

	basket := Ball{Color: "shend", Count: 25}
	st := reflect.TypeOf(basket)
	algin := st.Align()
	fmt.Printf("========= aligin: %d\n", algin)
	fmt.Println("NumMethod", st.NumMethod())
	fmt.Println("NumField", st.NumField())
	ss, exists := st.FieldByName("color")
	fmt.Println("--", ss, exists)
	fmt.Println("name", st.Name())
	fmt.Println("algin", algin)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	task := TaskInfo{TaskId: "haha", Name: "heeh"}
	v := reflect.ValueOf(task)
	fmt.Println("testvalue", v.FieldByName("Name").Interface().(string))

	interfacetest(&basket)
	interfacetest(&task)

	var bt = Ball{}

	SetData(bt)
	SetData2(&bt)

	fmt.Printf("Bt is %v\n", bt)

	type myfloat float64
	var x myfloat = 3.4
	fmt.Println("type: ", reflect.TypeOf(x), reflect.ValueOf(x).Kind())
	fmt.Println("value:", reflect.ValueOf(x).String())
}

func interfacetest(b TestHandler) {
	st := reflect.TypeOf(b).Name()
	fmt.Println(st)
	b.printinfo()
}

func SetData(data interface{}) {
	fmt.Println("=========== SetData")
	basket := Ball{Color: "red", Count: 25}
	v := reflect.ValueOf(data).Elem()
	if !v.CanAddr() {
		fmt.Println("========= cannot addr")
		return
	}
	if !v.CanSet() {
		log.Println("Can not set")
		return
	}

	v.Elem().Set(reflect.ValueOf(&basket).Elem())
}

func SetData2(data interface{}) {
	fmt.Println("=========== SetData2")
	basket := Ball{Color: "red", Count: 25}
	tmp := data.(*Ball)
	v := reflect.ValueOf(tmp.Color)
	if !v.CanAddr() {
		fmt.Println("========= cannot addr")
		return
	}
	if !v.CanSet() {
		log.Println("Can not set")
		return
	}

	v.Elem().Set(reflect.ValueOf(&basket).Elem())
}
