package main

import (
	"fmt"
	"io"
	"os"
)

// commons

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// utils

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(dataByte))
}

func main() {
	content := "This needs to go in a file"
	file, err := os.Create("./test1.txt")
	checkNilErr(err)
	length, err2 := io.WriteString(file, content)
	checkNilErr(err2)
	fmt.Printf("length: %d\n", length)
	readFile("./test1.txt")
	file.Close()
}
