package main

import "fmt"

// Functions in Go can return multiple values. This is particularly useful for returning both a result and an error value.
// Use: When you need to return more than one piece of information from a function, such as a computed value and a status indicator.

// Syntax:
// func functionName(parameter1 type1, parameter2 type2) (returnType1, returnType2) {
//     // function body
//     return value1, value2
// }

func main() {
	q, r := divide(10, 3)
	fmt.Println(q, r)
}

// Function with multiple return values: returns both quotient and remainder
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
