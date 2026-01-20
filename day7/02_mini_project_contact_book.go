// Day 7, Exercise 2: Mini Project - Contact Book
//
// This project combines everything from Days 1-6:
// - Variables and types
// - Control flow (loops, conditionals)
// - Functions with multiple returns
// - Slices for dynamic storage
// - Maps for fast lookup
// - Pointers for modification

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Contact represents a person in the contact book
// (We're using a map here; in Day 8 we'll learn about structs!)
type Contact = map[string]string

// ContactBook holds all contacts, indexed by name
var contacts = make(map[string]Contact)

func main() {
	fmt.Println("=== Contact Book ===")
	fmt.Println("Commands: add, find, list, update, delete, quit")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			handleAdd(reader)
		case "find":
			if len(parts) > 1 {
				handleFind(parts[1])
			} else {
				fmt.Println("Usage: find <name>")
			}
		case "list":
			handleList()
		case "update":
			if len(parts) > 1 {
				handleUpdate(reader, parts[1])
			} else {
				fmt.Println("Usage: update <name>")
			}
		case "delete":
			if len(parts) > 1 {
				handleDelete(parts[1])
			} else {
				fmt.Println("Usage: delete <name>")
			}
		case "quit", "exit", "q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Try: add, find, list, update, delete, quit")
		}
	}
}

// handleAdd creates a new contact (uses functions, maps, input)
func handleAdd(reader *bufio.Reader) {
	name := prompt(reader, "Name: ")
	if name == "" {
		fmt.Println("Name cannot be empty")
		return
	}

	if _, exists := contacts[name]; exists {
		fmt.Println("Contact already exists. Use 'update' to modify.")
		return
	}

	phone := prompt(reader, "Phone: ")
	email := prompt(reader, "Email: ")

	// Create contact using a map
	contacts[name] = Contact{
		"phone": phone,
		"email": email,
	}

	fmt.Printf("Added contact: %s\n", name)
}

// handleFind searches for a contact (uses comma-ok idiom)
func handleFind(name string) {
	contact, found := contacts[name]
	if !found {
		// Try case-insensitive search
		for n, c := range contacts {
			if strings.EqualFold(n, name) {
				printContact(n, c)
				return
			}
		}
		fmt.Println("Contact not found")
		return
	}
	printContact(name, contact)
}

// handleList shows all contacts (uses range, slices for sorting)
func handleList() {
	if len(contacts) == 0 {
		fmt.Println("No contacts yet. Use 'add' to create one.")
		return
	}

	fmt.Printf("\n--- All Contacts (%d) ---\n", len(contacts))

	// Collect names into a slice
	names := make([]string, 0, len(contacts))
	for name := range contacts {
		names = append(names, name)
	}

	// Print each contact
	for _, name := range names {
		printContact(name, contacts[name])
	}
}

// handleUpdate modifies an existing contact (uses pointers concept)
func handleUpdate(reader *bufio.Reader, name string) {
	contact, found := contacts[name]
	if !found {
		fmt.Println("Contact not found")
		return
	}

	fmt.Println("Leave blank to keep current value")

	phone := prompt(reader, fmt.Sprintf("Phone [%s]: ", contact["phone"]))
	if phone != "" {
		contact["phone"] = phone
	}

	email := prompt(reader, fmt.Sprintf("Email [%s]: ", contact["email"]))
	if email != "" {
		contact["email"] = email
	}

	contacts[name] = contact
	fmt.Println("Contact updated")
}

// handleDelete removes a contact
func handleDelete(name string) {
	if _, found := contacts[name]; !found {
		fmt.Println("Contact not found")
		return
	}

	delete(contacts, name)
	fmt.Printf("Deleted: %s\n", name)
}

// printContact displays a single contact
func printContact(name string, contact Contact) {
	fmt.Printf("\n  %s\n", name)
	fmt.Printf("    Phone: %s\n", contact["phone"])
	fmt.Printf("    Email: %s\n", contact["email"])
}

// prompt reads input with a custom prompt (helper function)
func prompt(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// TO RUN: go run day7/02_mini_project_contact_book.go
//
// EXAMPLE SESSION:
// > add
// Name: Alice
// Phone: 555-1234
// Email: alice@example.com
// Added contact: Alice
//
// > find alice
//   Alice
//     Phone: 555-1234
//     Email: alice@example.com
//
// > list
// --- All Contacts (1) ---
//   Alice
//     Phone: 555-1234
//     Email: alice@example.com
//
// CONCEPTS USED:
// - Variables & constants
// - Control flow (for loop, switch, if/else)
// - Functions with parameters and returns
// - Slices (collecting names)
// - Maps (storing contacts and contact data)
// - Pointers (bufio.Reader is passed by pointer)
//
// EXTENSIONS TO TRY:
// 1. Add a "search" command for partial name matching
// 2. Add more fields (address, birthday, notes)
// 3. Save/load contacts to a file (Day 8+ material)
