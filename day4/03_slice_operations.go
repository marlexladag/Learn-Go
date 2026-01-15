// Day 4, Exercise 3: Slice Operations
//
// Key concepts:
// - Common slice manipulation patterns
// - Removing elements from slices
// - Inserting elements into slices
// - Filtering and transforming slices
// - Reversing and sorting slices

package main

import (
	"fmt"
	"slices" // Go 1.21+ standard library
)

func main() {
	fmt.Println("=== Removing Elements ===")

	// Remove element at index i
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", nums)

	i := 2 // Remove element at index 2 (value 3)
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println("After removing index 2:", nums) // [1 2 4 5]

	// Remove first element
	letters := []string{"a", "b", "c", "d"}
	letters = letters[1:]
	fmt.Println("Remove first:", letters) // [b c d]

	// Remove last element
	letters = letters[:len(letters)-1]
	fmt.Println("Remove last:", letters) // [b c]

	fmt.Println("\n=== Inserting Elements ===")

	// Insert at index
	data := []int{1, 2, 4, 5}
	fmt.Println("Before insert:", data)

	// Insert 3 at index 2
	insertIdx := 2
	insertVal := 3
	data = append(data[:insertIdx], append([]int{insertVal}, data[insertIdx:]...)...)
	fmt.Println("After insert 3 at index 2:", data) // [1 2 3 4 5]

	// Cleaner insert using slices package (Go 1.21+)
	data2 := []int{1, 2, 4, 5}
	data2 = slices.Insert(data2, 2, 3)
	fmt.Println("Using slices.Insert:", data2)

	fmt.Println("\n=== Filtering Slices ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original:", numbers)

	// Filter: keep only even numbers
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Println("Evens only:", evens)

	// Filter in-place (modifies original, more memory efficient)
	odds := numbers[:0] // Reuse backing array
	for _, n := range numbers {
		if n%2 != 0 {
			odds = append(odds, n)
		}
	}
	fmt.Println("Odds (in-place):", odds)

	fmt.Println("\n=== Transforming Slices (Map pattern) ===")

	prices := []float64{10.0, 20.0, 30.0, 40.0}
	fmt.Println("Original prices:", prices)

	// Apply 10% discount
	discounted := make([]float64, len(prices))
	for i, price := range prices {
		discounted[i] = price * 0.9
	}
	fmt.Println("Discounted prices:", discounted)

	// Square all numbers
	nums2 := []int{1, 2, 3, 4, 5}
	squared := make([]int, len(nums2))
	for i, n := range nums2 {
		squared[i] = n * n
	}
	fmt.Println("Squared:", squared)

	fmt.Println("\n=== Reversing Slices ===")

	forward := []int{1, 2, 3, 4, 5}
	fmt.Println("Forward:", forward)

	// Manual reverse (in-place)
	reversed := make([]int, len(forward))
	copy(reversed, forward)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	fmt.Println("Reversed:", reversed)

	// Using slices.Reverse (Go 1.21+)
	toReverse := []string{"a", "b", "c", "d"}
	slices.Reverse(toReverse)
	fmt.Println("slices.Reverse:", toReverse)

	fmt.Println("\n=== Sorting Slices ===")

	unsorted := []int{5, 2, 8, 1, 9, 3}
	fmt.Println("Unsorted:", unsorted)

	// Using slices.Sort (Go 1.21+)
	slices.Sort(unsorted)
	fmt.Println("Sorted:", unsorted)

	// Check if sorted
	fmt.Println("Is sorted:", slices.IsSorted(unsorted))

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	slices.Sort(words)
	fmt.Println("Sorted strings:", words)

	fmt.Println("\n=== Finding Elements ===")

	items := []string{"apple", "banana", "cherry", "date"}

	// Check if element exists (Go 1.21+)
	fmt.Println("Contains 'banana':", slices.Contains(items, "banana"))
	fmt.Println("Contains 'grape':", slices.Contains(items, "grape"))

	// Find index (Go 1.21+)
	idx := slices.Index(items, "cherry")
	fmt.Println("Index of 'cherry':", idx)

	// Manual search (works on any Go version)
	target := "banana"
	foundIdx := -1
	for i, item := range items {
		if item == target {
			foundIdx = i
			break
		}
	}
	fmt.Println("Manual search - 'banana' at index:", foundIdx)

	fmt.Println("\n=== Removing Duplicates ===")

	withDups := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	fmt.Println("With duplicates:", withDups)

	// Using slices.Compact (requires sorted slice, Go 1.21+)
	unique := slices.Compact(withDups)
	fmt.Println("Unique (compact):", unique)

	// Manual dedup (preserves order, any Go version)
	input := []string{"a", "b", "a", "c", "b", "d"}
	seen := make(map[string]bool)
	var deduped []string
	for _, item := range input {
		if !seen[item] {
			seen[item] = true
			deduped = append(deduped, item)
		}
	}
	fmt.Println("Original:", input)
	fmt.Println("Deduped:", deduped)

	fmt.Println("\n=== Practical Example: Stack Operations ===")

	var stack []int

	// Push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println("Stack after pushes:", stack)

	// Peek (look at top without removing)
	top := stack[len(stack)-1]
	fmt.Println("Peek:", top)

	// Pop (remove and return top)
	top, stack = stack[len(stack)-1], stack[:len(stack)-1]
	fmt.Println("Popped:", top)
	fmt.Println("Stack after pop:", stack)

	fmt.Println("\n=== Practical Example: Queue Operations ===")

	var queue []string

	// Enqueue (add to back)
	queue = append(queue, "first")
	queue = append(queue, "second")
	queue = append(queue, "third")
	fmt.Println("Queue:", queue)

	// Dequeue (remove from front)
	front := queue[0]
	queue = queue[1:]
	fmt.Println("Dequeued:", front)
	fmt.Println("Queue after dequeue:", queue)
}

// TO RUN: go run day4/03_slice_operations.go
//
// Note: This file uses the 'slices' package from Go 1.21+
// If using an older version, use the manual implementations shown
//
// OUTPUT:
// === Removing Elements ===
// Original: [1 2 3 4 5]
// After removing index 2: [1 2 4 5]
// ...
//
// EXERCISE:
// 1. Create a slice of 10 random numbers
// 2. Remove all numbers less than 5
// 3. Sort the remaining numbers in descending order
// 4. Find and remove any duplicates
//
// KEY POINTS:
// - append() is the key function for slice manipulation
// - Remove: append(s[:i], s[i+1:]...)
// - Insert: append(s[:i], append([]T{v}, s[i:]...)...)
// - The 'slices' package (Go 1.21+) provides useful helpers
// - Always be mindful of the underlying array when modifying
