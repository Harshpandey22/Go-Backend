package main

// Range in Go is used to iterate over elements in various data structures like arrays, slices, maps, strings, and channels.
// Use: When you need to loop through elements in a collection without manually managing the index or iterator.

// Syntax:
// for index, value := range collection {
//     // use index and value
// }

func main() {

	message := "Hello World"
	for i, v := range message {
		println(i, string(v))
	}
}
