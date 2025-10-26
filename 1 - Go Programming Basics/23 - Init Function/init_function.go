package main

import "fmt"

// init function is a special function in Go that is automatically executed when the package is initialized.
// Use: The init function is used when you need to perform setup tasks, variables, some state before the main

// Syntax:
// func init() {
//     // initialization code here
// }

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {
	fmt.Println("Inside main function.")
}
