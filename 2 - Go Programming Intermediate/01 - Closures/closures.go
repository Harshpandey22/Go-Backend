package main

import "fmt"

// Closures in Go are function that references variables defined outside of them. They can capture and remember the environment in which they were created.
// Use: Closures are useful for creating function factories, maintaining state between function calls, and implementing callbacks.

// Syntax:
// func outerFunction() func() {
//     // outer function code
//     return func() {
//         // inner function code
//     }
// }

func main() {
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	// Creating new closure in which i is re-initialized to 0
	sequence2 := adder()
	fmt.Println(sequence2())
}

// adder() is a closure that captures the variable i from its surrounding environment.
// Each time the returned function is called, it increments and returns the value of i.
func adder() func() int {
	i := 0
	fmt.Println("previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
