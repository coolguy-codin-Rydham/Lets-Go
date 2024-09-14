package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTIme := time.Now()
	fmt.Println(presentTIme)

	fmt.Println(presentTIme.Format("01.02.2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.August, 10, 22, 22, 12, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02.01.2006 Monday"))
}
