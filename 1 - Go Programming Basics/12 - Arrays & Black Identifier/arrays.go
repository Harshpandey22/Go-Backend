package main

import "fmt"

// Arrays are used to store multiple values of the same type in a single variable. Arrays in Go have a fixed size
// Syntax: var arrayName [size]dataType
// Use: When you need to store a collection of items of the same type and the size of the collection is known and fixed.

// Blank Identifier (_) can be used to ignore values in Go. Underscore (_) is known as the blank identifier in Go and is used to store unused values.
// Example: _, value := someFunction() // Ignore the first return value
// In loops: for _, value := range someSlice { // Ignore the index

func main() {

	// Declare an array
	var arr [5]int // Declare an array of integers with size 5
	// Assign values to the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	// Access and print array elements
	for i := 0; i < len(arr); i++ {
		println("Element at index", i, "is", arr[i])
	}

	// Declare and initialize an array in a single line
	names := [3]string{"Alice", "Bob", "Charlie"}
	for i, name := range names {
		println("Name at index", i, "is", name)
	}

	// Multi-dimensional array
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			println("Element at matrix[", i, "][", j, "] is", matrix[i][j])
		}
	}

	// Array is non-reference type, copying an array creates a new array, modifying one does not affect the other.
	originalArray := [4]int{1, 2, 3, 4}
	copyArray := originalArray // Copying the array
	copyArray[0] = 10          // Modifying the copied array

	fmt.Println("Original array: ", originalArray)
	fmt.Println("Copied array: ", copyArray)

	// Using range to iterate over an array
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Using blank identifier to ignore a value
	// a, b := someFunction()
	// fmt.Println(a)
	// fmt.Println(b)

	// Using blank identifier to ignore a value
	a, _ := someFunction()
	fmt.Println(a)

	// Using blank identifier in a loop
	numbers := []int{10, 20, 30, 40, 50}
	for _, num := range numbers {
		fmt.Println("Number:", num)
	}

}

func someFunction() (int, int) {
	return 1, 2
}
