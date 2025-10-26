package main

import "fmt"

// Variadic functions in Go can accept a variable number of arguments. This is useful when you don't know beforehand how many arguments will be passed to the function.
// Variadic parameters are defined using an ellipsis (...) before the type.
// Use: When you want to create functions that can handle a flexible number of inputs, such as summing numbers, concatenating strings, etc.

// Syntax:
// func functionName(parameter1 type1, parameter2 ...type2) returnType {
//     // function body
//     return value
// }

func main() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))
}

// Variadic function that takes a variable number of integers and returns their sum
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
