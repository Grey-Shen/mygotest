package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println(ListAllDirNames("/tmp/data"))
}

func ListAllDirNames(path string) []string {
	var dirNames []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("dirtest", file.Name())
			dirNames = append(dirNames, file.Name())
		}
	}
	return dirNames
}
