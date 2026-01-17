// Day 5, Exercise 5: Maps and Functions
//
// Key concepts:
// - Passing maps to functions (reference semantics)
// - Returning maps from functions
// - Maps as function parameters are passed by reference
// - Modifying maps in functions affects the original

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Passing Maps to Functions ===")

	scores := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
	}

	fmt.Println("Before:", scores)

	// Function modifies the original map!
	addBonus(scores, 5)
	fmt.Println("After adding bonus:", scores)

	fmt.Println("\n=== Returning Maps from Functions ===")

	text := "the quick brown fox jumps over the lazy dog"
	wordFreq := countWords(text)
	fmt.Println("Word frequencies:", wordFreq)

	fmt.Println("\n=== Map as Configuration ===")

	config := map[string]string{
		"host":     "localhost",
		"port":     "8080",
		"protocol": "https",
	}

	url := buildURL(config)
	fmt.Println("Built URL:", url)

	fmt.Println("\n=== Merging Maps ===")

	defaults := map[string]int{
		"timeout":     30,
		"maxRetries":  3,
		"bufferSize":  1024,
		"concurrency": 4,
	}

	userSettings := map[string]int{
		"timeout":    60,
		"bufferSize": 2048,
	}

	merged := mergeMaps(defaults, userSettings)
	fmt.Println("Default settings:", defaults)
	fmt.Println("User settings:", userSettings)
	fmt.Println("Merged settings:", merged)

	fmt.Println("\n=== Filtering Maps ===")

	inventory := map[string]int{
		"apple":  50,
		"banana": 5,
		"orange": 30,
		"grape":  2,
		"mango":  15,
	}

	lowStock := filterMap(inventory, func(k string, v int) bool {
		return v < 10
	})

	fmt.Println("Full inventory:", inventory)
	fmt.Println("Low stock items:", lowStock)

	fmt.Println("\n=== Transforming Maps ===")

	prices := map[string]float64{
		"laptop":   999.99,
		"mouse":    29.99,
		"keyboard": 79.99,
	}

	// Apply 20% discount
	discounted := transformValues(prices, func(price float64) float64 {
		return price * 0.8
	})

	fmt.Println("Original prices:", prices)
	fmt.Println("Discounted prices:", discounted)

	fmt.Println("\n=== Inverting a Map ===")

	countryCapital := map[string]string{
		"France":  "Paris",
		"Japan":   "Tokyo",
		"Germany": "Berlin",
	}

	capitalCountry := invertMap(countryCapital)

	fmt.Println("Country -> Capital:", countryCapital)
	fmt.Println("Capital -> Country:", capitalCountry)

	fmt.Println("\n=== Checking Map Equality ===")

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	map3 := map[string]int{"a": 1, "b": 2, "c": 4}

	fmt.Println("map1 == map2:", mapsEqual(map1, map2))
	fmt.Println("map1 == map3:", mapsEqual(map1, map3))

	fmt.Println("\n=== Copying Maps (Deep Copy) ===")

	original := map[string]int{"x": 10, "y": 20}
	copied := copyMap(original)

	copied["x"] = 100
	copied["z"] = 30

	fmt.Println("Original:", original) // Unchanged
	fmt.Println("Copied:", copied)     // Modified
}

// addBonus modifies the map in place
func addBonus(scores map[string]int, bonus int) {
	for name := range scores {
		scores[name] += bonus
	}
}

// countWords returns a frequency map of words in text
func countWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		result[word]++
	}
	return result
}

// buildURL creates a URL from configuration map
func buildURL(config map[string]string) string {
	protocol := config["protocol"]
	host := config["host"]
	port := config["port"]
	return fmt.Sprintf("%s://%s:%s", protocol, host, port)
}

// mergeMaps combines two maps, with the second taking precedence
func mergeMaps(base, override map[string]int) map[string]int {
	result := make(map[string]int)

	// Copy base
	for k, v := range base {
		result[k] = v
	}

	// Override with second map
	for k, v := range override {
		result[k] = v
	}

	return result
}

// filterMap returns a new map with only entries that pass the predicate
func filterMap(m map[string]int, predicate func(string, int) bool) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// transformValues applies a function to all values
func transformValues(m map[string]float64, transform func(float64) float64) map[string]float64 {
	result := make(map[string]float64)
	for k, v := range m {
		result[k] = transform(v)
	}
	return result
}

// invertMap swaps keys and values
func invertMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[v] = k
	}
	return result
}

// mapsEqual checks if two maps have the same key-value pairs
func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

// copyMap creates a shallow copy of a map
func copyMap(m map[string]int) map[string]int {
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// TO RUN: go run day5/05_maps_and_functions.go
//
// OUTPUT:
// === Passing Maps to Functions ===
// Before: map[Alice:85 Bob:90 Charlie:78]
// After adding bonus: map[Alice:90 Bob:95 Charlie:83]
// ...
//
// EXERCISE:
// 1. Write a function that takes a map[string]int and returns the key with the highest value
// 2. Write a function that removes all entries where value is below a threshold
// 3. Write a function that returns only keys that exist in both of two maps
//
// KEY POINTS:
// - Maps are reference types - functions modify originals
// - Return new maps when you want immutability
// - Helper functions make map operations reusable
// - Filter, transform, merge are common patterns
// - Deep copy required to avoid shared references
