package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["jv"] = "Java"
	languages["py"] = "Python"

	fmt.Println("languages: ", languages)
	fmt.Println("JS", languages["js"])

	delete(languages, "js")
	fmt.Println("languages: ", languages)

	// looping map
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
