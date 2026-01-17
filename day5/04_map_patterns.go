// Day 5, Exercise 4: Common Map Patterns
//
// Key concepts:
// - Maps as sets (using map[T]bool or map[T]struct{})
// - Counting/frequency maps
// - Grouping data with maps
// - Caching/memoization patterns
// - Default values pattern

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Map as Set ===")

	// Using map[T]bool for set semantics
	uniqueNumbers := make(map[int]bool)

	numbers := []int{1, 2, 3, 2, 4, 1, 5, 3, 6}
	for _, n := range numbers {
		uniqueNumbers[n] = true
	}

	fmt.Println("Original:", numbers)
	fmt.Print("Unique: ")
	for n := range uniqueNumbers {
		fmt.Print(n, " ")
	}
	fmt.Println()

	// Check membership
	if uniqueNumbers[3] {
		fmt.Println("3 is in the set")
	}
	if !uniqueNumbers[10] {
		fmt.Println("10 is NOT in the set")
	}

	// Memory-efficient set using empty struct
	// struct{} takes 0 bytes!
	efficientSet := make(map[string]struct{})
	efficientSet["apple"] = struct{}{}
	efficientSet["banana"] = struct{}{}

	if _, exists := efficientSet["apple"]; exists {
		fmt.Println("apple is in efficient set")
	}

	fmt.Println("\n=== Counting / Frequency Map ===")

	text := "hello world hello go world go go"
	words := strings.Fields(text)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++ // Defaults to 0 if not exists
	}

	fmt.Println("Word frequencies:")
	for word, count := range wordCount {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// Find most frequent
	maxWord := ""
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}
	fmt.Printf("Most frequent: '%s' (%d times)\n", maxWord, maxCount)

	fmt.Println("\n=== Grouping Pattern ===")

	type Person struct {
		Name string
		City string
		Age  int
	}

	people := []Person{
		{Name: "Alice", City: "NYC", Age: 30},
		{Name: "Bob", City: "LA", Age: 25},
		{Name: "Charlie", City: "NYC", Age: 35},
		{Name: "Diana", City: "LA", Age: 28},
		{Name: "Eve", City: "Chicago", Age: 32},
	}

	// Group by city
	byCity := make(map[string][]Person)
	for _, p := range people {
		byCity[p.City] = append(byCity[p.City], p)
	}

	fmt.Println("People by city:")
	for city, residents := range byCity {
		fmt.Printf("  %s:\n", city)
		for _, p := range residents {
			fmt.Printf("    - %s (%d)\n", p.Name, p.Age)
		}
	}

	fmt.Println("\n=== Default Value Pattern ===")

	// Get with default using helper function
	settings := map[string]string{
		"theme":    "dark",
		"language": "en",
	}

	theme := getOrDefault(settings, "theme", "light")
	fontSize := getOrDefault(settings, "fontSize", "14")

	fmt.Println("Theme:", theme)        // dark (exists)
	fmt.Println("Font size:", fontSize) // 14 (default)

	fmt.Println("\n=== Two-Way Lookup ===")

	// Bidirectional mapping
	codeToName := map[string]string{
		"US": "United States",
		"UK": "United Kingdom",
		"FR": "France",
	}

	// Build reverse mapping
	nameToCode := make(map[string]string)
	for code, name := range codeToName {
		nameToCode[name] = code
	}

	fmt.Println("US ->", codeToName["US"])
	fmt.Println("France ->", nameToCode["France"])

	fmt.Println("\n=== Counting Duplicates ===")

	items := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}

	fmt.Println("Duplicates (count > 1):")
	for item, count := range counts {
		if count > 1 {
			fmt.Printf("  %s: %d times\n", item, count)
		}
	}

	fmt.Println("\n=== First/Last Occurrence Tracking ===")

	chars := "abracadabra"

	firstSeen := make(map[rune]int)
	lastSeen := make(map[rune]int)

	for i, ch := range chars {
		if _, exists := firstSeen[ch]; !exists {
			firstSeen[ch] = i
		}
		lastSeen[ch] = i
	}

	fmt.Printf("String: %s\n", chars)
	fmt.Println("First occurrence:")
	for ch, pos := range firstSeen {
		fmt.Printf("  '%c' at index %d\n", ch, pos)
	}

	fmt.Println("\n=== Intersection of Sets ===")

	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	intersection := make(map[int]bool)
	for k := range set1 {
		if set2[k] {
			intersection[k] = true
		}
	}

	fmt.Print("Set1: ")
	for k := range set1 {
		fmt.Print(k, " ")
	}
	fmt.Print("\nSet2: ")
	for k := range set2 {
		fmt.Print(k, " ")
	}
	fmt.Print("\nIntersection: ")
	for k := range intersection {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

// getOrDefault returns the value for key if it exists, otherwise returns defaultVal
func getOrDefault(m map[string]string, key, defaultVal string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return defaultVal
}

// TO RUN: go run day5/04_map_patterns.go
//
// OUTPUT:
// === Map as Set ===
// Original: [1 2 3 2 4 1 5 3 6]
// Unique: 1 2 3 4 5 6
// ...
//
// EXERCISE:
// 1. Given two slices of strings, find all elements that appear in both (intersection)
// 2. Find all elements that appear in only one slice (symmetric difference)
// 3. Remove duplicate strings from a slice while preserving order
//
// KEY POINTS:
// - map[T]bool or map[T]struct{} for set semantics
// - map[T]int is perfect for counting
// - map[T][]U for grouping data
// - Build reverse maps for bidirectional lookup
// - Missing keys return zero values (useful for counting)
