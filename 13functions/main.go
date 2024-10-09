package main

import "fmt"

func foo() {
	fmt.Println("Hello From Functions")
}
func greeterTwo() {
	fmt.Println("Anoter method")
}

// func sumOfTwo(valOne int, valTwo int) (int, int, string) {
// 	return (valOne + valTwo), (valOne - valTwo), "Hey I jus calced"
// }

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func main() {
	fmt.Println("Functions in golang")
	foo()
	// summ, diff, sting := sumOfTwo(10, 21)
	// fmt.Println(summ, diff, sting)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println(proRes)

	greeterTwo()
}
