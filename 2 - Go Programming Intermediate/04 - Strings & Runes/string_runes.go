package main

import (
	"fmt"
	"unicode/utf8"
)

// Strings in Go are sequences of bytes that represent text.They are immutable, meaning once created, their content cannot be changed.
// Strings can contain any byte values, including UTF-8 encoded characters.

// Runes are Go's representation of Unicode code points.
// A rune is an alias for int32 and can represent a single character.
// Runes are used to handle characters beyond the ASCII set, allowing for internationalization.

// Syntax:
// To declare a string: var s string = "Hello, World!"
// To declare a rune: var r rune = 'a' (single quotes for runes)
// To convert a string to a slice of runes: runes := []rune(s)
// var s string = "Hello, 世界"
// var r rune = '世'
// runes := []rune(s) // Converts string to slice of runes

func main() {

	message := "Hello, \nWorld"     // Normal string with escape sequence
	rawMessage := `Hello\nWordld`   // Raw string literal, ignores escape sequences
	rawMessage1 := "Hello, \tWorld" // Normal string with tab escape sequence
	rawMessage2 := "Hello, \rWorld" // Normal string with carriage return escape sequence

	fmt.Println(message)     // Output: Hello, /nWorld
	fmt.Println(rawMessage)  // Output: Hello\nWorld
	fmt.Println(rawMessage1) // Output: Hello, 	World
	fmt.Println(rawMessage2) // Output: World,

	
	fmt.Println("The first character in message var is:", message[0]) // Accessing individual byte (character) by index, ASCII


	str1 := "Apple"  // A has an ASCII value of 65
	str2 := "apple"  // a has an ASCII value of 97
	str3 := "Banana" // B has an ASCII value of 66
	str4 := "app"    // a has an ASCII value of 97

	fmt.Println(str1 < str2) // true, because 'A' (65) is less than 'a' (97)
	fmt.Println(str3 < str1) // false, because 'B' (66) is greater than 'A' (65)
	fmt.Println(str2 > str1) // true, because 'a' (97) is greater than 'A' (65)
	fmt.Println(str2 > str4) // true, because 'a' (97) is greater than 'a' (97) but 'p' (112) is greater than end of string


	// Loop over a string
	for i, char := range message {
		fmt.Printf("Character at index %d is %c \n", i, char) // Using range to iterate over string, char is of type rune
	}


	// using %x to print hexadecimal values of runes in the string
	for _, char := range message {
		fmt.Printf("%x\n", char) // Print the hexadecimal value of each rune in the string
	}


	// using %v to print the integer values of runes in the string
	for _, char := range message {
		fmt.Printf("%v\n", char) // Print the default format (integer value of rune)
	}


	fmt.Println("Rune count: ", utf8.RuneCountInString(message)) // Counting runes in the string

	
	// Declaring runes
	var ch rune = 'a' // Declaring a rune
	jcg := '世'        // Declaring a rune with a Unicode character

	fmt.Println("Rune ch:", ch)              // Output: 97
	fmt.Println("Rune jcg:", jcg)            // Output: 19990
	fmt.Printf("Rune ch in hex: %x\n", ch)   // Output: 61
	fmt.Printf("Rune jcg in hex: %x\n", jcg) // Output: 4e16


	// Converting runes to strings
	cstr := string(ch)
	fmt.Println("String from rune ch:", cstr) // Output: a
	fmt.Printf("Type of cstr: %T\n", cstr)    // Output: string

}
