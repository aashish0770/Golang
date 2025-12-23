package main

import (
	"fmt"
	"slices"
)

// slices are dynamic in size
// more flexible than arrays
// most commonly used data structure in go
//useful methods to work with slices

func main() {

	// declare a slice of integers
	var nums []int

	// check length and capacity of slice
	fmt.Println("Length:", len(nums), "Capacity:", cap(nums))

	var num1 = make([]int, 5) // length 5, capacity 5 for initialization size

	fmt.Println("Length:", len(num1), "Capacity:", cap(num1))

	// array with fixed capacity
	num2 := make([]int, 3, 5) // length 3, capacity 5
	fmt.Println("Length:", len(num2), "Capacity:", cap(num2))

	// append elements to slice
	num2 = append(num2, 10)
	num2 = append(num2, 20)
	num2 = append(num2, 30) // exceeds initial capacity of 5, so capacity will double

	fmt.Println("After appending value:", num2)
	fmt.Println("Length:", len(num2), "Capacity:", cap(num2))

	// another way to declare and initialize a slice
	num3 := []int{}
	fmt.Println("Length:", len(num3), "Capacity:", cap(num3))

	// copy slice function
	num4 := make([]int, 0, 5)
	num4 = append(num4, 3)
	var num5 = make([]int, len(num4))
	fmt.Println("Before copy num5:", num5)

	// first give the destination slice then source slice
	copy(num5, num4) // copy num4 to num5
	fmt.Println("After copy num5:", num5)

	// slicing operator
	var num6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original slice num6:", num6[0:4])           // from index 0 to 4 (excluding 4)
	fmt.Println("Original slice num6 to end:", num6[3:])     // from index 3 to end
	fmt.Println("Original slice num6 to index 6:", num6[:6]) // from start to index 6 (excluding 6)

	// slice package
	var num7 = []int{5, 2, 6, 3, 1, 4}
	var num8 = []int{7, 8, 9}

	fmt.Println("compare slice it return true or false", slices.Equal(num7, num8))

	// 2d slice
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2d slice", matrix)
}
