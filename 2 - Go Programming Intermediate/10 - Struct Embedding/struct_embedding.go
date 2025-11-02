package main

import "fmt"

// Struct embedding in Go is the mechanism used to promote fields and methods of one struct (the embedded type) into another struct (the outer type),
// providing a way to achieve composition and code reuse similar to inheritance in other languages, but with a simpler and more explicit design.
// Use: The main use of struct embedding in Go is to achieve composition and code reuse without the rigidity of traditional class inheritance, particularly by promoting methods for interface satisfaction and simplifying wrappers.

// Syntax:
// // 1. Define the embedded (inner) type
// type CarEngine struct {
// 	Horsepower int
// }

// // 2. Define the embedding (outer) type
// type SportsCar struct {
// 	// Embedding the CarEngine struct.
// 	// Notice there is NO field name here—just the type name.
// 	CarEngine

// 	// Regular field
// 	Color string
// }

type person struct {
	name string
	age  int
}

type Employee struct {
	person // Embedded Struct
	empId  string
	salary float64
}

// Method Inheritance (Embedding struct in Methods)
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.", p.name, p.age)
}

func main() {

	emp := Employee{
		person: person{name: "John", age: 30},
		empId:  "E001",
		salary: 50000,
	}

	fmt.Println("Name:", emp.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", emp.age)   // Same as above
	fmt.Println("Emp Id", emp.empId)
	fmt.Println("Salary", emp.salary)

	emp.introduce() // Output: Hi, I'm John and I'm 30 years old.%
}
