package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Hello"

	fmt.Print(fruitList, "\n")
	fmt.Print(len(fruitList), "\n")

	var vegetables = [3]string{"Potato", "Tomato", "mushrooms"}

	fmt.Print(vegetables, "\n")
	fmt.Print(len(vegetables), "\n")
}
