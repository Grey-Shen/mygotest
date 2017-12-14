package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type T struct {
	A int `yaml:"a,omitempty"`
	B int `yaml:"b"`
}

func main() {
	var t T
	buf, err := ioutil.ReadFile("/Users/chenyajun/Documents/goproject/src/download/test.yml")

	if err != nil {
		return
	}

	fmt.Println("test is ", string(buf))
	yaml.Unmarshal(buf, &t)
	fmt.Println("t.A:", t.A, "t.B: ", t.B)
}
