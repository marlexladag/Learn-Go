// Day 3, Exercise 1: Function Basics
//
// Key concepts:
// - Declaring functions with func keyword
// - Functions help organize and reuse code
// - Functions can be called multiple times
// - Code inside a function runs when called, not when defined

package main

import "fmt"

// A simple function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// Another simple function
func printSeparator() {
	fmt.Println("================================")
}

// Function that does a specific task
func greetUser() {
	fmt.Println("Welcome to Day 3!")
	fmt.Println("Today we learn about functions.")
}

// Functions can call other functions
func showHeader() {
	printSeparator()
	fmt.Println("    Functions in Go")
	printSeparator()
}

func main() {
	// Calling our functions
	showHeader()

	// Call sayHello multiple times
	fmt.Println("\n=== Calling sayHello() ===")
	sayHello()
	sayHello()
	sayHello()

	// Call greetUser
	fmt.Println("\n=== Calling greetUser() ===")
	greetUser()

	// Functions make code readable
	fmt.Println("\n=== Functions Organize Code ===")
	printSeparator()
	fmt.Println("Code is more organized with functions!")
	printSeparator()

	// main() is also a function - the entry point
	fmt.Println("\n=== About main() ===")
	fmt.Println("main() is a special function")
	fmt.Println("Go programs start execution in main()")
}

// TO RUN: go run 01_functions_basics.go
//
// OUTPUT:
// ================================
//     Functions in Go
// ================================
//
// === Calling sayHello() ===
// Hello, World!
// Hello, World!
// Hello, World!
//
// === Calling greetUser() ===
// Welcome to Day 3!
// Today we learn about functions.
//
// === Functions Organize Code ===
// ================================
// Code is more organized with functions!
// ================================
//
// === About main() ===
// main() is a special function
// Go programs start execution in main()
//
// KEY POINTS:
// - Use 'func' keyword to declare a function
// - Function names follow same rules as variables (camelCase)
// - Parentheses () are required even with no parameters
// - Curly braces {} contain the function body
// - Functions must be called to execute their code
