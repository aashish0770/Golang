package main

import (
	"fmt"
	"maps"
)

// maps -> key value pairs
func main() {

	// declaring and initializing a map
	m := make(map[string]string)
	// adding key value pairs to the map
	m["name"] = "Golang"
	m["type"] = "Programming Language"
	m["year"] = "2009"

	// get an element from the map
	fmt.Println("Name:", m["name"])
	fmt.Println("Type:", m["type"])
	fmt.Println("Year:", m["year"])

	// what if we try to get a key that doesn't exist
	fmt.Println("Version:", m["version"]) // returns zero value which is empty string in this case

	// deleting a key value pair from the map
	delete(m, "year")
	fmt.Println("After deleting year, map m:", m)

	// clearing a map
	clear(m)

	// another way to declare and initialize a map
	m1 := map[string]string{"name": "Golang", "type": "Programming Language", "year": "2009"}
	fmt.Println("Map m1:", m1)

	// getting the length of the map
	fmt.Println("Length of map m1:", len(m1))

	// checking if a key exists in the map
	k, ok := m1["name"]
	fmt.Println("Value of key name:", k)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// map package
	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
	m3 := map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println(maps.Equal(m2, m3))
}
