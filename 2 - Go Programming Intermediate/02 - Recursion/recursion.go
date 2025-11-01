package main

// Recursion in Go is a programming technique where a function calls itself to solve a problem.
// It is commonly used for problems that can be broken down into smaller, similar subproblems.
// Use: Recursion is useful for tasks such as traversing data structures (like trees), solving mathematical problems (like factorials and Fibonacci numbers), and implementing algorithms (like quicksort and mergesort).

// Syntax:
// func recursiveFunction(parameters) returnType {
//     // base case
//     if baseCondition {
//         return baseValue
//     }
//     // recursive case
//     return recursiveFunction(modifiedParameters)
// }

func main() {
	n := 5
	result := factorial(n)
	println("Factorial of", n, "is", result)
}

// factorial calculates the factorial of a non-negative integer n using recursion.
func factorial(n int) int {
	// Base case: factorial of 0 or 1 is 1
	if n == 0 || n == 1 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)
}
