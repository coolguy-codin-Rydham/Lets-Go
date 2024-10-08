package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// How to print
	welcome := "Welcome User\n"
	fmt.Println(welcome)

	//How to take user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Type %T\n", input) //how to check the type of variable data

	// Type conversions
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	//Error handling and how to check errors with comma ok syntax
	// Conditionals
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating)
		fmt.Printf("New Type: %T\n", numRating)
	}

}
