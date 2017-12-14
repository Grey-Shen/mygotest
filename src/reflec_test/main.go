package main

import (
	"fmt"
	"reflect"
)

type TestHandler interface {
	printinfo()
}

type Ball struct {
	color string `species:"gopher" color:"blue"`
	count int
}

func (b *Ball) printinfo() {
	fmt.Println(b.color, b.color)
}

type TaskInfo struct {
	TaskId string
	Name   string
}

func (t *TaskInfo) printinfo() {
	fmt.Println(t.TaskId, t.Name)
}

func (ball Ball) getColor() string {
	return ball.color
}

func main() {

	basket := Ball{color: "shend", count: 25}
	st := reflect.TypeOf(basket)
	algin := st.Align()
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
}

func interfacetest(b TestHandler) {
	st := reflect.TypeOf(b).Name()
	fmt.Println(st)
	b.printinfo()
}

// func FindColor(i interface{}) {
// 	st := reflect.TypeOf(i)
// 	s, b := st.FieldByName("TaskId")

// }
