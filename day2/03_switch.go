// Day 2, Exercise 3: Switch Statements
//
// Key concepts:
// - No 'break' needed (no fall-through by default)
// - Multiple values per case
// - Expressionless switch (like if-else chain)
// - Type switch (for interfaces - we'll cover later)

package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch
	fmt.Println("=== Basic Switch ===")
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7: // Multiple values in one case
		fmt.Println("Weekend!")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with strings
	fmt.Println("\n=== Switch with Strings ===")
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("Apples are red or green")
	case "banana":
		fmt.Println("Bananas are yellow")
	case "orange", "tangerine":
		fmt.Println("Citrus fruits!")
	default:
		fmt.Println("Unknown fruit")
	}

	// Switch with initialization
	fmt.Println("\n=== Switch with Init ===")
	switch num := 15; { // Note: no expression after switch
	case num < 0:
		fmt.Println("Negative")
	case num == 0:
		fmt.Println("Zero")
	case num < 10:
		fmt.Println("Single digit")
	case num < 100:
		fmt.Println("Double digit")
	default:
		fmt.Println("Large number")
	}

	// Expressionless switch (like if-else chain)
	fmt.Println("\n=== Expressionless Switch ===")
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Excellent! Grade A")
	case score >= 80:
		fmt.Println("Good! Grade B")
	case score >= 70:
		fmt.Println("Average. Grade C")
	case score >= 60:
		fmt.Println("Below average. Grade D")
	default:
		fmt.Println("Failed. Grade F")
	}

	// Fallthrough keyword (explicit fall-through)
	fmt.Println("\n=== Fallthrough ===")
	num := 5

	switch num {
	case 5:
		fmt.Println("Five")
		fallthrough // Continue to next case
	case 4:
		fmt.Println("Four or came from five")
		fallthrough
	case 3:
		fmt.Println("Three or came from above")
	case 2:
		fmt.Println("Two")
	}

	// Practical example: Time-based greeting
	fmt.Println("\n=== Time-Based Greeting ===")
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	case hour < 21:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night!")
	}
}

// TO RUN: go run 03_switch.go
//
// EXERCISE: Create a simple command parser that takes a command
// and responds accordingly:
// - "help" -> prints available commands
// - "version" -> prints "v1.0.0"
// - "quit", "exit", "q" -> prints "Goodbye!"
// - anything else -> prints "Unknown command"
