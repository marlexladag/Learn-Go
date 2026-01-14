// Day 3, Exercise 3: Return Values
//
// Key concepts:
// - Functions can return values using 'return'
// - Return type is specified after parameters
// - Returned values can be stored in variables
// - Return immediately exits the function

package main

import (
	"fmt"
	"math"
)

// Function that returns an int
func addNumbers(a, b int) int {
	return a + b
}

// Function that returns a float64
func multiply(a, b float64) float64 {
	return a * b
}

// Function that returns a string
func getGreeting(name string) string {
	return "Hello, " + name + "!"
}

// Function that returns a bool
func isEven(n int) bool {
	return n%2 == 0
}

// Function that returns a bool
func isPositive(n int) bool {
	return n > 0
}

// Using return for early exit
func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: division by zero!")
		return 0 // Early return
	}
	return a / b
}

// Practical example: calculate area of a circle
func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// Practical example: convert Celsius to Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// Function that returns result of a comparison
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function that computes factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println("=== Basic Return Values ===")
	sum := addNumbers(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)

	product := multiply(3.5, 2.0)
	fmt.Printf("3.5 * 2.0 = %.1f\n", product)

	fmt.Println("\n=== Returning Strings ===")
	message := getGreeting("Gopher")
	fmt.Println(message)

	// Can also use return value directly
	fmt.Println(getGreeting("World"))

	fmt.Println("\n=== Returning Booleans ===")
	fmt.Printf("Is 10 even? %t\n", isEven(10))
	fmt.Printf("Is 7 even? %t\n", isEven(7))
	fmt.Printf("Is 5 positive? %t\n", isPositive(5))
	fmt.Printf("Is -3 positive? %t\n", isPositive(-3))

	fmt.Println("\n=== Using Return Values in Conditions ===")
	number := 15
	if isEven(number) {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}

	fmt.Println("\n=== Early Return Example ===")
	result1 := divide(10, 2)
	fmt.Printf("10 / 2 = %.1f\n", result1)
	result2 := divide(10, 0)
	fmt.Printf("10 / 0 = %.1f\n", result2)

	fmt.Println("\n=== Practical Examples ===")
	radius := 5.0
	area := circleArea(radius)
	fmt.Printf("Circle with radius %.1f has area %.2f\n", radius, area)

	celsius := 100.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("%.1f°C = %.1f°F\n", celsius, fahrenheit)

	fmt.Println("\n=== Using Return in Logic ===")
	fmt.Printf("max(10, 25) = %d\n", max(10, 25))
	fmt.Printf("max(100, 50) = %d\n", max(100, 50))

	fmt.Println("\n=== Factorial ===")
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}

// TO RUN: go run 03_return_values.go
//
// OUTPUT:
// === Basic Return Values ===
// 10 + 20 = 30
// 3.5 * 2.0 = 7.0
//
// === Returning Strings ===
// Hello, Gopher!
// Hello, World!
//
// === Returning Booleans ===
// Is 10 even? true
// Is 7 even? false
// Is 5 positive? true
// Is -3 positive? false
//
// ... (continues)
//
// KEY POINTS:
// - Return type comes after the parameter list: func name(params) returnType
// - Use 'return' keyword followed by the value
// - The returned value must match the declared return type
// - Return values can be used in expressions, conditions, or stored in variables
// - 'return' immediately exits the function
//
// EXERCISE: Write a function isPrime(n int) bool that returns
// true if n is a prime number, false otherwise
