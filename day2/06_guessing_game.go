// Day 2, Exercise 6: Guessing Game
//
// A complete mini-project combining:
// - Variables
// - User input
// - Loops
// - Conditionals
// - Random numbers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number between 1 and 100
	secretNumber := rand.Intn(100) + 1
	maxAttempts := 7
	attempts := 0

	fmt.Println("=================================")
	fmt.Println("   Welcome to the Guessing Game!")
	fmt.Println("=================================")
	fmt.Printf("I'm thinking of a number between 1 and 100.\n")
	fmt.Printf("You have %d attempts. Good luck!\n\n", maxAttempts)

	for attempts < maxAttempts {
		attempts++
		attemptsLeft := maxAttempts - attempts

		var guess int
		fmt.Printf("Attempt %d/%d - Enter your guess: ", attempts, maxAttempts)
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Please enter a valid number!")
			attempts-- // Don't count invalid input
			continue
		}

		// Validate range
		if guess < 1 || guess > 100 {
			fmt.Println("Please guess a number between 1 and 100!")
			attempts-- // Don't count invalid input
			continue
		}

		// Check the guess
		switch {
		case guess == secretNumber:
			fmt.Println("\n🎉 CONGRATULATIONS! 🎉")
			fmt.Printf("You guessed it in %d attempts!\n", attempts)

			// Rate the performance
			switch {
			case attempts <= 3:
				fmt.Println("Amazing! You're a mind reader!")
			case attempts <= 5:
				fmt.Println("Great job! Very impressive!")
			default:
				fmt.Println("Good work! You got it!")
			}
			return // Exit the program

		case guess < secretNumber:
			fmt.Printf("Too LOW! ⬆️")
		case guess > secretNumber:
			fmt.Printf("Too HIGH! ⬇️")
		}

		if attemptsLeft > 0 {
			fmt.Printf(" (%d attempts left)\n\n", attemptsLeft)
		}
	}

	// Out of attempts
	fmt.Println("\n😢 Game Over!")
	fmt.Printf("The number was: %d\n", secretNumber)
	fmt.Println("Better luck next time!")
}

// TO RUN: go run 06_guessing_game.go
//
// EXERCISE: Enhance the game:
// 1. Add difficulty levels (easy: 10 attempts, hard: 5)
// 2. Track high scores
// 3. Add a "play again?" feature
// 4. Give hints like "You're very close!" when within 5
