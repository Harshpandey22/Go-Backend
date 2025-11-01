package main

import "fmt"

// The fmt package in Go provides I/O formatting functions similar to C's printf and scanf.
// Use: The fmt package is used for formatted I/O operations, such as printing to the console or reading input.

// Syntax:
//  - Import the package: import "fmt"
//  - Print formatted output: fmt.Printf(format string, values...)

func main() {

	// Print function prints to standard output without a newline
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 465)

	// Println function adds a newline after the output
	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 465)

	// Printf function allows formatted output using verbs
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %x\n", age, age)

	// Formatting Functions

	// Sprint returns the formatted string without printing
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Println(s) // Outputs: Hello World!123456

	// Sprintf returns the formatted string according to a format specifier
	s2 := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(s2) // Outputs: Name: Alice, Age: 30

	s3 := fmt.Sprintln("Hello", "World", 123, 456)
	fmt.Println(s3) // Outputs: Hello World 123 456 (with newline)

	// Scanning Functions

	// Scan reads from standard input
	var inputName string
	var inputAge int
	fmt.Print("Enter your name and age: ")
	fmt.Scan(&inputName, &inputAge)
	fmt.Printf("You entered - Name: %s, Age: %d\n", inputName, inputAge)

	// Scanln reads a line from standard input
	var lineName string
	var lineAge int
	fmt.Print("Enter your name and age (line): ")
	fmt.Scanln(&lineName, &lineAge)
	fmt.Printf("You entered - Name: %s, Age: %d\n", lineName, lineAge)

	// Error formatting functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age <18 {
		return fmt.Errorf("age %d is below the legal limit", age)
	}
	return nil
}

