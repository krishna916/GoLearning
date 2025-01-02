package main

import "fmt"

func main() {
	fmt.Println("Structs")

	// no inheritence in go
	krish := User{"Krish", "krish@mytech.dev", true, 30}
	fmt.Println(krish)
	fmt.Printf("krish's details are: %+v\n", krish)
	fmt.Printf("Name is %v and email is %v\n", krish.Name, krish.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
