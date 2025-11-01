package main

import "fmt"

// Pointers are variables that store the memory address of another variable.
// Use: Pointers in Go are used for various purposes, including:
// 1. Efficiency: Passing large structs or arrays by pointer instead of by value to avoid copying data.
// 2. Mutability: Allowing functions to modify the original variable's value.
// 3. Using Pointers as arguments in functions is a real world usage of pointers.
// 4. Prevent memory leaks: By using pointers, we can manage memory more effectively and avoid unnecessary copies.

// Syntax:
// To declare a pointer, use the asterisk (*) before the type.
// To get the address of a variable, use the ampersand (&).
// To dereference a pointer (access the value it points to), use the asterisk (*) before the pointer variable.
// var x int = 42
// var p *int = &x      // p is a pointer to x
// fmt.Println(*p)      // Dereferencing p gives the value of x (42)

func main() {

	// Declaring a pointer to an integer
	var ptr *int

	var a int = 10

	ptr = &a // Assigning the address of 'a' to the pointer 'ptr', also known as referencing

	fmt.Println("Value of a:", a)                    // Output: 10
	fmt.Println("Address of a:", &a)                 // Output: Address of a
	fmt.Println("Value of ptr (address of a):", ptr) // Output: Address of a
	fmt.Println("Value pointed to by ptr:", *ptr)    // Output: 10m, dereferencing ptr to get the value of a

	// Modifying the value of 'a' using the pointer 'ptr'. The actual value of variable 'a' is changed using pointer not copying the value.
	modifyValue(ptr)
	fmt.Println(a)
}

// In Go we can pass pointers to functions to allow them to modify the original variable's value.
func modifyValue(ptr *int) {
	*ptr++
}
