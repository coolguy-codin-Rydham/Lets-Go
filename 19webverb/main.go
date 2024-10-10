package main

import (
	"fmt"
	"io"
	"net/http"
)

const myUrl = "http://localhost:8000/get"

func PerformGetRequest(myUrl string) {

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)

	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	PerformGetRequest(myUrl)
}
