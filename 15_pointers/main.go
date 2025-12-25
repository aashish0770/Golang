package main

import "fmt"

// func changeNum(num int) {
// 	num = 20
// 	fmt.Println("Before changeNum:", num)
// }

// passing the value by reference using pointers
func changeNum(num *int) {
	*num = 5
	fmt.Println("Inside changeNum:", *num)
}

func main() {
	num := 10
	// changeNum(num)
	changeNum(&num) // passing the memory address of num
	fmt.Println("Memory address of num:", &num)
	fmt.Println("After changeNum:", num)
}
