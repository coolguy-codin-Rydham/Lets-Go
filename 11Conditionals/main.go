package main

import (
	"fmt"
)

func main() {
	fmt.Print("If else in golang\n")
	// loginCount := 19

	// reader := bufio.NewReader(os.Stdin)

	// loginC, _ := reader.ReadString('\n')

	// loginC = strings.TrimSpace(loginC)

	// loginCInt, _ := strconv.ParseInt(loginC, 10, 64)

	// fmt.Println(loginC)
	// fmt.Println("Converted value is :", loginCInt)

	// var result string
	// if loginCInt < 10 {
	// 	result = "Regular User"
	// } else if loginCInt > 10 && loginCInt < 20 {
	// 	result = "Ehh Users"
	// } else {
	// 	result = "Supreme User"
	// }

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less that 10")
	}

	// fmt.Println(result)

}
