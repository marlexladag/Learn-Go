// Day 2, Exercise 5: Loop Patterns
//
// Practice with nested loops and patterns
// These exercises build loop intuition

package main

import "fmt"

func main() {
	// Pattern 1: Rectangle
	fmt.Println("=== Rectangle 5x3 ===")
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Pattern 2: Right Triangle
	fmt.Println("\n=== Right Triangle ===")
	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Pattern 3: Inverted Triangle
	fmt.Println("\n=== Inverted Triangle ===")
	for i := 5; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Pattern 4: Number Triangle
	fmt.Println("\n=== Number Triangle ===")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Pattern 5: Pyramid (centered)
	fmt.Println("\n=== Pyramid ===")
	rows := 5
	for i := 1; i <= rows; i++ {
		// Print leading spaces
		for j := 0; j < rows-i; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Pattern 6: Hollow Square
	fmt.Println("\n=== Hollow Square ===")
	size := 5
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	// Pattern 7: Diamond
	fmt.Println("\n=== Diamond ===")
	n := 5
	// Upper half
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// Lower half
	for i := n - 1; i >= 1; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// TO RUN: go run 05_patterns.go
//
// EXERCISE: Create your own pattern:
// 1. Checkerboard pattern (alternating * and spaces)
// 2. Letter pattern (like the letter X or H)
