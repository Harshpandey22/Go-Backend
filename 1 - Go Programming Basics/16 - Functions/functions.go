package main

import "fmt"

// Functions in Go are defined using the 'func' keyword. They can take parameters and return values.
// First class citizens: Functions can be assigned to variables, passed as arguments, and returned from other functions.
// USe: When you want to encapsulate reusable logic, perform operations, or create higher-order functions.

// Syntax:
// func functionName(parameter1 type1, parameter2 type2) returnType {
//     // function body
//     return value
// }

func main() {

	add(1, 2)
	fmt.Println(add(1, 2))

	// Function without explicit giving a name (anonymous function)
	func() {
		fmt.Println("Hello anonymous function!")
	}()

	// Assigning a variable to an anonymous function
	greet := func() {
		fmt.Println("Hello!")
	}
	greet()

	// Passing a function as an argument to another function
	result := applyOperation(3, 4, add)
	fmt.Println("3 + 4 = ", result)

	// Returning a function from another function
	multiplier := createMultiplier(3)
	fmt.Println("5 * 3 = ", multiplier(5))
}

// Simple function that adds two integers and returns the result
func add(a int, b int) int {
	return a + b
}

// Function that takes another function as a parameter
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
