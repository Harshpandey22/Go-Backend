package main

import "fmt"

// Arithmetic Operators in Go
// + : Addition
// - : Subtraction
// * : Multiplication
// / : Division
// % : Modulus (Remainder)
// ++ : Increment (Increases a variable's value by 1)
// -- : Decrement (Decreases a variable's value by 1)

// Follows the incedence rules of BODMAS (Brackets, Orders (i.e. powers and square roots, etc.), Division and Multiplication, Addition and Subtraction)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Modulus: ", result)
}
