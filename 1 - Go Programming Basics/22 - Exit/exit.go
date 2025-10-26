package main

import (
	"fmt"
	"os"
)

// Exit in Go is a function from the "os" package that terminates the program immediately with a specified exit code.
// When Exit is called, it does not run any deferred functions, so any cleanup code in deferred functions will not be executed.
// Use: Exit is typically used when a program needs to terminate due to an error or when it has completed its task successfully.

// Syntax:
// os.Exit(code int)

func main() {
	fmt.Println("Starting the main function.")

	// Exit with a status code of 1 to indicate an error
	os.Exit(1)

	fmt.Println("End of the main function.") // This line will not execute
}
