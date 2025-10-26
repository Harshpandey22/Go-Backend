package main

import "fmt"

// Switch Condition is used to execute one block of code among many based on the value of a variable or expression
// Use: When you have multiple possible values for a variable and want to execute different code for each value, making the code cleaner and more readable than multiple if-else statements.

// Switch statement in go is (switch case default)
// Syntax:
// switch expression {
// case value1:
//     // code to be executed if expression == value1
// case value2:
//     // code to be executed if expression == value2
// ...
// default:
//     // code to be executed if expression doesn't match any case
// }

// Switch statement in other languages is (switch case break default)
// Syntax:
// switch (expression) {
// case value1:
//     // code to be executed if expression == value1
//     break;
// case value2:
//     // code to be executed if expression == value2
//     break;
// ...
// default:
//     // code to be executed if expression doesn't match any case
// }

func main() {

	// Simple switch case
	fruit := "apple"
	switch fruit {
	case "apple":
		println("This is an apple")
	case "banana":
		println("This is a banana")
	case "orange":
		println("This is an orange")
	default:
		println("Unknown fruit")
	}

	// Multiple cases
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		println("It's a weekday")
	case "Saturday", "Sunday":
		println("It's the weekend")
	default:
		println("Unknown day")
	}

	// Switch with expression
	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number <= 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is greater than 20")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2")
	default:
		fmt.Println("Not two")
	}
}
