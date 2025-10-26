package main

import "fmt"

// Recover is a built-in function in Go that allows a program to regain control of a panicking goroutine.
// Recover can be useful for handling errors gracefully and allowing the program to continue running.
// Use: Recover is typically used in situations where you want to handle panics and recover from them without terminating the program.

// Syntax:
// recover() interface{}

func main() {
	process()
	fmt.Println("Returned from process.")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("This is deferred and will execute last.")
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process") // This line will not execute
}
