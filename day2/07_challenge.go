// Day 2 Challenge: Prime Number Checker
//
// Build a program that:
// 1. Asks for a number
// 2. Determines if it's prime
// 3. Lists all primes up to that number
//
// A prime number is only divisible by 1 and itself

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("================================")
	fmt.Println("    Prime Number Checker")
	fmt.Println("================================")

	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number < 2 {
		fmt.Println("Please enter a number >= 2")
		return
	}

	// Check if the input number is prime
	fmt.Printf("\n=== Is %d Prime? ===\n", number)
	if isPrime(number) {
		fmt.Printf("Yes! %d is a prime number.\n", number)
	} else {
		fmt.Printf("No, %d is not a prime number.\n", number)
		// Show factors
		fmt.Print("Factors: ")
		for i := 1; i <= number; i++ {
			if number%i == 0 {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}

	// List all primes up to the number
	fmt.Printf("\n=== All Primes from 2 to %d ===\n", number)
	primeCount := 0

	for i := 2; i <= number; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			primeCount++
		}
	}
	fmt.Printf("\n\nTotal: %d prime numbers found\n", primeCount)

	// Bonus: Show the next prime after the input
	fmt.Println("\n=== Next Prime ===")
	next := number + 1
	for !isPrime(next) {
		next++
	}
	fmt.Printf("The next prime after %d is %d\n", number, next)
}

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Only check up to square root
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// TO RUN: go run 07_challenge.go
//
// SAMPLE OUTPUT:
// ================================
//     Prime Number Checker
// ================================
// Enter a number: 30
//
// === Is 30 Prime? ===
// No, 30 is not a prime number.
// Factors: 1 2 3 5 6 10 15 30
//
// === All Primes from 2 to 30 ===
// 2 3 5 7 11 13 17 19 23 29
//
// Total: 10 prime numbers found
//
// === Next Prime ===
// The next prime after 30 is 31
//
// NOTE: We used a function here (isPrime). We'll learn
// about functions in detail on Day 3!
//
// BONUS EXERCISE:
// Implement the Sieve of Eratosthenes algorithm
// for finding primes (more efficient for large ranges)
