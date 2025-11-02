package main

import (
	"errors"
	"fmt"
)

// Errors in Go are straightforward: they are values that represent an abnormal condition and are primarily handled by checking the last return value of a function.
// Use: The core uses of errors in Go are to signal failure and to provide contextual information about what went wrong, allowing the calling code to handle the specific failure condition.


func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative number") // Creating custom error using errors.New
	}
	// compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}
	// Process data
	return nil
}
func main() {

	result, err1 := sqrt(-1)
	if err1 != nil {
		fmt.Println(err1) // Output: Math Error: square root of negative number
		return
	}
	fmt.Println(result)


	data := []byte{}
	if err2 := process(data); err2 != nil {
		fmt.Println(err2) // Output: Error: Empty data
		return
	}
	fmt.Println("Data Processed Successfully")

}
