package main

import "fmt"

func main() {
	fmt.Print("Maps in Golang\n")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Print("Js is short for ", languages["JS"], "\n")

	// delete(languages, "RB")
	fmt.Print(languages, "\n")

	//Loops

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}
