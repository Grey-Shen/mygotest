package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

func main() {
	var contents []byte
	f, err := os.Open("./filter.txt")
	if err != nil {
		log.Printf("Failed to open faile: %s", err)
		return
	}

	fileStat, err := f.Stat()
	if err != nil {
		return
	}
	contents = make([]byte, fileStat.Size())

	if _, err := f.Read(contents); err != nil {
		if err == io.EOF {
			log.Println("done")
		} else {
			log.Printf("Failed to read file: %s", err)
			return
		}
	}

	filterArray := strings.Split(string(contents), "\n")
	fmt.Println(filterArray)
	var input string
	ff := bufio.NewReader(os.Stdin)

	fmt.Println("Please input a string:")
	input, _ = ff.ReadString('\n')

	for _, tmp := range filterArray {
		fmt.Println("tmptest", tmp)
		rege, err := regexp.Compile(tmp)
		if err != nil {
			return
		}

		myrep := strings.Repeat("*", utf8.RuneCountInString(tmp))
		fmt.Println("myrep", myrep)
		input = rege.ReplaceAllString(input, myrep)
	}

	fmt.Println("input", input)
}

func readLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
