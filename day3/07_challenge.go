// Day 3 Challenge: Calculator with Functions
//
// Build a calculator that demonstrates all function concepts learned today:
// - Basic functions
// - Parameters
// - Return values
// - Multiple returns
// - Named returns
// - Variadic functions
//
// The calculator should:
// 1. Perform basic arithmetic operations
// 2. Handle multiple numbers (variadic)
// 3. Return error status for invalid operations
// 4. Calculate statistics on a set of numbers

package main

import (
	"fmt"
	"math"
)

// =====================================================
// BASIC OPERATIONS (Parameters + Returns)
// =====================================================

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

// Multiple returns: result and success status
func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// =====================================================
// ADVANCED OPERATIONS (Multiple Returns)
// =====================================================

// Returns quotient, remainder, and error message
func intDivide(a, b int) (quotient, remainder int, err string) {
	if b == 0 {
		err = "cannot divide by zero"
		return
	}
	quotient = a / b
	remainder = a % b
	return
}

// Returns both roots of quadratic equation ax² + bx + c = 0
func quadraticRoots(a, b, c float64) (root1, root2 float64, hasRealRoots bool) {
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		hasRealRoots = false
		return
	}

	hasRealRoots = true
	sqrtDisc := math.Sqrt(discriminant)
	root1 = (-b + sqrtDisc) / (2 * a)
	root2 = (-b - sqrtDisc) / (2 * a)
	return
}

// =====================================================
// VARIADIC OPERATIONS
// =====================================================

// Sum any number of values
func sumAll(numbers ...float64) float64 {
	total := 0.0
	for _, n := range numbers {
		total += n
	}
	return total
}

// Multiply any number of values
func multiplyAll(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	result := 1.0
	for _, n := range numbers {
		result *= n
	}
	return result
}

// Find max of any number of values
func maximum(first float64, rest ...float64) float64 {
	max := first
	for _, n := range rest {
		if n > max {
			max = n
		}
	}
	return max
}

// Find min of any number of values
func minimum(first float64, rest ...float64) float64 {
	min := first
	for _, n := range rest {
		if n < min {
			min = n
		}
	}
	return min
}

// =====================================================
// STATISTICS (Named Returns + Variadic)
// =====================================================

// Calculate comprehensive statistics
func statistics(numbers ...float64) (count int, sum, avg, min, max float64) {
	count = len(numbers)
	if count == 0 {
		return // All zeros
	}

	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		sum += n
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	avg = sum / float64(count)
	return
}

// =====================================================
// UTILITY FUNCTIONS
// =====================================================

func printSeparator() {
	fmt.Println("────────────────────────────────────────")
}

func printSection(title string) {
	fmt.Println()
	printSeparator()
	fmt.Printf("  %s\n", title)
	printSeparator()
}

// =====================================================
// MAIN PROGRAM
// =====================================================

