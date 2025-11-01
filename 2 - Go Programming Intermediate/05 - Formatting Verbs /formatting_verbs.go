package main

import "fmt"

// Formatting verbs in Go are used with the fmt package to format strings, numbers, and other data types.

func main() {

	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v 	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints a literal percent sign

	i := 15.5
	str := "Hello World"
	fmt.Printf("%v\n", i)    // Default format, outputs: 15.5
	fmt.Printf("%#v\n", i)   // Go-syntax format, outputs: 15.5
	fmt.Printf("%T\n", i)    // Type of the value, outputs: float64
	fmt.Printf("%%\n")       // Literal percent sign, outputs: %
	fmt.Printf("%v\n", str)  // Default format, outputs: Hello World
	fmt.Printf("%#v\n", str) // Go-syntax format, outputs: "Hello World"
	fmt.Printf("%T\n", str)  // Type of the value, outputs: string

	

	// --- Integer Formatting Verbs
	// %b	Base 2
	// %d	Base 10
	// %+d 	Base 10 and always show sign
	// %o 	Base 8
	// %O 	Base 8 with leading 0o
	// %x	Base 16, with lower-case letters for a-f
	// %X	Base 16, with upper-case letters for A-F
	// %#x	Base 16, with leading 0x
	// %4d 	Pad with spaces (width 4, right-aligned)
	// %-4d Pad with spaces (width 4, left-aligned)
	// %04d Pad with leading zeros (width 4)

	a := 18
	fmt.Printf("%b\n", a)   // Base 2, outputs: 10010
	fmt.Printf("%d\n", a)   // Base 10, outputs: 18
	fmt.Printf("%+d\n", a)  // Base 10 with sign, outputs: +18
	fmt.Printf("%o\n", a)   // Base 8, outputs: 22
	fmt.Printf("%O\n", a)   // Base 8 with leading 0o, outputs: 0o22
	fmt.Printf("%x\n", a)   // Base 16 lower-case, outputs: 12
	fmt.Printf("%X\n", a)   // Base 16 upper-case, outputs: 12
	fmt.Printf("%#x\n", a)  // Base 16 with leading 0x, outputs: 0x12
	fmt.Printf("%4d\n", a)  // Width 4 right-aligned, outputs: "  18"
	fmt.Printf("%-4d\n", a) // Width 4 left-aligned, outputs: "18  "
	fmt.Printf("%04d\n", a) // Width 4 with leading zeros, outputs: 0018



	// --- String Formatting Verbs
	// %s 	Prints the value as plain string
	// %q 	Prints the value as double-quoted string
	// %8s 	Prints the value as plain string (width 8, right-aligned)
	// %-8s Prints the value as plain string (width 8, left-aligned)
	// %x 	Prints the value as hex dump of byte values
	// % x 	Prints the value as hex dump with spaces between bytes
	name := "Gopher"
	fmt.Printf("%s\n", name)   // Plain string, outputs: Gopher
	fmt.Printf("%q\n", name)   // Double-quoted string, outputs: "Gopher"
	fmt.Printf("%8s\n", name)  // Width 8 right-aligned, outputs: "  Gopher"
	fmt.Printf("%-8s\n", name) // Width 8 left-aligned, outputs: "Gopher  "
	fmt.Printf("%x\n", name)   // Hex dump, outputs: 476f70686572
	fmt.Printf("% x\n", name)  // Hex dump with spaces, outputs: 47 6f 70 68 65 72



	// --- Boolean Formatting Verbs
	// %t	Prints the value as true or false
	flag := true
	fmt.Printf("%t\n", flag) // Outputs: true



	// --- Float Formatting Verbs
	// %e 		Scientific notation with 'e' as exponent (e.g., 1.23e+04)
	// %f 		Decimal point but no exponent (e.g., 1234.5678)
	// %.2f 	Decimal point with 2 digits after the point (e.g., 1234.57)
	// %6.2f 	Width 6, with 2 digits after the point (e.g., "1234.57" or " 123.45")
	// %-6.2f 	Width 6, left-aligned, with 2 digits after the point
	// %g 	  	Exponent as needed, but shorter representation of %e or %f
	num := 1234.56789
	fmt.Printf("%e\n", num)     // Scientific notation, outputs: 1.234568e+03
	fmt.Printf("%f\n", num)     // Decimal point, outputs: 1234.567890
	fmt.Printf("%.2f\n", num)   // 2 digits after point, outputs: 1234.57
	fmt.Printf("%6.2f\n", num)  // Width 6 with 2 digits, outputs: 1234.57
	fmt.Printf("%-6.2f\n", num) // Width 6 left-aligned with 2 digits, outputs: 1234.57
	fmt.Printf("%g\n", num)     // Shorter representation, outputs: 1234.57
}
