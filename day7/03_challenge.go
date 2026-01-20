// Day 7, Challenge: Build a Simple Calculator
//
// Combine your knowledge to build an interactive calculator!
//
// Requirements:
// 1. Support +, -, *, / operations
// 2. Handle division by zero gracefully
// 3. Keep a history of calculations
// 4. Allow viewing history
// 5. Support "clear" to reset history

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// History stores past calculations
var history []string

func main() {
	fmt.Println("=== Go Calculator ===")
	fmt.Println("Enter: number operator number (e.g., 5 + 3)")
	fmt.Println("Commands: history, clear, quit")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("calc> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		// Handle commands
		switch strings.ToLower(input) {
		case "history":
			showHistory()
			continue
		case "clear":
			history = nil
			fmt.Println("History cleared")
			continue
		case "quit", "exit", "q":
			fmt.Println("Goodbye!")
			return
		}

		// Parse calculation
		result, expression, err := calculate(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Store in history and display
		entry := fmt.Sprintf("%s = %.2f", expression, result)
		history = append(history, entry)
		fmt.Printf("= %.2f\n", result)
	}
}

// calculate parses input and performs the operation
// Returns: result, formatted expression, error
func calculate(input string) (float64, string, error) {
	parts := strings.Fields(input)

	if len(parts) != 3 {
		return 0, "", fmt.Errorf("expected: number operator number")
	}

	// Parse first number
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid first number: %s", parts[0])
	}

	// Parse operator
	operator := parts[1]

	// Parse second number
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, "", fmt.Errorf("invalid second number: %s", parts[2])
	}

	// Perform calculation
	var result float64
	expression := fmt.Sprintf("%.2f %s %.2f", num1, operator, num2)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return 0, "", fmt.Errorf("division by zero")
		}
		result = num1 / num2
	default:
		return 0, "", fmt.Errorf("unknown operator: %s (use +, -, *, /)", operator)
	}

	return result, expression, nil
}

// showHistory displays all past calculations
func showHistory() {
	if len(history) == 0 {
		fmt.Println("No calculations yet")
		return
	}

	fmt.Println("\n--- History ---")
	for i, entry := range history {
		fmt.Printf("%d. %s\n", i+1, entry)
	}
	fmt.Println()
}

// TO RUN: go run day7/03_challenge.go
//
// EXAMPLE SESSION:
// calc> 10 + 5
// = 15.00
// calc> 100 / 4
// = 25.00
// calc> 3.14 * 2
// = 6.28
// calc> 10 / 0
// Error: division by zero
// calc> history
// --- History ---
// 1. 10.00 + 5.00 = 15.00
// 2. 100.00 / 4.00 = 25.00
// 3. 3.14 * 2.00 = 6.28
//
// BONUS CHALLENGES:
// 1. Add support for parentheses: (5 + 3) * 2
// 2. Add memory functions: M+, M-, MR, MC
// 3. Add scientific functions: sqrt, pow, sin, cos
// 4. Add a "replay" command to redo last calculation
// 5. Support chained operations: 5 + 3 - 2
