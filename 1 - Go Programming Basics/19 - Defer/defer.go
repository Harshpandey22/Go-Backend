package main

import "fmt"

// In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.
// LIFO (Last In First Out) order is used to execute the deferred functions.
// Use: Defer is often used to clean up resources, unlock a mutex, logging & tracing, etc.

// Syntax:
// defer functionName(parameters)

func main() {
	process()
	lifoDefer()
}

// Defer example
func process() {
	defer fmt.Println("This is deferred and will execute last.")
	fmt.Println("This will execute first.")
}

func lifoDefer() {
	defer fmt.Println("First deferred call")
	defer fmt.Println("Second deferred call")
	defer fmt.Println("Third deferred call")
	fmt.Println("Function execution started")
}
