// Day 4, Exercise 1: Arrays
//
// Key concepts:
// - Arrays are fixed-size sequences of elements of the same type
// - Array size is part of the type ([3]int is different from [4]int)
// - Arrays are value types (copied when assigned or passed)
// - Zero values: arrays are initialized with zero values of their element type

package main

import "fmt"

func main() {
	fmt.Println("=== Array Declaration ===")

	// Method 1: Declare with var (initialized to zero values)
	var numbers [5]int
	fmt.Println("Zero-value array:", numbers) // [0 0 0 0 0]

	// Method 2: Declare and initialize with values
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("Initialized array:", primes)

	// Method 3: Let compiler count the elements with [...]
	vowels := [...]string{"a", "e", "i", "o", "u"}
	fmt.Println("Auto-sized array:", vowels)
	fmt.Println("Length:", len(vowels))

	// Method 4: Initialize specific indices
	sparse := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Println("Sparse array:", sparse) // [10 0 30 0 50]

	fmt.Println("\n=== Accessing Elements ===")

	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	// Access by index (0-based)
	fmt.Println("First day:", days[0])
	fmt.Println("Last day:", days[6])
	fmt.Println("Last day (using len):", days[len(days)-1])

	// Modify elements
	days[0] = "Monday"
	fmt.Println("Modified first day:", days[0])

	fmt.Println("\n=== Iterating Over Arrays ===")

	scores := [4]int{95, 87, 92, 88}

	// Method 1: Traditional for loop
	fmt.Print("Traditional loop: ")
	for i := 0; i < len(scores); i++ {
		fmt.Print(scores[i], " ")
	}
	fmt.Println()

	// Method 2: Range (preferred) - gives index and value
	fmt.Print("Range loop: ")
	for index, value := range scores {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println()

	// Method 3: Range with only index
	fmt.Print("Index only: ")
	for i := range scores {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Method 4: Range with only value (use _ for index)
	fmt.Print("Value only: ")
	for _, score := range scores {
		fmt.Print(score, " ")
	}
	fmt.Println()

	fmt.Println("\n=== Arrays Are Value Types ===")

	original := [3]int{1, 2, 3}
	copied := original // Creates a COPY, not a reference

	copied[0] = 100

	fmt.Println("Original:", original) // [1 2 3] - unchanged!
	fmt.Println("Copied:", copied)     // [100 2 3]

	fmt.Println("\n=== Array Comparison ===")

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	a3 := [3]int{1, 2, 4}

	fmt.Println("a1 == a2:", a1 == a2) // true
	fmt.Println("a1 == a3:", a1 == a3) // false
	// Note: Can only compare arrays of the same type and size

	fmt.Println("\n=== Array Length ===")

	arr := [10]int{1, 2, 3} // Rest will be 0
	fmt.Println("Array:", arr)
	fmt.Println("Length:", len(arr)) // Always 10, not 3!

	fmt.Println("\n=== Practical Example: Calculate Average ===")

	grades := [5]float64{85.5, 92.0, 78.5, 90.0, 88.5}

	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	average := sum / float64(len(grades))

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Average: %.2f\n", average)

	fmt.Println("\n=== Finding Min and Max ===")

	temperatures := [7]int{72, 68, 75, 80, 77, 65, 71}

	min, max := temperatures[0], temperatures[0]
	for _, temp := range temperatures {
		if temp < min {
			min = temp
		}
		if temp > max {
			max = temp
		}
	}

	fmt.Println("Temperatures:", temperatures)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}

// TO RUN: go run day4/01_arrays.go
//
// OUTPUT:
// === Array Declaration ===
// Zero-value array: [0 0 0 0 0]
// Initialized array: [2 3 5 7 11]
// Auto-sized array: [a e i o u]
// Length: 5
// Sparse array: [10 0 30 0 50]
//
// === Accessing Elements ===
// First day: Mon
// Last day: Sun
// ...
//
// EXERCISE:
// 1. Create an array of your 5 favorite foods
// 2. Print them in reverse order
// 3. Count how many start with a vowel
//
// KEY POINTS:
// - Arrays have FIXED size (part of the type)
// - Arrays are VALUE types (copying creates independent copy)
// - Use [...] to let compiler count elements
// - Use range for idiomatic iteration
// - len() returns the declared size, not "used" elements
