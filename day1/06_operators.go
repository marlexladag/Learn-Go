// Day 1, Exercise 6: Operators
//
// Key concepts:
// - Arithmetic operators: + - * / %
// - Comparison operators: == != < > <= >=
// - Logical operators: && || !
// - Assignment operators: = += -= *= /=

package main

import "fmt"

func main() {
	// Arithmetic Operators
	fmt.Println("=== Arithmetic Operators ===")
	a, b := 10, 3

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)  // integer division
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // modulo (remainder)

	// Float division
	x, y := 10.0, 3.0
	fmt.Printf("%.1f / %.1f = %.4f\n", x, y, x/y)

	// Comparison Operators
	fmt.Println("\n=== Comparison Operators ===")
	fmt.Printf("%d == %d : %t\n", a, b, a == b)
	fmt.Printf("%d != %d : %t\n", a, b, a != b)
	fmt.Printf("%d < %d  : %t\n", a, b, a < b)
	fmt.Printf("%d > %d  : %t\n", a, b, a > b)
	fmt.Printf("%d <= %d : %t\n", a, b, a <= b)
	fmt.Printf("%d >= %d : %t\n", a, b, a >= b)

	// Logical Operators
	fmt.Println("\n=== Logical Operators ===")
	isAdult := true
	hasLicense := false

	fmt.Printf("isAdult && hasLicense: %t\n", isAdult && hasLicense) // AND
	fmt.Printf("isAdult || hasLicense: %t\n", isAdult || hasLicense) // OR
	fmt.Printf("!isAdult: %t\n", !isAdult)                           // NOT

	// Assignment Operators
	fmt.Println("\n=== Assignment Operators ===")
	num := 10
	fmt.Println("num starts at:", num)

	num += 5 // num = num + 5
	fmt.Println("num += 5:", num)

	num -= 3 // num = num - 3
	fmt.Println("num -= 3:", num)

	num *= 2 // num = num * 2
	fmt.Println("num *= 2:", num)

	num /= 4 // num = num / 4
	fmt.Println("num /= 4:", num)

	// Increment and Decrement
	fmt.Println("\n=== Increment/Decrement ===")
	counter := 0
	fmt.Println("counter:", counter)

	counter++ // increment by 1
	fmt.Println("counter++:", counter)

	counter++
	fmt.Println("counter++:", counter)

	counter-- // decrement by 1
	fmt.Println("counter--:", counter)

	// Note: Go does NOT have ++counter or --counter (prefix)
	// Note: counter++ is a statement, not an expression
	//       You cannot do: x := counter++ (this is an error)
}

// TO RUN: go run 06_operators.go
//
// EXERCISE: Write a program that:
// 1. Takes two numbers
// 2. Performs all arithmetic operations
// 3. Checks if the first is greater than the second
// 4. Checks if both are positive (> 0)
