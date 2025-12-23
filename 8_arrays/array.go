package main

import "fmt"

func main() {

	// arrays number of elements are fixed in size

	var nums [5]int // declare an array of integers with size 5

	//array length
	fmt.Println(len(nums))

	// assign values to array elements
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30

	fmt.Println("This access only one array element", nums[2]) // access array element at index 2
	fmt.Println("This is the complete array", nums)            // print the complete array

	// array for boolean values
	var vals [4]bool
	vals[2] = true
	fmt.Println("Boolean array:", vals)

	// array for strings
	var strs [3]string
	strs[0] = "Hello"
	fmt.Println("String array:", strs)

	// declare and initialize an array in one line
	nums1 := [3]int{1, 2, 3}
	fmt.Println("Initialized array:", nums1)

	// 2d array
	nums2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("2d array:", nums2)

	// fixed size, that is predictable
	// memory allocation is easier
	// contatnt time access to elements


}
