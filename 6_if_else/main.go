package main

import "fmt"

func main() {

	age := 18

	// if else example

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is a minor")
	}

	// if else if else example

	if age < 13 {
		fmt.Println("Person is a child")
	} else if age >= 13 && age < 18 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is an adult")
	}

	// logical operators with if else

	var role string = "admin"
	var hasAccess bool = false

	if role == "admin" && hasAccess {
		fmt.Println("Admin has access to the system")
	} else {
		fmt.Println("Access denied")
	}

	// we can also declare and initialize a variable in the if statement
	if age := 14; age < 18 {
		fmt.Println("Person is a minor")
	} else if age >= 10 {
		fmt.Println("Person is a child")
	} else {
		fmt.Println("Person is an adult")
	}

	// go does not have ternary operators like other languages
	// but we can achieve similar functionality using if else if else
}
