package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statement example
	// switch is used to select one of many code blocks to be executed
	// no need to use break statement in go switch case
	i := 5

	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2 or 3")
	}

	// switch with multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch example
	// type switch is used to compare the type of an interface variable
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("hello")

}
