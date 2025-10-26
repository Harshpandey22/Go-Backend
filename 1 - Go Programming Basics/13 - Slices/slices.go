package main

import (
	"fmt"
	"slices"
)

// Slices are dynamically-sized, flexible views into the elements of an array.They are more powerful than arrays and are used extensively in Go programming.
// Slices are similar to lists in other programming languages.
// Use: When you need a collection that can grow and shrink in size, and when you want to work with subsets of arrays.

// Syntax to create a slice:
// var sliceName []dataType

func main() {

	// Creating a slice using make function
	arr1 := make([]int, 5) // creates a slice of integers with length 5
	fmt.Println("Slice created using make:", arr1)

	arr2 := []int{1, 2, 3, 4, 5}       // Creating a slice with initial values
	slice1 := append(arr2, 10, 11, 12) // Appending values to the slice
	fmt.Println("Slice after appending values:", slice1)

	// Copy function to copy one slice to another
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(sliceCopy)

	// Using slices.Equal to compare two slices
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}
}
