package main

import "fmt"

// Structs in Go are used to group related data together. They are similar to classes in other programming languages but do not support inheritance.
// Use: Structs are used to create complex data types that group multiple fields.

// Syntax:
// type StructName struct {
//     Field1 Type1
//     Field2 Type2
//     ...
// }

// Defining a struct named Person. Structs are always defined outside of functions.
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// Creating instances of the Person struct
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	p1 := Person{
		firstName: "Jane",
		age:       30,
	}

	fmt.Println(p.firstName)  // Output: John
	fmt.Println(p1.firstName) // Output: Jane

	

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "user123@gmail.com",
	}

	fmt.Println(user.username, user.email) // Output: user123 user123@gmail.com
}
