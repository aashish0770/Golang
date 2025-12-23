package main

import "fmt"

// range is used to iterate over a collection of values
func main() {
	nums := []int{10, 20, 30, 40, 50}

	// using for loop with index
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println("Index:", i, "Value:", nums[i])
	// }

	sum := 0

	// using range keyword to iterate over slice
	for index, num := range nums {
		fmt.Println("Index:", index, "Value:", num)

		sum += num
	}
	fmt.Println("Sum of all numbers:", sum)

	// using range to iterate over map
	m := map[string]string{
		"name": "Golang",
		"type": "Programming Language",
		"year": "2009",
	}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// using range to iterate over string
	for i, c := range "hello" {
		fmt.Println("Index:", i, "Character:", c) // c will print unicode code point
		fmt.Println("Character as string:", string(c)) // convert code point to string
	}
}
