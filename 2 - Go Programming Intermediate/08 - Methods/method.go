package main

import "fmt"

// Go methods are like functions but with a key difference: they have a receiver argument, which allows access to the receiver's properties.
// The receiver can be a struct or non-struct type, but both must be in the same package. Methods cannot be created for types defined in other packages, including built-in types like int or string; otherwise, the compiler will raise an error.
// Use: Primarily for encapsulation and behavioral extension, where a type (like a struct) is given specific actions it can perform, most visibly in method chaining for fluent APIs (e.g., db.Where(...).Order(...).Find(...)).

// Syntax:
// func (receiverName ReceiverType) MethodName(parameters) (returnTypes) {
//     // method body
// }

type Rectangle struct {
	length float64
	width  float64
}

// Method with Value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width

}

// Method with Pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // r.length = r.length*factor
	r.width *= factor
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle is ", area) // Output: Area of rectangle is  90

	rect.Scale(2)
	fmt.Println("Area of rectangle with a factor of 2 is", area) // Output: Area of rectangle with a factor of 2 is 90
}
