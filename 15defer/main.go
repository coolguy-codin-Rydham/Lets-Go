package main

import "fmt"

// defer keyword moves the execution of tbe line at the end in LIFO manner down the line earlier it is executes

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Hehehe")
	fmt.Println("Hello")
	myDefer()
}

// hello 4 3 2 1 0 heheheh world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
