// Day 3, Exercise 2: Function Parameters
//
// Key concepts:
// - Passing data to functions via parameters
// - Parameter types must be specified
// - Multiple parameters with same type shorthand
// - Parameters are copies (pass by value)

package main

import "fmt"

// Function with one parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with two parameters of different types
func printInfo(name string, age int) {
	fmt.Printf("%s is %d years old\n", name, age)
}

// Multiple parameters of same type - shorthand syntax
// Instead of: func add(a int, b int)
func add(a, b int) {
	result := a + b
	fmt.Printf("%d + %d = %d\n", a, b, result)
}

// More parameters with mixed shorthand
func describePerson(firstName, lastName string, age int, height float64) {
	fmt.Printf("Name: %s %s\n", firstName, lastName)
	fmt.Printf("Age: %d, Height: %.1f cm\n", age, height)
}

// Demonstrating pass-by-value
func tryToModify(x int) {
	fmt.Printf("  Inside function, x = %d\n", x)
	x = 999 // This only modifies the local copy!
	fmt.Printf("  After modification, x = %d\n", x)
}

// Function with a string parameter
func repeat(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(text)
	}
}

func main() {
	fmt.Println("=== Single Parameter ===")
	greet("Alice")
	greet("Bob")
	greet("Go Developer")

	fmt.Println("\n=== Two Parameters ===")
	printInfo("Alice", 25)
	printInfo("Bob", 30)

	fmt.Println("\n=== Same Type Shorthand ===")
	add(5, 3)
	add(100, 200)
	add(-10, 10)

	fmt.Println("\n=== Multiple Parameters ===")
	describePerson("John", "Doe", 28, 175.5)

	fmt.Println("\n=== Pass By Value Demo ===")
	number := 42
	fmt.Printf("Before function call, number = %d\n", number)
	tryToModify(number)
	fmt.Printf("After function call, number = %d\n", number)
	fmt.Println("(Notice: original value unchanged!)")

	fmt.Println("\n=== Practical Example ===")
	repeat("Go is awesome!", 3)
}

// TO RUN: go run 02_parameters.go
//
// OUTPUT:
// === Single Parameter ===
// Hello, Alice!
// Hello, Bob!
// Hello, Go Developer!
//
// === Two Parameters ===
// Alice is 25 years old
// Bob is 30 years old
//
// === Same Type Shorthand ===
// 5 + 3 = 8
// 100 + 200 = 300
// -10 + 10 = 0
//
// === Multiple Parameters ===
// Name: John Doe
// Age: 28, Height: 175.5 cm
//
// === Pass By Value Demo ===
// Before function call, number = 42
//   Inside function, x = 42
//   After modification, x = 999
// After function call, number = 42
// (Notice: original value unchanged!)
//
// === Practical Example ===
// Go is awesome!
// Go is awesome!
// Go is awesome!
//
// KEY POINTS:
// - Parameters are declared as: name type
// - Multiple params of same type: a, b int (shorthand for a int, b int)
// - Go is pass-by-value: function gets a copy of the argument
// - Changes to parameters inside function don't affect original
// - We'll learn about pointers later to modify original values
//
// EXERCISE: Write a function calculateArea(length, width float64)
// that prints the area of a rectangle
