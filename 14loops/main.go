package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Saturday"}

	fmt.Println(days)

	// standard loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, value)
	// }

	randomValue := 1

	for randomValue < 10 {

		if randomValue == 2 {
			goto lco
		}

		if randomValue == 5 {
			randomValue++
			continue
		}

		fmt.Println("Value is: ", randomValue)
		randomValue++
	}

	lco:
		fmt.Println("Jumping to GOTO!")
}
