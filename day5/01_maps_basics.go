// Day 5, Exercise 1: Maps Basics
//
// Key concepts:
// - Maps are key-value data structures (like dictionaries/hash tables)
// - Keys must be comparable types (no slices, maps, or functions as keys)
// - Maps are reference types (changes affect original)
// - Zero value of a map is nil (must initialize before use)

package main

import "fmt"

func main() {
	fmt.Println("=== Map Declaration ===")

	// Method 1: Using make() - most common
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	fmt.Println("Ages map:", ages)

	// Method 2: Map literal - declare and initialize
	capitals := map[string]string{
		"France":  "Paris",
		"Japan":   "Tokyo",
		"Germany": "Berlin",
		"Italy":   "Rome", // Note: trailing comma required on last item
	}
	fmt.Println("Capitals:", capitals)

	// Method 3: Empty map literal
	scores := map[string]int{}
	scores["Math"] = 95
	fmt.Println("Scores:", scores)

	// Note: This creates a NIL map (can read but NOT write!)
	var nilMap map[string]int
	fmt.Println("Nil map:", nilMap)
	fmt.Println("Nil map is nil:", nilMap == nil)
	// nilMap["key"] = 1  // This would PANIC!

	fmt.Println("\n=== Accessing Values ===")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Direct access
	fmt.Println("Red hex:", colors["red"])

	// Accessing non-existent key returns zero value
	fmt.Println("Yellow hex:", colors["yellow"]) // Returns ""

	// Check if key exists with "comma ok" idiom
	hex, exists := colors["yellow"]
	if exists {
		fmt.Println("Yellow exists:", hex)
	} else {
		fmt.Println("Yellow does not exist")
	}

	// Common pattern: check and use in one line
	if hex, ok := colors["green"]; ok {
		fmt.Println("Green found:", hex)
	}

	fmt.Println("\n=== Adding and Updating ===")

	inventory := make(map[string]int)

	// Adding new keys
	inventory["apples"] = 50
	inventory["bananas"] = 30
	inventory["oranges"] = 40
	fmt.Println("Initial inventory:", inventory)

	// Updating existing keys
	inventory["apples"] = 45 // Sold 5 apples
	fmt.Println("After selling apples:", inventory)

	// Increment pattern
	inventory["bananas"] += 10 // Restocked
	fmt.Println("After restocking bananas:", inventory)

	fmt.Println("\n=== Deleting Keys ===")

	pets := map[string]string{
		"dog":     "Buddy",
		"cat":     "Whiskers",
		"hamster": "Squeaky",
	}
	fmt.Println("Before delete:", pets)

	delete(pets, "hamster")
	fmt.Println("After delete:", pets)

	// Deleting non-existent key does nothing (no error)
	delete(pets, "fish")
	fmt.Println("After deleting non-existent:", pets)

	fmt.Println("\n=== Map Length ===")

	fruits := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	fmt.Println("Number of fruit types:", len(fruits))

	fmt.Println("\n=== Maps Are Reference Types ===")

	original := map[string]int{"a": 1, "b": 2}
	reference := original // Both point to same underlying data

	reference["a"] = 100
	reference["c"] = 3

	fmt.Println("Original:", original)   // Also changed!
	fmt.Println("Reference:", reference) // Same as original

	fmt.Println("\n=== Practical Example: Word Counter ===")

	wordCount := make(map[string]int)

	// Simple word splitting (for demo purposes)
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog", "the", "fox"}

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("Word counts:", wordCount)
	fmt.Println("'the' appears", wordCount["the"], "times")
}

// TO RUN: go run day5/01_maps_basics.go
//
// OUTPUT:
// === Map Declaration ===
// Ages map: map[Alice:30 Bob:25 Charlie:35]
// Capitals: map[France:Paris Germany:Berlin Italy:Rome Japan:Tokyo]
// ...
//
// EXERCISE:
// 1. Create a map of country codes to country names (e.g., "US" -> "United States")
// 2. Add at least 5 countries
// 3. Check if "UK" exists, if not add it
// 4. Print the total number of countries
//
// KEY POINTS:
// - Use make() or map literal to create maps
// - Zero value is nil (can't write to nil map)
// - Use "comma ok" idiom to check key existence
// - delete() removes keys, len() counts keys
// - Maps are reference types (copies share data)
