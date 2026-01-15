// Day 4, Exercise 2: Slice Basics
//
// Key concepts:
// - Slices are dynamic, flexible views into arrays
// - Slices have length (current elements) and capacity (max elements before reallocation)
// - Slices are reference types (point to underlying array)
// - make() creates slices with specified length and capacity
// - nil slices vs empty slices

package main

import "fmt"

func main() {
	fmt.Println("=== Creating Slices ===")

	// Method 1: Slice literal (most common)
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice literal:", fruits)

	// Method 2: make(type, length, capacity)
	numbers := make([]int, 3, 5) // length=3, capacity=5
	fmt.Println("make() slice:", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Method 3: make(type, length) - capacity = length
	scores := make([]int, 5)
	fmt.Println("make() with length only:", scores)

	// Method 4: Empty slice
	empty := []int{}
	fmt.Println("Empty slice:", empty, "Length:", len(empty))

	// Method 5: nil slice (declared but not initialized)
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice, "Is nil:", nilSlice == nil)

	fmt.Println("\n=== Slice vs Array ===")

	// This is an ARRAY (fixed size)
	array := [3]int{1, 2, 3}

	// This is a SLICE (dynamic)
	slice := []int{1, 2, 3}

	fmt.Printf("Array type: %T\n", array) // [3]int
	fmt.Printf("Slice type: %T\n", slice) // []int

	fmt.Println("\n=== Slicing Arrays and Slices ===")

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", original)

	// slice[start:end] - includes start, excludes end
	fmt.Println("original[2:5]:", original[2:5])   // [2 3 4]
	fmt.Println("original[:3]:", original[:3])     // [0 1 2] - from start
	fmt.Println("original[7:]:", original[7:])     // [7 8 9] - to end
	fmt.Println("original[:]:", original[:])       // [0 1 2 3 4 5 6 7 8 9] - full copy reference

	// Slice from array
	arr := [5]int{10, 20, 30, 40, 50}
	sliceFromArr := arr[1:4]
	fmt.Println("Slice from array:", sliceFromArr) // [20 30 40]

	fmt.Println("\n=== Slices Share Underlying Array ===")

	data := []int{1, 2, 3, 4, 5}
	slice1 := data[1:4]  // [2 3 4]
	slice2 := data[2:5]  // [3 4 5]

	fmt.Println("Before modification:")
	fmt.Println("data:", data)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	// Modify slice1
	slice1[1] = 100 // This changes data[2]

	fmt.Println("\nAfter slice1[1] = 100:")
	fmt.Println("data:", data)     // [1 2 100 4 5] - changed!
	fmt.Println("slice1:", slice1) // [2 100 4]
	fmt.Println("slice2:", slice2) // [100 4 5] - also changed!

	fmt.Println("\n=== Appending to Slices ===")

	colors := []string{"red", "green"}
	fmt.Println("Original:", colors)

	// append() returns a NEW slice
	colors = append(colors, "blue")
	fmt.Println("After append:", colors)

	// Append multiple elements
	colors = append(colors, "yellow", "purple")
	fmt.Println("After multi-append:", colors)

	// Append one slice to another using ...
	moreColors := []string{"orange", "pink"}
	colors = append(colors, moreColors...)
	fmt.Println("After slice append:", colors)

	fmt.Println("\n=== Nil vs Empty Slice Behavior ===")

	var nilS []int       // nil slice
	emptyS := []int{}    // empty slice
	makeS := make([]int, 0) // empty slice via make

	fmt.Printf("nil slice: %v, len=%d, is nil=%t\n", nilS, len(nilS), nilS == nil)
	fmt.Printf("empty slice: %v, len=%d, is nil=%t\n", emptyS, len(emptyS), emptyS == nil)
	fmt.Printf("make slice: %v, len=%d, is nil=%t\n", makeS, len(makeS), makeS == nil)

	// All work the same with append!
	nilS = append(nilS, 1)
	emptyS = append(emptyS, 1)
	makeS = append(makeS, 1)

	fmt.Println("After append - all work:", nilS, emptyS, makeS)

	fmt.Println("\n=== Practical Example: Building a List ===")

	// Start with nil slice (idiomatic Go)
	var todoList []string

	// Add items
	todoList = append(todoList, "Learn arrays")
	todoList = append(todoList, "Learn slices")
	todoList = append(todoList, "Practice coding")

	fmt.Println("Todo List:")
	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Println("\n=== Copying Slices (Safe Copy) ===")

	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))

	copied := copy(dest, source) // Returns number of elements copied
	fmt.Printf("Copied %d elements\n", copied)

	dest[0] = 100
	fmt.Println("Source (unchanged):", source)
	fmt.Println("Dest (modified):", dest)
}

// TO RUN: go run day4/02_slices_basics.go
//
// OUTPUT:
// === Creating Slices ===
// Slice literal: [apple banana cherry]
// make() slice: [0 0 0]
// Length: 3, Capacity: 5
// ...
//
// EXERCISE:
// 1. Create a slice of your hobbies
// 2. Add two more hobbies using append
// 3. Create a sub-slice of your top 3 hobbies
// 4. Make a safe copy and modify the copy
//
// KEY POINTS:
// - Slices are dynamic (can grow with append)
// - Slices are REFERENCE types (share underlying array)
// - Use make() when you know the size upfront
// - append() may return a new slice (always reassign!)
// - Use copy() for independent copies
