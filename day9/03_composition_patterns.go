// Day 9, Exercise 3: Composition Patterns
//
// Key concepts:
// - Go favors composition over inheritance
// - Embedding provides "has-a" relationships with promoted fields/methods
// - Explicit composition gives more control
// - Combine small, focused types into larger ones

package main

import (
	"fmt"
	"time"
)

// ========================================
// Pattern 1: Explicit Composition
// ========================================

// Address is a reusable component
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s %s, %s",
		a.Street, a.City, a.State, a.ZipCode, a.Country)
}

// Person uses explicit composition - has an Address
type Person struct {
	Name    string
	Email   string
	Address Address // explicit field
}

// Company also uses Address
type Company struct {
	Name    string
	Address Address
}

// ========================================
// Pattern 2: Embedding (Promoted Fields)
// ========================================

// Timestamps provides created/updated tracking
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Timestamps) Touch() {
	t.UpdatedAt = time.Now()
}

func (t *Timestamps) SetCreated() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = t.CreatedAt
}

// Article embeds Timestamps - fields and methods are promoted
type Article struct {
	Timestamps        // embedded - fields/methods promoted
	Title      string
	Content    string
	Author     string
}

// Comment also embeds Timestamps
type Comment struct {
	Timestamps
	Text   string
	Author string
}

// ========================================
// Pattern 3: Multiple Embedding
// ========================================

// Identifiable provides ID functionality
type Identifiable struct {
	ID string
}

// Auditable tracks who modified what
type Auditable struct {
	CreatedBy  string
	ModifiedBy string
}

// Document combines multiple embedded types
type Document struct {
	Identifiable // has ID
	Timestamps   // has CreatedAt, UpdatedAt, Touch()
	Auditable    // has CreatedBy, ModifiedBy
	Title        string
	Content      string
}

// ========================================
// Pattern 4: Embedding with Method Override
// ========================================

// Logger provides basic logging
type Logger struct {
	Prefix string
}

func (l *Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

// Service embeds Logger but can override behavior
type Service struct {
	*Logger // embed pointer for shared state
	Name    string
}

// Log overrides the embedded Logger's Log method
func (s *Service) Log(msg string) {
	// Add service-specific behavior
	fmt.Printf("[%s:%s] %s\n", s.Logger.Prefix, s.Name, msg)
}

// ========================================
// Pattern 5: Composition for Behavior Reuse
// ========================================

// EmailSender handles email functionality
type EmailSender struct {
	SMTPHost string
	Port     int
}

func (e *EmailSender) SendEmail(to, subject, body string) {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
}

// NotificationSystem composes multiple capabilities
type NotificationSystem struct {
	emailer *EmailSender // explicit composition (private)
}

func NewNotificationSystem(smtpHost string) *NotificationSystem {
	return &NotificationSystem{
		emailer: &EmailSender{SMTPHost: smtpHost, Port: 587},
	}
}

func (n *NotificationSystem) NotifyUser(email, message string) {
	n.emailer.SendEmail(email, "Notification", message)
}

func main() {
	fmt.Println("=== Explicit Composition ===")

	person := Person{
		Name:  "Alice",
		Email: "alice@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			State:   "MA",
			ZipCode: "02101",
			Country: "USA",
		},
	}

	fmt.Printf("%s lives at: %s\n", person.Name, person.Address.FullAddress())

	// Company reuses the same Address type
	company := Company{
		Name: "Acme Corp",
		Address: Address{
			Street:  "456 Corporate Blvd",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
			Country: "USA",
		},
	}
	fmt.Printf("%s is located at: %s\n", company.Name, company.Address.FullAddress())

	fmt.Println("\n=== Embedding (Promoted Fields) ===")

	article := Article{
		Title:   "Go Composition Patterns",
		Content: "Go favors composition over inheritance...",
		Author:  "Alice",
	}
	article.SetCreated() // Method promoted from Timestamps

	fmt.Printf("Article: %s by %s\n", article.Title, article.Author)
	fmt.Printf("Created: %v\n", article.CreatedAt) // Field promoted from Timestamps

	time.Sleep(10 * time.Millisecond)
	article.Touch() // Update the timestamp
	fmt.Printf("Updated: %v\n", article.UpdatedAt)

	fmt.Println("\n=== Multiple Embedding ===")

	doc := Document{
		Identifiable: Identifiable{ID: "DOC-001"},
		Auditable:    Auditable{CreatedBy: "alice", ModifiedBy: "alice"},
		Title:        "Important Document",
		Content:      "This is the content...",
	}
	doc.SetCreated()

	// All fields are promoted to top level
	fmt.Printf("ID: %s\n", doc.ID)
	fmt.Printf("Title: %s\n", doc.Title)
	fmt.Printf("Created by: %s at %v\n", doc.CreatedBy, doc.CreatedAt)

	fmt.Println("\n=== Embedding with Override ===")

	logger := &Logger{Prefix: "APP"}
	service := &Service{
		Logger: logger,
		Name:   "UserService",
	}

	// Logger's method
	logger.Log("Starting application")

	// Service's overridden method
	service.Log("Processing request")

	// Can still access embedded logger directly
	service.Logger.Log("Direct logger access")

	fmt.Println("\n=== Private Composition ===")

	notifier := NewNotificationSystem("smtp.example.com")
	notifier.NotifyUser("user@example.com", "Welcome to our platform!")

	fmt.Println("\n=== Composition vs Inheritance ===")
	fmt.Println("Go has NO inheritance - only composition!")
	fmt.Println("")
	fmt.Println("Inheritance (other languages):")
	fmt.Println("  class Dog extends Animal { }")
	fmt.Println("")
	fmt.Println("Composition (Go):")
	fmt.Println("  type Dog struct { Animal } // embedding")
	fmt.Println("  type Dog struct { animal Animal } // explicit")
	fmt.Println("")
	fmt.Println("Benefits of composition:")
	fmt.Println("  - More flexible")
	fmt.Println("  - Avoids fragile base class problem")
	fmt.Println("  - Easier to understand relationships")
	fmt.Println("  - Can compose multiple types")
}

// TO RUN: go run day9/03_composition_patterns.go
//
// OUTPUT:
// === Explicit Composition ===
// Alice lives at: 123 Main St, Boston, MA 02101, USA
// ...
//
// EXERCISE:
// 1. Create a Vehicle struct with Make, Model, Year
// 2. Create a Engine struct with Horsepower, Type (gas/electric)
// 3. Create a Car that composes Vehicle and Engine
// 4. Create a Motorcycle that composes Vehicle and Engine differently
// 5. Add methods to each that use the composed parts
//
// KEY POINTS:
// - Explicit composition: type A struct { b B } - access via a.b
// - Embedding: type A struct { B } - fields/methods promoted to a.Field
// - Embedding is NOT inheritance - it's automatic delegation
// - Multiple types can be embedded
// - Embedded method can be overridden by defining same method on outer type
