package main

import "fmt"

// Maps are unordered collections of key-value pairs. They are used to store data in a way that allows for efficient retrieval based on keys.
// Maps are similar to dictionaries in other programming languages.
// Use: When you need to associate values with unique keys and perform fast lookups, insertions, and deletions.

// Syntax to create a map:
// var mapName map[keyDataType]valueDataType
// or
// mapName := make(map[keyDataType]valueDataType)	Using make function
// or
// mapName := map[keyDataType]valueDataType{key1: value1, key2: value2}		Using map literals

func main() {

	myMap := make(map[string]int) // Creating a map using make function
	fmt.Println(myMap)

	myMap["Alice"] = 25 // Adding key-value pairs to the map
	myMap["Bob"] = 30
	myMap["Charlie"] = 35
	fmt.Println("Map after adding key-value pairs:", myMap)

	// Updating value for an existing key
	myMap["Alice"] = 26
	fmt.Println("Map after updating Alice's age:", myMap)

	// Deleting a key-value pair from the map
	delete(myMap, "Bob")
	fmt.Println("Map after deleting Bob:", myMap)

}