func main() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║     Day 3 Challenge: Calculator      ║")
	fmt.Println("╚══════════════════════════════════════╝")

	// Basic Operations
	printSection("Basic Operations")
	fmt.Printf("5 + 3 = %.2f\n", add(5, 3))
	fmt.Printf("10 - 4 = %.2f\n", subtract(10, 4))
	fmt.Printf("6 × 7 = %.2f\n", multiply(6, 7))

	result, ok := divide(15, 3)
	if ok {
		fmt.Printf("15 ÷ 3 = %.2f\n", result)
	}

	result, ok = divide(10, 0)
	if !ok {
		fmt.Println("10 ÷ 0 = Error: division by zero!")
	}

	// Integer Division
	printSection("Integer Division")
	q, r, err := intDivide(17, 5)
	if err == "" {
		fmt.Printf("17 ÷ 5 = %d remainder %d\n", q, r)
	}

	q, r, err = intDivide(100, 7)
	if err == "" {
		fmt.Printf("100 ÷ 7 = %d remainder %d\n", q, r)
	}

	// Quadratic Equation Solver
	printSection("Quadratic Equation Solver")

	// x² - 5x + 6 = 0 → roots: 2, 3
	r1, r2, hasRoots := quadraticRoots(1, -5, 6)
	fmt.Print("x² - 5x + 6 = 0 → ")
	if hasRoots {
		fmt.Printf("x = %.2f, %.2f\n", r1, r2)
	}

	// x² + 2x + 1 = 0 → root: -1 (double)
	r1, r2, hasRoots = quadraticRoots(1, 2, 1)
	fmt.Print("x² + 2x + 1 = 0 → ")
	if hasRoots {
		fmt.Printf("x = %.2f, %.2f\n", r1, r2)
	}

	// x² + x + 1 = 0 → no real roots
	r1, r2, hasRoots = quadraticRoots(1, 1, 1)
	fmt.Print("x² + x + 1 = 0 → ")
	if !hasRoots {
		fmt.Println("No real roots")
	}

	// Variadic Operations
	printSection("Variadic Operations")
	fmt.Printf("sumAll(1, 2, 3, 4, 5) = %.2f\n", sumAll(1, 2, 3, 4, 5))
	fmt.Printf("sumAll(10, 20, 30) = %.2f\n", sumAll(10, 20, 30))
	fmt.Printf("multiplyAll(2, 3, 4) = %.2f\n", multiplyAll(2, 3, 4))
	fmt.Printf("maximum(3, 7, 2, 9, 5) = %.2f\n", maximum(3, 7, 2, 9, 5))
	fmt.Printf("minimum(3, 7, 2, 9, 5) = %.2f\n", minimum(3, 7, 2, 9, 5))

	// Using slice with variadic
	values := []float64{15, 22, 8, 42, 11, 36}
	fmt.Printf("\nSlice: %v\n", values)
	fmt.Printf("Sum of slice: %.2f\n", sumAll(values...))
	fmt.Printf("Max of slice: %.2f\n", maximum(values[0], values[1:]...))

	// Statistics
	printSection("Statistics Calculator")
	testData := []float64{85, 92, 78, 95, 88, 72, 90, 85, 91, 87}
	fmt.Printf("Data: %v\n\n", testData)

	count, sum, avg, min, max := statistics(testData...)
	fmt.Printf("Count:   %d\n", count)
	fmt.Printf("Sum:     %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Minimum: %.2f\n", min)
	fmt.Printf("Maximum: %.2f\n", max)

	// Interactive Mode
	printSection("Interactive Calculator")
	fmt.Println("Enter two numbers for calculation:")

	var a, b float64
	fmt.Print("First number: ")
	fmt.Scan(&a)
	fmt.Print("Second number: ")
	fmt.Scan(&b)

	fmt.Printf("\nResults for %.2f and %.2f:\n", a, b)
	fmt.Printf("  Add:      %.2f\n", add(a, b))
	fmt.Printf("  Subtract: %.2f\n", subtract(a, b))
	fmt.Printf("  Multiply: %.2f\n", multiply(a, b))

	if result, ok := divide(a, b); ok {
		fmt.Printf("  Divide:   %.2f\n", result)
	} else {
		fmt.Println("  Divide:   Error (division by zero)")
	}

	fmt.Println("\n✓ Day 3 Challenge Complete!")
}

// TO RUN: go run 07_challenge.go
//
// SAMPLE OUTPUT:
// ╔══════════════════════════════════════╗
// ║     Day 3 Challenge: Calculator      ║
// ╚══════════════════════════════════════╝
//
// ────────────────────────────────────────
//   Basic Operations
// ────────────────────────────────────────
// 5 + 3 = 8.00
// 10 - 4 = 6.00
// 6 × 7 = 42.00
// 15 ÷ 3 = 5.00
// 10 ÷ 0 = Error: division by zero!
//
// ... (continues with all sections)
//
// CONCEPTS DEMONSTRATED:
// ✓ Basic functions (printSeparator, printSection)
// ✓ Parameters (add, subtract, multiply)
// ✓ Single return value (add, subtract, multiply)
// ✓ Multiple return values (divide, quadraticRoots)
// ✓ Named return values (intDivide, statistics)
// ✓ Variadic functions (sumAll, multiplyAll, maximum, minimum)
// ✓ Spreading slices to variadic (values...)
//
// BONUS EXERCISES:
// 1. Add a power function: power(base, exponent float64) float64
// 2. Add a function to calculate standard deviation
// 3. Add support for more operations (modulo, absolute value)
// 4. Create a function that returns all prime factors of a number
