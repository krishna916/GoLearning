package main

import "fmt"

func main() {
	fmt.Println("if else in GO")

	loginCount := 43

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 100 {
		result = "Power user"
	} else {
		result = "Admin User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less that 10")
	}

    
}
