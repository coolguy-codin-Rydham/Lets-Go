package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome User\n"
	fmt.Print(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza: ")

	// Comma Ok syntax || err ok syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Reading: ", input)

}
