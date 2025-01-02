package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Pear", "Melon"}
	fmt.Printf("type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Papaya")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 934
	highScores[2] = 434
	highScores[3] = 834
	// highScores[4] = 777

	highScores = append(highScores, 555, 333, 321)

	fmt.Println(highScores)

	// sort.Ints(highScores)
	slices.Sort(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove value for slices based on index
	var courses = []string{ "react", "java", "javascript", "python", "Typescript"};
	fmt.Println(courses)
	var index int = 2;

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses);
}
