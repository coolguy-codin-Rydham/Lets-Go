package main

import "fmt"

func main() {
	fmt.Printf("Hello\n")
	// var one = 2

	// var ptr *int
	// fmt.Print(ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println(*ptr)
	fmt.Printf("%T", myNumber)
}
