// Day 1, Exercise 3: Constants
//
// Key concepts:
// - const keyword for values that don't change
// - Constants must be known at compile time
// - iota for auto-incrementing constants

package main

import "fmt"

// Package-level constants
const Pi = 3.14159
const AppName = "GoLearner"

// Grouped constants
const (
	MaxUsers    = 100
	MinPassword = 8
	Version     = "1.0.0"
)

// iota - auto-incrementing constant generator
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// Practical iota example - file sizes
const (
	_  = iota             // ignore first value (0)
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30
	TB                    // 1 << 40
)

func main() {
	fmt.Println("=== Constants ===")
	fmt.Println("Pi:", Pi)
	fmt.Println("App Name:", AppName)
	fmt.Println("Max Users:", MaxUsers)

	fmt.Println("\n=== Days (iota) ===")
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Friday:", Friday)

	fmt.Println("\n=== File Sizes (iota with bit shift) ===")
	fmt.Println("1 KB =", KB, "bytes")
	fmt.Println("1 MB =", MB, "bytes")
	fmt.Println("1 GB =", GB, "bytes")

	// Using constants in calculations
	fileSize := 5 * MB
	fmt.Println("\n5 MB file =", fileSize, "bytes")
}

// TO RUN: go run 03_constants.go
//
// EXERCISE: Create constants for:
// 1. HTTP status codes (200, 404, 500)
// 2. Log levels using iota (DEBUG, INFO, WARN, ERROR)
