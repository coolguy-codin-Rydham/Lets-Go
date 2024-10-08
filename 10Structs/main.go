package main

import "fmt"

func main() {
	//We dont have inheritance, no super no parent
	fmt.Print("Struct in golang\n")

	rydham := User{"Rydham", "rydhampreetsingh.gindra@gmail.com", false, 20}

	fmt.Printf("Type of rydham %T\n", rydham)
	fmt.Printf("\nDetails of Rydham are %+v\n\n", rydham) //Good one
	fmt.Print(rydham, "\n")
	fmt.Print(rydham.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
