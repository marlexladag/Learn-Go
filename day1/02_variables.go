// Day 1, Exercise 2: Variables and Types
//
// Key concepts:
// - var keyword declares variables
// - := is shorthand declaration (type inferred)
// - Go is statically typed
// - Basic types: int, float64, string, bool

package main

import "fmt"

func main() {
	// Method 1: var with explicit type
	var name string = "Gopher"
	var age int = 25

	// Method 2: var with type inference
	var city = "San Francisco"

	// Method 3: Short declaration (most common)
	country := "USA"
	temperature := 72.5 // float64
	isLearning := true  // bool

	// Printing variables
	fmt.Println("=== Variable Examples ===")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Is Learning:", isLearning)

	// Multiple variable declaration
	var x, y, z int = 1, 2, 3
	fmt.Println("\nMultiple vars:", x, y, z)

	// Zero values (default values)
	var defaultInt int
	var defaultString string
	var defaultBool bool
	var defaultFloat float64

	fmt.Println("\n=== Zero Values ===")
	fmt.Println("int:", defaultInt)       // 0
	fmt.Println("string:", defaultString) // "" (empty)
	fmt.Println("bool:", defaultBool)     // false
	fmt.Println("float64:", defaultFloat) // 0
}

// TO RUN: go run 02_variables.go
//
// EXERCISE: Add your own variables for:
// 1. Your favorite number
// 2. Your hobby (string)
// 3. Whether you like coffee (bool)
// Then print them!
