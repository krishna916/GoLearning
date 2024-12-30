package main

import "fmt"

const LoginToken string = "this is constant" // public

func main() {
	var username string = "krishna916"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455344343433453454
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicait type
	var website = "this is string"
	fmt.Println(website);
	fmt.Printf("variable is of type: %T \n", website)

	// no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("login token is of type: %T \n", LoginToken)
}
