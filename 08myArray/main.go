package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Pear"

	fmt.Println("fruitList: ", fruitList)
	fmt.Println("fruitList: ", len(fruitList))

	var vegList = [3]string { "potato", "carrot", "cabbage"}
	fmt.Println("vegList: ", vegList)

}
