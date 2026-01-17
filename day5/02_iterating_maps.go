// Day 5, Exercise 2: Iterating Over Maps
//
// Key concepts:
// - Use range to iterate over maps
// - Map iteration order is NOT guaranteed (randomized)
// - Can iterate over keys only, or keys and values
// - For sorted iteration, extract keys to slice and sort

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Basic Iteration with Range ===")

	prices := map[string]float64{
		"apple":  1.20,
		"banana": 0.50,
		"orange": 0.80,
		"mango":  2.50,
	}

	// Iterate over key-value pairs
	fmt.Println("All prices:")
	for fruit, price := range prices {
		fmt.Printf("  %s: $%.2f\n", fruit, price)
	}

	fmt.Println("\n=== Iterate Over Keys Only ===")

	// Use _ to ignore value
	fmt.Print("Fruits available: ")
	for fruit := range prices {
		fmt.Print(fruit, " ")
	}
	fmt.Println()

	fmt.Println("\n=== Iteration Order Is Random ===")

	// Run this multiple times - order may change!
	numbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	fmt.Println("First iteration:")
	for k, v := range numbers {
		fmt.Printf("  %d: %s\n", k, v)
	}

	fmt.Println("Second iteration (may be different order):")
	for k, v := range numbers {
		fmt.Printf("  %d: %s\n", k, v)
	}

	fmt.Println("\n=== Sorted Iteration by Keys ===")

	grades := map[string]int{
		"Charlie": 85,
		"Alice":   92,
		"Bob":     78,
		"Diana":   95,
	}

	// Step 1: Extract keys to a slice
	names := make([]string, 0, len(grades))
	for name := range grades {
		names = append(names, name)
	}

	// Step 2: Sort the slice
	sort.Strings(names)

	// Step 3: Iterate using sorted keys
	fmt.Println("Grades (alphabetical order):")
	for _, name := range names {
		fmt.Printf("  %s: %d\n", name, grades[name])
	}

	fmt.Println("\n=== Sorted by Values ===")

	// To sort by values, we need a different approach
	type nameGrade struct {
		name  string
		grade int
	}

	var students []nameGrade
	for name, grade := range grades {
		students = append(students, nameGrade{name, grade})
	}

	// Sort by grade (descending)
	sort.Slice(students, func(i, j int) bool {
		return students[i].grade > students[j].grade
	})

	fmt.Println("Grades (highest first):")
	for _, s := range students {
		fmt.Printf("  %s: %d\n", s.name, s.grade)
	}

	fmt.Println("\n=== Practical Example: Grouping Data ===")

	// Group people by their first letter
	people := []string{"Alice", "Bob", "Anna", "Charlie", "Ben", "Carol"}
	byLetter := make(map[byte][]string)

	for _, name := range people {
		firstLetter := name[0]
		byLetter[firstLetter] = append(byLetter[firstLetter], name)
	}

	fmt.Println("People grouped by first letter:")
	// Get sorted letters
	var letters []byte
	for letter := range byLetter {
		letters = append(letters, letter)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	for _, letter := range letters {
		fmt.Printf("  %c: %v\n", letter, byLetter[letter])
	}

	fmt.Println("\n=== Filtering While Iterating ===")

	products := map[string]float64{
		"laptop":    999.99,
		"mouse":     29.99,
		"keyboard":  79.99,
		"monitor":   299.99,
		"headphone": 49.99,
	}

	// Find products under $100
	fmt.Println("Products under $100:")
	for product, price := range products {
		if price < 100 {
			fmt.Printf("  %s: $%.2f\n", product, price)
		}
	}

	// Calculate total value
	var total float64
	for _, price := range products {
		total += price
	}
	fmt.Printf("Total inventory value: $%.2f\n", total)
}

// TO RUN: go run day5/02_iterating_maps.go
//
// OUTPUT:
// === Basic Iteration with Range ===
// All prices:
//   apple: $1.20
//   banana: $0.50
//   ...
//
// EXERCISE:
// 1. Create a map of month numbers to month names (1 -> "January", etc.)
// 2. Print all months in order (1-12)
// 3. Find and print only months with more than 5 letters in their name
//
// KEY POINTS:
// - range gives key, value pairs
// - Iteration order is NOT deterministic
// - For sorted output, extract keys, sort, then iterate
// - Can sort by values using slice of structs
