package main // main package is the entry point of the go program

import "fmt"

func main() { // main function is the entry point of every go program
	fmt.Println("hello, world")
}

// Note: In go, there is build option which means creatung executable file from the source code.
// Syntax: go build <filename.go>	Example: go build main.go
// Advantage of build option is that it creates an executable file which can be run without the need of go compiler.
// We can simply share the binary file with the APIs to use it without the need of go compiler.
