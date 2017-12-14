package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
        {"Name": "Platypus"},
        {"Name": "Quoll",    "order": false}
    ]`)
	type Animal struct {
		Name  string `json:"name"`
		Order bool   `json:"order"`
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
