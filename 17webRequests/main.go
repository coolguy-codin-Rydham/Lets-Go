package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.centralspaceofelites.com/"

// const url = "http://localhost:3000/data"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("hellllllllo")
		panic("")
	}
	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println(response.Body)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	data := string(dataBytes)

	fmt.Println(data)

}
