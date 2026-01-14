// Day 1 Challenge: Build a Simple Calculator
//
// This exercise combines everything you learned today:
// - Variables and types
// - Constants
// - Formatted printing
// - User input
// - Operators
//
// YOUR TASK: Complete the TODOs below!

package main

import "fmt"

// Define constants for operation symbols
const (
	ADD      = "+"
	SUBTRACT = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
)

func main() {
	fmt.Println("================================")
	fmt.Println("    Simple Go Calculator")
	fmt.Println("================================")

	// TODO 1: Declare variables for two numbers (float64)
	var num1, num2 float64

	// TODO 2: Get first number from user
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	// TODO 3: Get second number from user
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Display all results
	fmt.Println("\n=== Results ===")

	// TODO 4: Calculate and print addition
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, ADD, num2, num1+num2)

	// TODO 5: Calculate and print subtraction
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, SUBTRACT, num2, num1-num2)

	// TODO 6: Calculate and print multiplication
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, MULTIPLY, num2, num1*num2)

	// TODO 7: Calculate and print division (handle divide by zero!)
	if num2 != 0 {
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, DIVIDE, num2, num1/num2)
	} else {
		fmt.Println("Cannot divide by zero!")
	}

	// BONUS: Add more features!
	// - Modulo operation (for integers)
	// - Power calculation
	// - Compare which number is larger

	fmt.Println("\n=== Comparison ===")
	if num1 > num2 {
		fmt.Printf("%.2f is greater than %.2f\n", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%.2f is less than %.2f\n", num1, num2)
	} else {
		fmt.Println("Both numbers are equal!")
	}
}

// TO RUN: go run 07_challenge.go
//
// EXPECTED OUTPUT:
// ================================
//     Simple Go Calculator
// ================================
// Enter first number: 10
// Enter second number: 3
//
// === Results ===
// 10.00 + 3.00 = 13.00
// 10.00 - 3.00 = 7.00
// 10.00 * 3.00 = 30.00
// 10.00 / 3.00 = 3.33
//
// === Comparison ===
// 10.00 is greater than 3.00
