// Day 4, Exercise 4: Slice Internals
//
// Key concepts:
// - Slice header: pointer, length, capacity
// - How capacity affects append behavior
// - Memory allocation and growth strategy
// - Avoiding common slice gotchas
// - Full slice expressions s[low:high:max]

package main

import "fmt"

func main() {
	fmt.Println("=== Slice Header Structure ===")
	// A slice is a struct with 3 fields:
	// - Pointer to underlying array
	// - Length (number of accessible elements)
	// - Capacity (max elements before reallocation)

	s := make([]int, 3, 5)
	fmt.Printf("Slice: %v\n", s)
	fmt.Printf("Length: %d (elements you can access)\n", len(s))
	fmt.Printf("Capacity: %d (elements before reallocation)\n", cap(s))

	fmt.Println("\n=== Length vs Capacity ===")

	data := make([]int, 2, 5)
	data[0], data[1] = 10, 20

	fmt.Println("Initial:", data)
	fmt.Printf("len=%d, cap=%d\n", len(data), cap(data))

	// Append within capacity (no reallocation)
	data = append(data, 30)
	fmt.Println("After append 30:", data)
	fmt.Printf("len=%d, cap=%d\n", len(data), cap(data))

	// Keep appending
	data = append(data, 40, 50)
	fmt.Println("After append 40,50:", data)
	fmt.Printf("len=%d, cap=%d (still same capacity)\n", len(data), cap(data))

	// Now capacity is reached, next append triggers reallocation
	data = append(data, 60)
	fmt.Println("After append 60:", data)
	fmt.Printf("len=%d, cap=%d (capacity grew!)\n", len(data), cap(data))

	fmt.Println("\n=== Capacity Growth Strategy ===")

	var growth []int
	prevCap := 0
	for i := 0; i < 20; i++ {
		growth = append(growth, i)
		if cap(growth) != prevCap {
			fmt.Printf("len=%2d, cap=%2d (grew from %d)\n", len(growth), cap(growth), prevCap)
			prevCap = cap(growth)
		}
	}
	// Note: Go typically doubles capacity for small slices,
	// then grows by ~25% for larger ones

	fmt.Println("\n=== Slicing and Capacity ===")

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len=%d, cap=%d\n", original, len(original), cap(original))

	// Slice from index 2
	slice1 := original[2:5]
	fmt.Printf("original[2:5]: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	// Capacity is 8 because it can extend to end of original!

	// Slice from index 7
	slice2 := original[7:]
	fmt.Printf("original[7:]:  %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	fmt.Println("\n=== The Slice Gotcha: Shared Backing Array ===")

	base := []int{1, 2, 3, 4, 5}
	sub := base[1:3] // [2, 3]

	fmt.Println("Before append:")
	fmt.Printf("base: %v\n", base)
	fmt.Printf("sub:  %v (cap=%d)\n", sub, cap(sub))

	// Append to sub - it has capacity, so it overwrites base!
	sub = append(sub, 100)

	fmt.Println("\nAfter sub = append(sub, 100):")
	fmt.Printf("base: %v  <-- element 4 became 100!\n", base)
	fmt.Printf("sub:  %v\n", sub)

	fmt.Println("\n=== Full Slice Expression: Limiting Capacity ===")

	// s[low:high:max] - max limits the capacity
	base2 := []int{1, 2, 3, 4, 5}

	// Normal slice (inherits remaining capacity)
	normal := base2[1:3]
	fmt.Printf("base2[1:3]:   %v, cap=%d\n", normal, cap(normal))

	// Full slice expression (limited capacity)
	limited := base2[1:3:3] // capacity = max - low = 3 - 1 = 2
	fmt.Printf("base2[1:3:3]: %v, cap=%d\n", limited, cap(limited))

	// Now append forces a new backing array
	base3 := []int{1, 2, 3, 4, 5}
	safe := base3[1:3:3]
	safe = append(safe, 100) // Creates new array!

	fmt.Println("\nWith full slice expression:")
	fmt.Printf("base3: %v <-- unchanged!\n", base3)
	fmt.Printf("safe:  %v\n", safe)

	fmt.Println("\n=== Pre-allocating for Performance ===")

	// Bad: Grows multiple times
	var slow []int
	for i := 0; i < 10000; i++ {
		slow = append(slow, i)
	}

	// Good: Pre-allocate when size is known
	fast := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		fast = append(fast, i)
	}

	// Best: If exact size known, use length
	best := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		best[i] = i
	}

	fmt.Println("Pre-allocation prevents multiple reallocations")
	fmt.Println("Use make([]T, 0, n) when appending to known size")
	fmt.Println("Use make([]T, n) when setting by index")

	fmt.Println("\n=== Nil Slice vs Empty Slice ===")

	var nilSlice []int          // nil, no allocation
	emptySlice := []int{}       // empty, has allocation
	makeSlice := make([]int, 0) // empty, has allocation

	fmt.Printf("nil slice:   %v, len=%d, cap=%d, is nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("empty slice: %v, len=%d, cap=%d, is nil: %t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("make slice:  %v, len=%d, cap=%d, is nil: %t\n",
		makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)

	// Prefer nil slices - they work the same but use less memory
	fmt.Println("\nTip: var s []T is preferred over s := []T{}")

	fmt.Println("\n=== Clearing a Slice ===")

	// Method 1: Set to nil (releases memory)
	s1 := []int{1, 2, 3, 4, 5}
	s1 = nil
	fmt.Println("Set to nil:", s1, "cap:", cap(s1))

	// Method 2: Re-slice to zero length (keeps capacity)
	s2 := []int{1, 2, 3, 4, 5}
	s2 = s2[:0]
	fmt.Println("Re-slice [:0]:", s2, "cap:", cap(s2))

	// Method 3: clear() function (Go 1.21+) - zeros elements, keeps length
	s3 := []int{1, 2, 3, 4, 5}
	clear(s3)
	fmt.Println("clear():", s3, "cap:", cap(s3))
}

// TO RUN: go run day4/04_slice_internals.go
//
// OUTPUT:
// === Slice Header Structure ===
// Slice: [0 0 0]
// Length: 3 (elements you can access)
// Capacity: 5 (elements before reallocation)
// ...
//
// EXERCISE:
// 1. Create a slice with length 2 and capacity 10
// 2. Append 8 more elements and verify no reallocation
// 3. Create a sub-slice and use full slice expression to prevent gotchas
// 4. Benchmark: compare append with/without pre-allocation
//
// KEY POINTS:
// - Slice = pointer + length + capacity
// - Appending may or may not reallocate
// - Sub-slices share the backing array (gotcha!)
// - Use full slice expression s[low:high:max] to limit capacity
// - Pre-allocate with make() for better performance
// - Prefer nil slices over empty slice literals
