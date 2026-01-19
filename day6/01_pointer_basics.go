// Day 6, Exercise 1: Pointer Basics
//
// Key concepts:
// - A pointer holds the memory address of a value
// - The & operator gets the address of a variable
// - The * operator dereferences a pointer (gets the value at the address)
// - Zero value of a pointer is nil

package main

import "fmt"

func main() {
	fmt.Println("=== What is a Pointer? ===")

	// A regular variable holds a value
	x := 42
	fmt.Println("Value of x:", x)

	// A pointer holds the memory address of that value
	p := &x // & means "address of"
	fmt.Println("Address of x:", p)
	fmt.Println("Type of p:", fmt.Sprintf("%T", p)) // *int

	// Dereferencing: get the value at the address
	fmt.Println("Value at address p:", *p) // * means "value at"

	fmt.Println("\n=== Modifying Through Pointers ===")

	num := 10
	ptr := &num

	fmt.Println("Before:", num)
	*ptr = 20                  // Change value through pointer
	fmt.Println("After:", num) // Original changed!

	fmt.Println("\n=== Pointer Declaration ===")

	// Method 1: Declare and assign later
	var intPtr *int // nil pointer
	fmt.Println("Nil pointer:", intPtr)
	fmt.Println("Is nil:", intPtr == nil)

	value := 100
	intPtr = &value
	fmt.Println("After assignment:", *intPtr)

	// Method 2: Short declaration with &
	name := "Go"
	namePtr := &name
	fmt.Println("Name pointer points to:", *namePtr)

	// Method 3: Using new() - allocates memory, returns pointer
	newPtr := new(int)                      // Allocates int, returns *int
	fmt.Println("new(int) value:", *newPtr) // Zero value: 0
	*newPtr = 42
	fmt.Println("After assignment:", *newPtr)

	fmt.Println("\n=== Pointer vs Value ===")

	a := 5
	b := a  // b is a COPY of a
	c := &a // c points to a

	fmt.Println("a:", a, "b:", b, "*c:", *c)

	b = 10  // Only changes b
	*c = 15 // Changes a through pointer

	fmt.Println("After changes:")
	fmt.Println("a:", a, "b:", b, "*c:", *c)

	fmt.Println("\n=== Why Use Pointers? ===")

	// 1. To modify the original value
	counter := 0
	increment(&counter)
	increment(&counter)
	fmt.Println("Counter after increments:", counter)

	// 2. To avoid copying large data structures
	// (We'll see more of this with structs)

	// 3. To signal "no value" with nil
	result := findValue([]int{1, 2, 3}, 2)
	if result != nil {
		fmt.Println("Found value:", *result)
	}

	result = findValue([]int{1, 2, 3}, 5)
	if result == nil {
		fmt.Println("Value not found")
	}

	fmt.Println("\n=== Pointer Arithmetic (NOT supported) ===")

	// Unlike C/C++, Go does NOT support pointer arithmetic
	arr := [3]int{10, 20, 30}
	arrPtr := &arr[0]
	fmt.Println("First element:", *arrPtr)
	// arrPtr++ // This would NOT compile in Go!
	// This is intentional - Go is memory-safe

	fmt.Println("\n=== Common Mistake: Dereferencing Nil ===")

	var nilPtr *int
	// fmt.Println(*nilPtr) // PANIC! nil pointer dereference

	// Always check for nil before dereferencing
	if nilPtr != nil {
		fmt.Println(*nilPtr)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}
}

// increment modifies the original value through a pointer
func increment(n *int) {
	*n++
}

// findValue returns a pointer to the found value, or nil if not found
func findValue(slice []int, target int) *int {
	for i := range slice {
		if slice[i] == target {
			return &slice[i]
		}
	}
	return nil
}

// TO RUN: go run day6/01_pointer_basics.go
//
// OUTPUT:
// === What is a Pointer? ===
// Value of x: 42
// Address of x: 0xc0000b4008 (address will vary)
// Type of p: *int
// Value at address p: 42
// ...
//
// EXERCISE:
// 1. Create an int variable and a pointer to it
// 2. Print both the value and the address
// 3. Modify the value through the pointer
// 4. Verify the original variable changed
//
// KEY POINTS:
// - & gets the address of a variable
// - * dereferences a pointer (gets value)
// - Pointers allow modifying original values
// - Zero value is nil - always check before dereferencing
// - Go does NOT support pointer arithmetic (memory safety)
