package main

import (
	"encoding/json"
	"fmt"
)

var test = json.RawMessage(`{"category": "侦探","user": "毛利小五郎"}`)

func main() {
	var a tt
	json.Unmarshal(test, &a)
	fmt.Println("a:", a)

}
