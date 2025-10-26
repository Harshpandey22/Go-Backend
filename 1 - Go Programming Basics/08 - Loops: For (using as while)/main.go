package main

import "fmt"

// Unlike other programming languages, Go doesn't have a dedicated keyword for a while loop.
// However, we can use the for loop to perform the functionality of a while loop.

func main() {

	// Here is an example of using a for loop as a while loop:
	i := 1
	for i <= 10 {
		fmt.Println("Iteration: ", i)
		i++
	}

	// Infinite Loop
	// for{
	// 	fmt.Println("Infinite Loop")
	// }

	// for as while with break
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum: ", sum)
		if sum >= 50 {
			break
		}
	}

	// for as while with continue
	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number: ", num)
		num++
	}
}
