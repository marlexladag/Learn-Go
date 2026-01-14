// Day 1, Exercise 4: Formatted Printing
//
// Key concepts:
// - fmt.Printf for formatted output
// - Format verbs: %s, %d, %f, %v, %T, etc.
// - Escape sequences: \n, \t

package main

import "fmt"

func main() {
	name := "Alice"
	age := 30
	height := 5.75
	isStudent := false

	fmt.Println("=== Printf Format Verbs ===")

	// %s - string
	fmt.Printf("Name: %s\n", name)

	// %d - integer (decimal)
	fmt.Printf("Age: %d years old\n", age)

	// %f - floating point
	fmt.Printf("Height: %f feet\n", height)

	// %.2f - floating point with 2 decimal places
	fmt.Printf("Height (2 decimals): %.2f feet\n", height)

	// %t - boolean
	fmt.Printf("Is Student: %t\n", isStudent)

	// %v - default format (works for any type)
	fmt.Printf("Name: %v, Age: %v, Student: %v\n", name, age, isStudent)

	// %T - type of the variable
	fmt.Printf("\n=== Types ===\n")
	fmt.Printf("name is type: %T\n", name)
	fmt.Printf("age is type: %T\n", age)
	fmt.Printf("height is type: %T\n", height)
	fmt.Printf("isStudent is type: %T\n", isStudent)

	// %+v - prints struct field names (useful later)
	// %#v - Go syntax representation

	// Width and padding
	fmt.Println("\n=== Width & Padding ===")
	fmt.Printf("|%10s|\n", "hello")  // right-aligned, width 10
	fmt.Printf("|%-10s|\n", "hello") // left-aligned, width 10
	fmt.Printf("|%010d|\n", 42)      // zero-padded, width 10

	// Sprintf - returns string instead of printing
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println("\nSprintf result:", message)
}

// TO RUN: go run 04_printf.go
//
// COMMON FORMAT VERBS:
// %v  - default format
// %T  - type
// %s  - string
// %d  - integer
// %f  - float
// %t  - boolean
// %p  - pointer
// %x  - hexadecimal
// %%  - literal percent sign
