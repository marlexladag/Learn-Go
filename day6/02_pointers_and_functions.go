// Day 6, Exercise 2: Pointers and Functions
//
// Key concepts:
// - Pass by value: function gets a copy (default in Go)
// - Pass by pointer: function can modify the original
// - When to use pointers vs values in function parameters

package main

import "fmt"

func main() {
	fmt.Println("=== Pass by Value (Default) ===")

	num := 10
	fmt.Println("Before double:", num)
	doubleValue(num)
	fmt.Println("After double:", num) // Still 10!

	fmt.Println("\n=== Pass by Pointer ===")

	num2 := 10
	fmt.Println("Before double:", num2)
	doublePointer(&num2)
	fmt.Println("After double:", num2) // Now 20!

	fmt.Println("\n=== Multiple Values ===")

	a, b := 5, 10
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	fmt.Println("\n=== Returning Pointers ===")

	// Go allows returning pointers to local variables
	// The variable escapes to the heap automatically
	ptr := createInt(42)
	fmt.Println("Created int:", *ptr)

	// This is safe because Go's escape analysis
	// moves the variable to the heap when needed

	fmt.Println("\n=== Practical Example: Updating State ===")

	player := Player{
		Name:   "Hero",
		Health: 100,
		Score:  0,
	}

	fmt.Println("Initial:", player)

	takeDamage(&player, 30)
	fmt.Println("After damage:", player)

	addScore(&player, 100)
	fmt.Println("After scoring:", player)

	heal(&player, 20)
	fmt.Println("After healing:", player)

	fmt.Println("\n=== Optional Parameters with Pointers ===")

	// Using nil to represent "not provided"
	config := Config{}

	// Apply defaults, then overrides
	timeout := 30
	applyConfig(&config, &timeout, nil, nil)
	fmt.Printf("Config with timeout: %+v\n", config)

	retries := 5
	debug := true
	applyConfig(&config, nil, &retries, &debug)
	fmt.Printf("Config with retries and debug: %+v\n", config)

	fmt.Println("\n=== In-Place Modification ===")

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Before doubling slice:", numbers)
	doubleSlice(numbers) // Slices are already reference types!
	fmt.Println("After doubling slice:", numbers)

	// But slice header (length, capacity) needs pointer to modify
	fmt.Println("\nBefore append through pointer:")
	fmt.Println("Slice:", numbers)
	appendToSlice(&numbers, 6, 7, 8)
	fmt.Println("After append through pointer:", numbers)
}

// doubleValue receives a COPY - original unchanged
func doubleValue(n int) {
	n *= 2
	fmt.Println("Inside function:", n)
}

// doublePointer receives a pointer - can modify original
func doublePointer(n *int) {
	*n *= 2
	fmt.Println("Inside function:", *n)
}

// swap exchanges two values using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

// createInt demonstrates returning a pointer to a local variable
func createInt(value int) *int {
	result := value // local variable
	return &result  // safe - Go moves to heap
}

// Player represents a game player
type Player struct {
	Name   string
	Health int
	Score  int
}

func takeDamage(p *Player, damage int) {
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}

func heal(p *Player, amount int) {
	p.Health += amount
	if p.Health > 100 {
		p.Health = 100
	}
}

func addScore(p *Player, points int) {
	p.Score += points
}

// Config demonstrates optional parameters
type Config struct {
	Timeout int
	Retries int
	Debug   bool
}

func applyConfig(c *Config, timeout, retries *int, debug *bool) {
	if timeout != nil {
		c.Timeout = *timeout
	}
	if retries != nil {
		c.Retries = *retries
	}
	if debug != nil {
		c.Debug = *debug
	}
}

// doubleSlice modifies slice elements in place
// Note: slice itself is a reference type, so no pointer needed
// for element modification
func doubleSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// appendToSlice needs a pointer to modify the slice header
func appendToSlice(s *[]int, values ...int) {
	*s = append(*s, values...)
}

// TO RUN: go run day6/02_pointers_and_functions.go
//
// OUTPUT:
// === Pass by Value (Default) ===
// Before double: 10
// Inside function: 20
// After double: 10
// ...
//
// EXERCISE:
// 1. Write a function that takes a pointer to an int and triples it
// 2. Write a function that takes two int pointers and sets both to their sum
// 3. Create a simple Counter struct with an Increment method using pointer
//
// KEY POINTS:
// - Go is pass-by-value by default (copies are made)
// - Use pointers when you need to modify the original
// - Returning pointers to local variables is safe (escape analysis)
// - Slices are reference types but need pointers to modify header
