package main

// If-Else Condition is used to execute code based on certain conditions
// Use: When you need to make decisions in your code based on specific conditions, allowing for different execution paths.

// Syntax:
// if condition {
//     // code to be executed if condition is true
// } else if anotherCondition {
//     // code to be executed if anotherCondition is true
// } else {
//     // code to be executed if both conditions are false
// }

func main() {

	// Example 1: Simple If-Else
	age := 20
	if age < 18 {
		println("Minor")

	} else {
		println("Senior Citizen")
	}

	// Example 2: If-Else If-Else
	marks := 85
	if marks >= 90 {
		println("Grade A")
	} else if marks >= 75 {
		println("Grade B")
	} else if marks >= 60 {
		println("Grade C")
	} else {
		println("Grade F")
	}

	// Example 3: Nested If
	num := 10
	if num >= 0 {
		if num == 0 {
			println("Zero")
		} else {
			println("Positive Number")
		}
	} else {
		println("Negative Number")
	}
}
