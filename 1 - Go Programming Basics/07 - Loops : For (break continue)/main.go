package main

import "fmt"

func main() {

	// Simple iteration over range
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("Index: ", index, "Value: ", value)
		fmt.Printf("Index: %d, Value: %d\n", index, value) // %d is for integer
		fmt.Printf("Index: %v, Value: %v\n", index, value) // %v is for any value
	}

	// Break & Continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // continue the loop but skip the rest of lines/statements in this iteration
		}
		fmt.Println(i)
		if i == 5 {
			break // break out of the loop
		}
	}

}
