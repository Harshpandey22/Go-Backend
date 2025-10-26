package main

func main() {
	// PascalCase: Each word starts with a capital letter. Used for naming types, structs, interfaces, and exported functions/variables.
	// Pascal case is used for structs, enums, types, interfaces.
	// Example: CalculatedArea, UserProfile, NewHttpRequest

	// Lower -> snake_case: Words are separated by underscores. Not commonly used in Go, but sometimes seen in constants.
	// Snake case is used in constants.
	// Example: user_id, max_value, default_timeout, api_key, etc.

	// UPPER -> SNAKE_CASE: All letters are uppercase, and words are separated by underscores. Typically used for constants that are meant to be immutable.
	// Upper snake case is used in constants that are immutable.
	// Example: MAX_CONNECTIONS, DEFAULT_TIMEOUT, API_KEY

	// camelCase: The first word starts with a lowercase letter, and subsequent words start with a capital letter. Commonly used for naming variables and functions.
	// Camel case is used in variables and functions.
	// Example: calculateArea, userProfile, newHttpRequest (variables and functions)
}
