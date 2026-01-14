// Day 1, Exercise 5: User Input
//
// Key concepts:
// - fmt.Scan, fmt.Scanf, fmt.Scanln for input
// - Pointers (&) to store input in variables
// - bufio.Scanner for line input

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Method 1: fmt.Scan (reads space-separated values)
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // & gives the memory address

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello %s! You are %d years old.\n\n", name, age)

	// Method 2: bufio.Scanner (reads full lines, better for strings with spaces)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your favorite quote: ")
	scanner.Scan() // reads the line
	quote := scanner.Text()

	fmt.Printf("Your quote: \"%s\"\n\n", quote)

	// Method 3: fmt.Scanf (formatted input)
	var month, day, year int
	fmt.Print("Enter date (MM/DD/YYYY): ")
	fmt.Scanf("%d/%d/%d", &month, &day, &year)
	fmt.Printf("Date entered: %02d/%02d/%d\n", month, day, year)
}

// TO RUN: go run 05_input.go
//
// NOTE: The & symbol is crucial - it passes the memory address
// so Scan can store the value in your variable.
// We'll learn more about pointers later!
//
// EXERCISE: Create a program that asks for:
// 1. First name
// 2. Last name
// 3. Birth year
// Then calculates and displays their age
