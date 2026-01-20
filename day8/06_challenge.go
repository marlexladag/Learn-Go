// Day 8, Challenge: Build a Library System
//
// Create a simple library management system using structs and methods.
//
// Requirements:
// 1. Book struct with title, author, ISBN, available status
// 2. Library struct that holds books and members
// 3. Member struct with name, ID, borrowed books
// 4. Methods: AddBook, BorrowBook, ReturnBook, ListAvailable

package main

import (
	"encoding/json"
	"fmt"
)

// Book represents a book in the library
type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Available bool   `json:"available"`
}

// Member represents a library member
type Member struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Borrowed []string `json:"borrowed"` // ISBNs of borrowed books
}

// Library holds books and members
type Library struct {
	Name    string            `json:"name"`
	Books   map[string]*Book  `json:"books"`   // ISBN -> Book
	Members map[string]*Member `json:"members"` // ID -> Member
}

// NewLibrary creates a new library
func NewLibrary(name string) *Library {
	return &Library{
		Name:    name,
		Books:   make(map[string]*Book),
		Members: make(map[string]*Member),
	}
}

// AddBook adds a book to the library
func (l *Library) AddBook(title, author, isbn string) {
	l.Books[isbn] = &Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}
	fmt.Printf("Added: %q by %s\n", title, author)
}

// AddMember registers a new member
func (l *Library) AddMember(id, name string) {
	l.Members[id] = &Member{
		ID:       id,
		Name:     name,
		Borrowed: []string{},
	}
	fmt.Printf("Registered member: %s (%s)\n", name, id)
}

// BorrowBook allows a member to borrow a book
func (l *Library) BorrowBook(memberID, isbn string) error {
	// Check if member exists
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member %s not found", memberID)
	}

	// Check if book exists
	book, exists := l.Books[isbn]
	if !exists {
		return fmt.Errorf("book %s not found", isbn)
	}

	// Check if book is available
	if !book.Available {
		return fmt.Errorf("book %q is not available", book.Title)
	}

	// Borrow the book
	book.Available = false
	member.Borrowed = append(member.Borrowed, isbn)

	fmt.Printf("%s borrowed %q\n", member.Name, book.Title)
	return nil
}

// ReturnBook processes a book return
func (l *Library) ReturnBook(memberID, isbn string) error {
	// Check if member exists
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member %s not found", memberID)
	}

	// Check if book exists
	book, exists := l.Books[isbn]
	if !exists {
		return fmt.Errorf("book %s not found", isbn)
	}

	// Check if member has this book
	found := false
	for i, borrowed := range member.Borrowed {
		if borrowed == isbn {
			// Remove from borrowed list
			member.Borrowed = append(member.Borrowed[:i], member.Borrowed[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%s does not have %q borrowed", member.Name, book.Title)
	}

	// Return the book
	book.Available = true
	fmt.Printf("%s returned %q\n", member.Name, book.Title)
	return nil
}

// ListAvailable shows all available books
func (l *Library) ListAvailable() {
	fmt.Println("\nAvailable Books:")
	fmt.Println("================")

	count := 0
	for _, book := range l.Books {
		if book.Available {
			fmt.Printf("  %q by %s (ISBN: %s)\n", book.Title, book.Author, book.ISBN)
			count++
		}
	}

	if count == 0 {
		fmt.Println("  No books available")
	}
	fmt.Printf("\nTotal: %d available\n", count)
}

// ListBorrowed shows what a member has borrowed
func (l *Library) ListBorrowed(memberID string) {
	member, exists := l.Members[memberID]
	if !exists {
		fmt.Printf("Member %s not found\n", memberID)
		return
	}

	fmt.Printf("\nBooks borrowed by %s:\n", member.Name)
	fmt.Println("========================")

	if len(member.Borrowed) == 0 {
		fmt.Println("  No books borrowed")
		return
	}

	for _, isbn := range member.Borrowed {
		if book, ok := l.Books[isbn]; ok {
			fmt.Printf("  %q by %s\n", book.Title, book.Author)
		}
	}
}

// Stats returns library statistics
func (l *Library) Stats() {
	totalBooks := len(l.Books)
	available := 0
	for _, book := range l.Books {
		if book.Available {
			available++
		}
	}

	fmt.Printf("\n%s Statistics:\n", l.Name)
	fmt.Println("==================")
	fmt.Printf("  Total Books: %d\n", totalBooks)
	fmt.Printf("  Available: %d\n", available)
	fmt.Printf("  Borrowed: %d\n", totalBooks-available)
	fmt.Printf("  Members: %d\n", len(l.Members))
}

// ToJSON exports library to JSON
func (l *Library) ToJSON() string {
	data, _ := json.MarshalIndent(l, "", "  ")
	return string(data)
}

func main() {
	fmt.Println("=== Library Management System ===\n")

	// Create library
	lib := NewLibrary("City Library")

	// Add some books
	fmt.Println("Adding books...")
	lib.AddBook("The Go Programming Language", "Alan Donovan", "978-0134190440")
	lib.AddBook("Clean Code", "Robert Martin", "978-0132350884")
	lib.AddBook("Design Patterns", "Gang of Four", "978-0201633610")
	lib.AddBook("The Pragmatic Programmer", "David Thomas", "978-0135957059")

	// Register members
	fmt.Println("\nRegistering members...")
	lib.AddMember("M001", "Alice")
	lib.AddMember("M002", "Bob")

	// Show available books
	lib.ListAvailable()

	// Borrow some books
	fmt.Println("\n--- Borrowing ---")
	lib.BorrowBook("M001", "978-0134190440")
	lib.BorrowBook("M001", "978-0132350884")
	lib.BorrowBook("M002", "978-0201633610")

	// Try to borrow unavailable book
	err := lib.BorrowBook("M002", "978-0134190440")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show what each member has
	lib.ListBorrowed("M001")
	lib.ListBorrowed("M002")

	// Return a book
	fmt.Println("\n--- Returning ---")
	lib.ReturnBook("M001", "978-0132350884")

	// Updated stats
	lib.Stats()

	// Show available books again
	lib.ListAvailable()

	// Export to JSON
	fmt.Println("\n--- JSON Export ---")
	fmt.Println(lib.ToJSON())
}

// TO RUN: go run day8/06_challenge.go
//
// OUTPUT:
// === Library Management System ===
//
// Adding books...
// Added: "The Go Programming Language" by Alan Donovan
// ...
//
// CONCEPTS USED:
// - Struct definitions with multiple fields
// - Struct tags for JSON serialization
// - Pointer receivers for methods that modify state
// - Maps for O(1) lookups
// - Slices for dynamic collections
// - Error handling with multiple returns
// - Constructor pattern (NewLibrary)
//
// EXTENSIONS TO TRY:
// 1. Add due dates for borrowed books
// 2. Implement book search by title/author
// 3. Add a maximum borrow limit per member
// 4. Save/load library state to file
// 5. Add book categories/genres
