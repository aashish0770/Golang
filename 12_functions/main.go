package main

import "fmt"

// function to add two integers and return the sum
func add(a, b int) int {
	return a + b
}

func getLanguage() (string, string, string, bool) {
	return "Go", "Golang", "Go Programming Language", true
}

func processIt(fn func(a int) int) {
	fn(10)
}

func main() {

	// calling the add function
	result := add(2, 5)
	fmt.Println("Sum:", result)

	// calling the getLanguage function
	lang1, lang2, lang3, isGo := getLanguage()
	fmt.Println("Languages:", lang1, lang2, lang3, isGo)

	// passing an anonymous function as an argument
	fn := func(a int) int {
		return 2
	}

	processIt(fn)
}
