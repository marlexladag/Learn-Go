// Day 7, Exercise 1: Review Quiz
//
// Test your knowledge of Days 1-6!
// Run this program and answer the questions in the comments.

package main

import "fmt"

func main() {
	fmt.Println("=== Day 7: Go Foundations Review ===\n")

	// QUIZ 1: Variables & Types
	fmt.Println("--- Quiz 1: Variables & Types ---")

	// What will this print?
	var a int
	b := 3.14
	c := "hello"
	fmt.Printf("a=%v (type: %T)\n", a, a)
	fmt.Printf("b=%v (type: %T)\n", b, b)
	fmt.Printf("c=%v (type: %T)\n", c, c)
	// Answer: a=0 (zero value), b=3.14 (float64), c=hello (string)

	// QUIZ 2: Control Flow
	fmt.Println("\n--- Quiz 2: Control Flow ---")

	// What numbers will this print?
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	// Answer: 1 3 (odd numbers only, continue skips even)

	// QUIZ 3: Functions
	fmt.Println("\n--- Quiz 3: Functions ---")

	// What will swap return?
	x, y := swap(10, 20)
	fmt.Printf("x=%d, y=%d\n", x, y)
	// Answer: x=20, y=10

	// QUIZ 4: Slices
	fmt.Println("\n--- Quiz 4: Slices ---")

	// What happens here?
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:4]
	slice[0] = 100
	fmt.Println("original:", original)
	fmt.Println("slice:", slice)
	// Answer: original changes too! Slices share underlying array

	// QUIZ 5: Maps
	fmt.Println("\n--- Quiz 5: Maps ---")

	// What will this print?
	scores := map[string]int{"Alice": 95}
	val, exists := scores["Bob"]
	fmt.Printf("val=%d, exists=%v\n", val, exists)
	// Answer: val=0 (zero value), exists=false

	// QUIZ 6: Pointers
	fmt.Println("\n--- Quiz 6: Pointers ---")

	// What will num be after this?
	num := 5
	double(&num)
	fmt.Printf("num=%d\n", num)
	// Answer: num=10 (modified through pointer)

	fmt.Println("\n=== Summary of Key Concepts ===")
	fmt.Println(`
Day 1 - Basics:
  - var x int vs x := value (short declaration)
  - Zero values: 0, "", false, nil
  - Constants are immutable

Day 2 - Control Flow:
  - if/else, no parentheses needed
  - for is the only loop (no while)
  - switch doesn't need break

Day 3 - Functions:
  - func name(params) returnType
  - Multiple return values
  - Variadic: func(args ...int)

Day 4 - Arrays & Slices:
  - Arrays: fixed size [5]int
  - Slices: dynamic []int
  - append(), len(), cap()

Day 5 - Maps:
  - map[KeyType]ValueType
  - val, ok := m[key] (comma-ok idiom)
  - delete(m, key)

Day 6 - Pointers:
  - & gets address
  - * dereferences
  - Pass by pointer to modify
`)
}

func swap(a, b int) (int, int) {
	return b, a
}

func double(n *int) {
	*n *= 2
}

// TO RUN: go run day7/01_review_quiz.go
//
// SELF-CHECK:
// Before running, predict each output!
// Understanding WHY is more important than memorizing.
