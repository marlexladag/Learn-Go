// Day 3, Exercise 5: Named Return Values
//
// Key concepts:
// - Return values can be named in the function signature
// - Named returns act as variables inside the function
// - "Naked return" returns the named values automatically
// - Useful for documentation and certain patterns

package main

import (
	"fmt"
	"math"
)

// Named return values - they're pre-declared as variables
func rectangleProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // "naked return" - returns area and perimeter
}

// Named returns with explicit return (recommended for clarity)
func circleProps(radius float64) (area, circumference float64) {
	area = math.Pi * radius * radius
	circumference = 2 * math.Pi * radius
	return area, circumference // Explicit is often clearer
}

// Named returns are initialized to zero values
func stats(numbers []int) (sum int, count int, average float64) {
	// sum and count start at 0, average starts at 0.0
	if len(numbers) == 0 {
		return // Returns 0, 0, 0.0
	}

	count = len(numbers)
	for _, n := range numbers {
		sum += n
	}
	average = float64(sum) / float64(count)
	return
}

// Named returns help document what a function returns
func divideWithRemainder(dividend, divisor int) (quotient, remainder int, err string) {
	if divisor == 0 {
		err = "division by zero"
		return // Returns 0, 0, "division by zero"
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// Comparing anonymous vs named returns
// Anonymous:
func getPointAnon() (int, int) {
	return 10, 20
}

// Named (self-documenting):
func getPointNamed() (x, y int) {
	x = 10
	y = 20
	return
}

// Named returns can be reassigned
func safeDivide(a, b float64) (result float64, ok bool) {
	ok = true // Assume success
	if b == 0 {
		ok = false
		return
	}
	result = a / b
	return
}

// Practical example: Parse a simple coordinate string
func parseCoordinates(s string) (x, y int, valid bool) {
	// Simple parser for "x,y" format
	commaPos := -1
	for i, c := range s {
		if c == ',' {
			commaPos = i
			break
		}
	}

	if commaPos == -1 {
		return // Returns 0, 0, false (zero values)
	}

	// Parse x (before comma)
	for _, c := range s[:commaPos] {
		if c < '0' || c > '9' {
			return
		}
		x = x*10 + int(c-'0')
	}

	// Parse y (after comma)
	for _, c := range s[commaPos+1:] {
		if c < '0' || c > '9' {
			return
		}
		y = y*10 + int(c-'0')
	}

	valid = true
	return
}

func main() {
	fmt.Println("=== Rectangle Properties ===")
	area, perim := rectangleProps(5, 3)
	fmt.Printf("Rectangle 5x3: Area=%.1f, Perimeter=%.1f\n", area, perim)

	fmt.Println("\n=== Circle Properties ===")
	cArea, circ := circleProps(4)
	fmt.Printf("Circle r=4: Area=%.2f, Circumference=%.2f\n", cArea, circ)

	fmt.Println("\n=== Statistics ===")
	numbers := []int{10, 20, 30, 40, 50}
	sum, count, avg := stats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d, Count: %d, Average: %.2f\n", sum, count, avg)

	// Empty slice returns zero values
	sum, count, avg = stats([]int{})
	fmt.Printf("Empty slice: Sum=%d, Count=%d, Average=%.2f\n", sum, count, avg)

	fmt.Println("\n=== Division with Remainder ===")
	q, r, err := divideWithRemainder(17, 5)
	if err == "" {
		fmt.Printf("17 / 5 = %d remainder %d\n", q, r)
	}

	q, r, err = divideWithRemainder(10, 0)
	if err != "" {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("\n=== Safe Divide ===")
	result, ok := safeDivide(10, 3)
	if ok {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	result, ok = safeDivide(10, 0)
	if !ok {
		fmt.Println("Cannot divide by zero")
	}

	fmt.Println("\n=== Coordinate Parsing ===")
	testCoords := []string{"10,20", "5,15", "invalid", "100,200"}
	for _, s := range testCoords {
		x, y, valid := parseCoordinates(s)
		if valid {
			fmt.Printf("'%s' -> x=%d, y=%d\n", s, x, y)
		} else {
			fmt.Printf("'%s' -> invalid format\n", s)
		}
	}
}

// TO RUN: go run 05_named_returns.go
//
// OUTPUT:
// === Rectangle Properties ===
// Rectangle 5x3: Area=15.0, Perimeter=16.0
//
// === Circle Properties ===
// Circle r=4: Area=50.27, Circumference=25.13
//
// ... (continues)
//
// KEY POINTS:
// - Named returns: func f() (name type, name2 type2)
// - Named return variables are initialized to zero values
// - "Naked return" (just 'return') returns the named values
// - Named returns serve as documentation
// - For complex functions, explicit return is often clearer
// - Named returns are useful when return values need explanation
//
// WHEN TO USE NAMED RETURNS:
// - When return values need documentation
// - When you want zero-value returns on error
// - Short functions where naked return is clear
//
// WHEN TO AVOID:
// - Long functions where naked return is confusing
// - When it makes code harder to follow
// - When explicit return values are clearer
//
// EXERCISE: Rewrite the minMax function from the previous exercise
// using named return values
