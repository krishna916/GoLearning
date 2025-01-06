package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := add(3, 5)

	fmt.Println("Result is: ", result)

	proRes, message := proAdd(3, 5, 2)

	fmt.Println("Result of proAdd:, ", message, proRes)

}

func add(a int, b int) int {
	return a + b
}

func proAdd(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Pro function example with second result"
}

func greeter() {
	fmt.Println("hello world from a function")
}
