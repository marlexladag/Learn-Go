// Day 3, Exercise 4: Multiple Return Values
//
// Key concepts:
// - Go functions can return multiple values
// - Common pattern: return value AND error
// - Use parentheses for multiple return types
// - Use _ (blank identifier) to ignore unwanted values

package main

import (
	"fmt"
	"math"
)

// Function returning two values
func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false // Return zero and "not ok"
	}
	return a / b, true // Return result and "ok"
}

// Return multiple related values
func minMax(numbers ...int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// Return calculation results
func calculate(a, b int) (int, int, int, int) {
	sum := a + b
	diff := a - b
	prod := a * b
	quot := 0
	if b != 0 {
		quot = a / b
	}
	return sum, diff, prod, quot
}

// Common pattern: return value and error status
func squareRoot(n float64) (float64, string) {
	if n < 0 {
		return 0, "cannot compute square root of negative number"
	}
	return math.Sqrt(n), ""
}

// Return multiple strings
func splitName(fullName string) (string, string) {
	// Simple split - assumes "FirstName LastName" format
	firstName := ""
	lastName := ""
	spaceFound := false

	for _, char := range fullName {
		if char == ' ' {
			spaceFound = true
			continue
		}
		if spaceFound {
			lastName += string(char)
		} else {
			firstName += string(char)
		}
	}
	return firstName, lastName
}

// Swap two values
func swap(a, b int) (int, int) {
	return b, a
}

// Return coordinates
func getPosition() (x, y, z int) {
	return 10, 20, 30
}

func main() {
	fmt.Println("=== Basic Multiple Returns ===")
	result, ok := divide(10, 3)
	if ok {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	result, ok = divide(10, 0)
	if !ok {
		fmt.Println("Division by zero attempted!")
	}

	fmt.Println("\n=== Ignoring Values with _ ===")
	// Sometimes we only want one of the returned values
	quotient, _ := divide(20, 4) // Ignore the 'ok' value
	fmt.Printf("20 / 4 = %.2f (ignored ok status)\n", quotient)

	fmt.Println("\n=== Multiple Calculations ===")
	sum, diff, prod, quot := calculate(20, 5)
	fmt.Printf("For 20 and 5:\n")
	fmt.Printf("  Sum: %d\n", sum)
	fmt.Printf("  Difference: %d\n", diff)
	fmt.Printf("  Product: %d\n", prod)
	fmt.Printf("  Quotient: %d\n", quot)

	fmt.Println("\n=== Min and Max ===")
	min, max := minMax(5, 2, 9, 1, 7, 3)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	fmt.Println("\n=== Error Handling Pattern ===")
	val, err := squareRoot(16)
	if err == "" {
		fmt.Printf("sqrt(16) = %.2f\n", val)
	}

	val, err = squareRoot(-4)
	if err != "" {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("\n=== String Splitting ===")
	first, last := splitName("John Doe")
	fmt.Printf("First: %s, Last: %s\n", first, last)

	first, last = splitName("Jane Smith")
	fmt.Printf("First: %s, Last: %s\n", first, last)

	fmt.Println("\n=== Swap Values ===")
	a, b := 100, 200
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	fmt.Println("\n=== Multiple Return Types ===")
	x, y, z := getPosition()
	fmt.Printf("Position: x=%d, y=%d, z=%d\n", x, y, z)

	// Ignore some return values
	x, _, z = getPosition() // Only get x and z
	fmt.Printf("Got x=%d and z=%d (ignored y)\n", x, z)
}

// TO RUN: go run 04_multiple_returns.go
//
// OUTPUT:
// === Basic Multiple Returns ===
// 10 / 3 = 3.33
// Division by zero attempted!
//
// === Ignoring Values with _ ===
// 20 / 4 = 5.00 (ignored ok status)
//
// ... (continues)
//
// KEY POINTS:
// - Multiple return types go in parentheses: func f() (int, string)
// - Return multiple values: return val1, val2
// - Receive multiple values: a, b := function()
// - Use _ (blank identifier) to ignore unwanted values
// - Common Go pattern: return (result, error) or (result, ok)
// - This pattern is used extensively in Go's standard library
//
// EXERCISE: Write a function that takes a slice of numbers and
// returns the sum, average, min, and max all at once
