package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print("Slices we are starting\n")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}

	// fmt.Print("Type is", fruitlist, "\n")
	// fmt.Printf("Type is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")

	// fmt.Print("Type is", fruitlist, "\n")
	// fmt.Printf("Type is %T\n", fruitlist)

	fruitlist = append(fruitlist[:3])
	// fmt.Print("Type is", fruitlist, "\n")
	// fmt.Printf("Type is %T\n", fruitlist)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 456
	highScore[3] = 567
	// highScore[4] = 5123

	highScore = append(highScore, 444, 1213, 21312)

	// fmt.Print(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)

	// fmt.Print(highScore)
	// fmt.Print(sort.IntsAreSorted(highScore))

	var courses = []string{"Reacjs", "Js", "Swift", "python", "Ruby"}
	fmt.Print(courses, "\n")

	//to remove only one index in
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
