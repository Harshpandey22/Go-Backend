package main

import "fmt"

// Panic in Go is a built-in function that stops the normal execution of a goroutine.
// When a function panics, it immediately stops executing and begins to unwind the stack, running any deferred functions along the way.
// Panics are typically used for unrecoverable errors, such as accessing an out-of-bounds index in a slice or dereferencing a nil pointer.
// Use: Panic is often used in situations where the program cannot continue to run safely.

// Syntax:
// panic(value interface{})

func main() {
	process(10)
	process(-3)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be non-negative")
	}
	fmt.Println("Processing:", input)
}
