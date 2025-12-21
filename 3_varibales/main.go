package main

import "fmt"

func main() {
	//string variables
	var name string = "Gopher"

	var name2 = "Go Programmer" // type inference: You can directly assign a value without specifying the type
	// In go the created variable must be used otherwise it will throw a compile time error
	fmt.Println("Hello, " + name + " " + name2)

	// boolean variables
	var isGopher bool = true // boolean variabe declaration with bool
	fmt.Println("Is Gopher a programming language mascot? ", isGopher)

	// integer variables
	var year int = 2009 // integer variable declaration with int
	fmt.Println("The Go programming language was released in the year ", year)

	// floating point variables
	// there are two types of floating point in go: float32 and float64
	var pi float64 = 3.14159 // floating point variable declaration with float64 
	fmt.Println("The value of pi is ", pi)


	// short variable declaration and assignment
	// can be used for int , float, string, bool and other types too
	name3 := "Go Enthusiast" // using := for variable declaration and assignment
	fmt.Println("Welcome, " + name3)
}
