package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch in Go Long")

	diceNumber := rand.Intn(6) + 1
	fmt.Printf("dice number is %v: \n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Move two spot")
	case 3:
		fmt.Println("Move three spot")
		fallthrough // will execute the next case instead of breaking the switch
	case 4:
		fmt.Println("Move four spot")
		fallthrough
	case 5:
		fmt.Println("Move five spot")
	case 6:
		fmt.Println("Move six spot and roll dice again")
	default:
		fmt.Println("Dice probably flew off")
	}

}
