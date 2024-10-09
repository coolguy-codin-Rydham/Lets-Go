package main

import "fmt"

//The things with no capital first letter are non exportable

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

//this is a method here

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println(u)
}

func main() {
	rydham := User{"Rydham", "rydham@gmail.com", false, 20, 10}
	fmt.Println(rydham)
	rydham.GetStatus()
	rydham.NewMail()
	fmt.Println(rydham)

}
