package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println(index, day)
	}
	for _, day := range days {
		fmt.Println(day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 3 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++

	}

lco:
	fmt.Println("Jumping at ahahhahahah")

}
