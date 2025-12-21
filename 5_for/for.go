package main

import "fmt"

// for -> only used for looping in go
func main() {
	// while loop example using for
	i := 1
	for i <= 5 {
		fmt.Println("While Loop iteration: ", i)
		i++
	}

	// infinate loop example using for

	// for {
	// 	fmt.Println("Infinate Loop: Press Ctrl + C to stop")
	// }

	//classic for loop example
	for a := 1; a <= 5; a++ {
		if a == 2 {
			continue // skip the iteration when a is 2
		}
		fmt.Println("Classic For Loop iteration: ", a)
	}

	// range based for loop example
	for i := range 5 {
		fmt.Println("Range based For Loop iteration: ", i)
	}
}
