package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://localhost:3000/data?varless=true&coursename=hahahha"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["varless"])
	for i, val := range qparams {
		fmt.Println(i, val)
	}
	fmt.Printf("Type of qparams: %T\n", qparams)
	//always pass reference
	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "localhost:3000",
		Path:     "data",
		RawQuery: "user=Rydham",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
