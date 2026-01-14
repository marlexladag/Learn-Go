// Day 3, Exercise 6: Variadic Functions
//
// Key concepts:
// - Variadic functions accept variable number of arguments
// - Use ... (ellipsis) before type: func f(nums ...int)
// - Inside function, variadic parameter is a slice
// - Can pass a slice using ... spread operator

package main

import "fmt"

// Variadic function - accepts any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Variadic with regular parameter (variadic must be last!)
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// Finding maximum of variable arguments
func max(first int, rest ...int) int {
	maximum := first
	for _, n := range rest {
		if n > maximum {
			maximum = n
		}
	}
	return maximum
}

// Finding minimum
func min(first int, rest ...int) int {
	minimum := first
	for _, n := range rest {
		if n < minimum {
			minimum = n
		}
	}
	return minimum
}

// Concatenate strings with separator
func joinStrings(separator string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += separator + parts[i]
	}
	return result
}

// Calculate average
func average(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := 0.0
	for _, n := range numbers {
		total += n
	}
	return total / float64(len(numbers))
}

// Printf is a famous variadic function!
// func Printf(format string, a ...interface{}) (n int, err error)

// Counting how many match a condition
func countPositive(numbers ...int) int {
	count := 0
	for _, n := range numbers {
		if n > 0 {
			count++
		}
	}
	return count
}

// Check if all values are true
func allTrue(values ...bool) bool {
	for _, v := range values {
		if !v {
			return false
		}
	}
	return true
}

// Check if any value is true
func anyTrue(values ...bool) bool {
	for _, v := range values {
		if v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("=== Basic Variadic ===")
	fmt.Printf("sum() = %d\n", sum())
	fmt.Printf("sum(5) = %d\n", sum(5))
	fmt.Printf("sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	fmt.Println("\n=== Passing Slice to Variadic ===")
	numbers := []int{10, 20, 30, 40, 50}
	// Use ... to spread a slice into variadic arguments
	fmt.Printf("sum(%v) = %d\n", numbers, sum(numbers...))

	fmt.Println("\n=== Mixed Parameters ===")
	greetAll("Hello", "Alice", "Bob", "Charlie")
	fmt.Println()
	greetAll("Welcome", "Go", "Developer")

	fmt.Println("\n=== Max and Min ===")
	fmt.Printf("max(3, 7, 2, 9, 5) = %d\n", max(3, 7, 2, 9, 5))
	fmt.Printf("min(3, 7, 2, 9, 5) = %d\n", min(3, 7, 2, 9, 5))
	fmt.Printf("max(42) = %d\n", max(42))

	fmt.Println("\n=== String Joining ===")
	fmt.Println(joinStrings(", ", "apple", "banana", "cherry"))
	fmt.Println(joinStrings(" - ", "Go", "is", "awesome"))
	fmt.Println(joinStrings("/", "home", "user", "documents"))

	fmt.Println("\n=== Average ===")
	fmt.Printf("average(1, 2, 3, 4, 5) = %.2f\n", average(1, 2, 3, 4, 5))
	fmt.Printf("average(10.5, 20.5, 30.0) = %.2f\n", average(10.5, 20.5, 30.0))

	fmt.Println("\n=== Count Positive ===")
	fmt.Printf("countPositive(1, -2, 3, -4, 5) = %d\n", countPositive(1, -2, 3, -4, 5))
	fmt.Printf("countPositive(-1, -2, -3) = %d\n", countPositive(-1, -2, -3))

	fmt.Println("\n=== Boolean Variadic ===")
	fmt.Printf("allTrue(true, true, true) = %t\n", allTrue(true, true, true))
	fmt.Printf("allTrue(true, false, true) = %t\n", allTrue(true, false, true))
	fmt.Printf("anyTrue(false, false, true) = %t\n", anyTrue(false, false, true))
	fmt.Printf("anyTrue(false, false, false) = %t\n", anyTrue(false, false, false))

	fmt.Println("\n=== Spreading Slices ===")
	scores := []float64{85.5, 90.0, 78.5, 92.0, 88.5}
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Average score: %.2f\n", average(scores...))

	words := []string{"one", "two", "three", "four"}
	fmt.Printf("Words: %v\n", words)
	fmt.Printf("Joined: %s\n", joinStrings("-", words...))
}

// TO RUN: go run 06_variadic.go
//
// OUTPUT:
// === Basic Variadic ===
// sum() = 0
// sum(5) = 5
// sum(1, 2, 3) = 6
// sum(1, 2, 3, 4, 5) = 15
//
// === Passing Slice to Variadic ===
// sum([10 20 30 40 50]) = 150
//
// ... (continues)
//
// KEY POINTS:
// - Variadic parameter: ...type (e.g., ...int, ...string)
// - Variadic parameter must be the last parameter
// - Inside function, variadic param is a slice
// - Call with any number of args: f(1, 2, 3)
// - Spread a slice with ...: f(slice...)
// - Empty variadic gives empty slice, not nil
//
// STANDARD LIBRARY EXAMPLES:
// - fmt.Printf(format string, a ...interface{})
// - fmt.Println(a ...interface{})
// - append(slice []T, elems ...T) []T
//
// EXERCISE: Write a function multiply(...int) that returns
// the product of all numbers (return 1 for empty input)
