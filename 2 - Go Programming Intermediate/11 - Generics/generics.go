package main

import "fmt"

// Generics in Go allow us to write functions and data structures that work with different types while maintaining type safety, eliminating the need to write redundant code for different types.
// Use: The primary use of generics in Go is to enable type-safe code reuse by writing functions and data structures that can operate on values of different types without sacrificing type checking or using the less safe any (empty interface).

// Generic in function
func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {

	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	x1, y1 := "John", "Jane"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)
}
