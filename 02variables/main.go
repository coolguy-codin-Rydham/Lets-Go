package main

import "fmt"

// Any variable made with its first letter as capital is a public variabe
const LoginToken string = "akjfhlashflkhgf"

func main() {

	var username string = "Rydhampreet"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float64 = 255.098765476545678765
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// Defauolt values

	var aVaraible int
	fmt.Println(aVaraible)
	fmt.Printf("Variable is of type: %T\n", aVaraible)

	//implicit way to declare varables

	var website = "rydhampreetsingh.vercel.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	//No var style
	//Not allowed outside the functions
	numOfUser := 300000
	fmt.Println(numOfUser)

	fmt.Print(LoginToken)

}
