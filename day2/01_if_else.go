// Day 2, Exercise 1: If/Else Statements
//
// Key concepts:
// - if, else if, else
// - No parentheses around conditions
// - Braces {} are required
// - Statement initialization in if

package main

import "fmt"

func main() {
	age := 18

	// Basic if/else
	fmt.Println("=== Basic If/Else ===")
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// If/else if/else chain
	fmt.Println("\n=== Grade Classification ===")
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// If with initialization statement (Go special!)
	fmt.Println("\n=== If with Initialization ===")
	if num := 10; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
	// Note: 'num' only exists inside the if block!

	// Multiple conditions
	fmt.Println("\n=== Multiple Conditions ===")
	temperature := 72
	humidity := 45

	if temperature >= 65 && temperature <= 80 && humidity < 60 {
		fmt.Println("Perfect weather!")
	} else if temperature > 80 || humidity >= 60 {
		fmt.Println("It's uncomfortable")
	} else {
		fmt.Println("It's a bit cold")
	}

	// Nested if
	fmt.Println("\n=== Nested If ===")
	username := "admin"
	password := "secret123"

	if username == "admin" {
		if password == "secret123" {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Wrong password")
		}
	} else {
		fmt.Println("User not found")
	}

	// Comparing strings
	fmt.Println("\n=== String Comparison ===")
	name := "Go"
	if name == "Go" {
		fmt.Println("You're learning Go!")
	}
}

// TO RUN: go run 01_if_else.go
//
// EXERCISE: Write a program that:
// 1. Takes a year as input
// 2. Determines if it's a leap year
// Leap year rules:
// - Divisible by 4
// - BUT not by 100
// - UNLESS also divisible by 400
