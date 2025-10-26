package main

// Constants in Go are immutable values that are defined at compile time and cannot be changed during the execution of the program.
// Use: When you have values that should remain constant throughout the program, such as mathematical constants, configuration values, etc.

// Syntax:
// const constantName type = value

import "fmt"

// Lower SnakeCase for constants that can be changed
const pi = 3.14
const gravity = 9.8

// UPPER_SNAKE_CASE for constants that are immutable
const PI = 3.14
const GRAVITY = 9.8

func main() {

	fmt.Println(pi)
	fmt.Println(gravity)
	fmt.Println(PI)
	fmt.Println(GRAVITY)

}
