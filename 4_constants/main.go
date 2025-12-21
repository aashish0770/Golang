package main

import "fmt"

const name3 = "you can create constants outside the main function" // constant declaration with type inference

func main() {

	//constant
	const name string = "Golang" // constant declaration with type string and value "Golang"

	const name1 = "Go Language" // constant declaration with type inference

	fmt.Println("Hello, " + name)
	fmt.Println("Welcome to " + name1)
	fmt.Println(name3)

	//multiple constant booking example
	const (
		port     = 8080
		host     = "localhost"
		isActive = true
	)

	fmt.Println(port, host)

}
	