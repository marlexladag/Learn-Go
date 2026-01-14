// Day 2, Exercise 2: For Loops
//
// Key concepts:
// - Go has ONLY the 'for' loop (no while, do-while)
// - Three forms: standard, while-style, infinite
// - break and continue keywords
// - range for iterating

package main

import "fmt"

func main() {
	// Standard for loop (like C, Java, JavaScript)
	fmt.Println("=== Standard For Loop ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// While-style loop (condition only)
	fmt.Println("\n=== While-Style Loop ===")
	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	// Infinite loop (with break)
	fmt.Println("\n=== Infinite Loop with Break ===")
	n := 0
	for {
		fmt.Printf("n = %d\n", n)
		n++
		if n >= 3 {
			break // exit the loop
		}
	}

	// Continue keyword
	fmt.Println("\n=== Continue (skip even numbers) ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip to next iteration
		}
		fmt.Printf("%d is odd\n", i)
	}

	// Looping with range over a string
	fmt.Println("\n=== Range over String ===")
	word := "Go!"
	for index, char := range word {
		fmt.Printf("Index %d: %c (Unicode: %d)\n", index, char, char)
	}

	// Ignoring index with blank identifier
	fmt.Println("\n=== Ignoring Index ===")
	for _, char := range "Hello" {
		fmt.Printf("%c ", char)
	}
	fmt.Println()

	// Nested loops
	fmt.Println("\n=== Nested Loops (Multiplication Table) ===")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// Loop with label (for breaking out of nested loops)
	fmt.Println("\n=== Break with Label ===")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of both loops!")
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}

// TO RUN: go run 02_for_loops.go
//
// EXERCISE 1: Print the first 10 Fibonacci numbers
// (0, 1, 1, 2, 3, 5, 8, 13, 21, 34)
//
// EXERCISE 2: Find the sum of all numbers from 1 to 100
