// Day 2, Exercise 4: FizzBuzz
//
// The classic programming challenge!
// Print numbers 1-100 with these rules:
// - If divisible by 3, print "Fizz"
// - If divisible by 5, print "Buzz"
// - If divisible by both, print "FizzBuzz"
// - Otherwise, print the number

package main

import "fmt"

func main() {
	fmt.Println("=== FizzBuzz ===")

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// TO RUN: go run 04_fizzbuzz.go
//
// ALTERNATIVE: Try solving this without switch,
// using only if/else statements
//
// BONUS: Modify to handle custom ranges
// (e.g., FizzBuzz from 50 to 75)
