package main

import "fmt"

var firstName = "John" // Global variable

func main () {
	var age int = 10 // var <name> <type> = <value>. Type is optional if value is given. For global variables, var is mandatory.
	var age1 = 10 // var <name> = <value> (type is inferred). We can also use := for this.
	fmt.Println(age)
	fmt.Println(age1)

	age2 := 10	// := is shorthand for var <name> = <value> (type is inferred) and can only be used inside functions.
	// var age3 int := 10	// We cannot use var and int together with :=
	fmt.Println(age2)
	fmt.Println(firstName)

	// Default values
	// Integers: 0
	// Floats: 0.0
	// Strings: ""
	// Booleans: false
	// Pointers: nil
	// Arrays, slices, maps, Functions: nil
	// Structs: each field is set to its default value
}

