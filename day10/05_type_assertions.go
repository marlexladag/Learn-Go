package main

import (
	"fmt"
	"strings"
)

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 5: Type Assertions and Type Switches
// ============================================================================
//
// TYPE ASSERTIONS
//
// When you have an interface value, you often need to access the concrete
// type underneath. Type assertions let you:
//
//   1. Extract the concrete value from an interface
//   2. Check if an interface holds a specific type
//   3. Convert between interface types
//
// Syntax:
//   value := x.(Type)        // panics if x doesn't hold Type
//   value, ok := x.(Type)    // safe version - ok is false if wrong type
//
// ============================================================================

// Messenger interface
type Messenger interface {
	SendMessage(to, content string) error
}

// EmailSender implements Messenger
type EmailSender struct {
	SMTPServer string
	From       string
}

func (e EmailSender) SendMessage(to, content string) error {
	fmt.Printf("Email from %s to %s via %s: %s\n", e.From, to, e.SMTPServer, content)
	return nil
}

// Extra method not in interface
func (e EmailSender) SendWithAttachment(to, content, attachment string) error {
	fmt.Printf("Email with attachment %s to %s: %s\n", attachment, to, content)
	return nil
}

// SMSSender implements Messenger
type SMSSender struct {
	PhoneNumber string
}

func (s SMSSender) SendMessage(to, content string) error {
	// SMS messages are limited to 160 chars
	if len(content) > 160 {
		content = content[:157] + "..."
	}
	fmt.Printf("SMS from %s to %s: %s\n", s.PhoneNumber, to, content)
	return nil
}

// Extra method
func (s SMSSender) GetRemainingCredits() int {
	return 42 // simulated
}

// ============================================================================
// BASIC TYPE ASSERTION
// ============================================================================

func demonstrateTypeAssertion(m Messenger) {
	// Unsafe assertion - panics if wrong type
	// email := m.(EmailSender) // Don't do this without checking!

	// Safe assertion with comma-ok idiom
	if email, ok := m.(EmailSender); ok {
		fmt.Println("This is an EmailSender!")
		fmt.Println("  SMTP Server:", email.SMTPServer)
		// Can call methods not in the interface
		email.SendWithAttachment("user@example.com", "Hello!", "doc.pdf")
	} else {
		fmt.Println("This is NOT an EmailSender")
	}

	if sms, ok := m.(SMSSender); ok {
		fmt.Println("This is an SMSSender!")
		fmt.Println("  Remaining credits:", sms.GetRemainingCredits())
	}
}

// ============================================================================
// TYPE SWITCHES
// ============================================================================
//
// Type switch is a cleaner way to handle multiple possible types:
//
//   switch v := x.(type) {
//   case Type1:
//       // v is Type1
//   case Type2:
//       // v is Type2
//   default:
//       // v is same as x
//   }
//
// ============================================================================

func describeMessenger(m Messenger) {
	switch v := m.(type) {
	case EmailSender:
		fmt.Printf("Email messenger using %s\n", v.SMTPServer)
	case SMSSender:
		fmt.Printf("SMS messenger from %s\n", v.PhoneNumber)
	case *EmailSender:
		fmt.Printf("Pointer to Email messenger using %s\n", v.SMTPServer)
	case *SMSSender:
		fmt.Printf("Pointer to SMS messenger from %s\n", v.PhoneNumber)
	default:
		fmt.Printf("Unknown messenger type: %T\n", v)
	}
}

// ============================================================================
// CHECKING FOR ADDITIONAL INTERFACES
// ============================================================================

// Formatter is an optional interface
type Formatter interface {
	Format() string
}

// PrettyPrinter is another optional interface
type PrettyPrinter interface {
	PrettyPrint() string
}

// Document implements multiple interfaces
type Document struct {
	Title   string
	Content string
}

func (d Document) String() string {
	return d.Title
}

func (d Document) Format() string {
	return fmt.Sprintf("[%s]\n%s", strings.ToUpper(d.Title), d.Content)
}

func (d Document) PrettyPrint() string {
	line := strings.Repeat("=", len(d.Title)+4)
	return fmt.Sprintf("%s\n| %s |\n%s\n%s", line, d.Title, line, d.Content)
}

// SimpleNote only implements String
type SimpleNote struct {
	Text string
}

func (n SimpleNote) String() string {
	return n.Text
}

// Print checks for optional interfaces and uses them if available
func Print(item fmt.Stringer) {
	// Check if it implements PrettyPrinter (most specific)
	if pp, ok := item.(PrettyPrinter); ok {
		fmt.Println("Using PrettyPrint:")
		fmt.Println(pp.PrettyPrint())
		return
	}

	// Check if it implements Formatter
	if f, ok := item.(Formatter); ok {
		fmt.Println("Using Format:")
		fmt.Println(f.Format())
		return
	}

	// Fall back to basic String()
	fmt.Println("Using String:")
	fmt.Println(item.String())
}

// ============================================================================
// TYPE ASSERTION WITH INTERFACE TO INTERFACE
// ============================================================================

// Saver interface
type Saver interface {
	Save() error
}

// Loader interface
type Loader interface {
	Load() error
}

// SaveLoader combines both
type SaveLoader interface {
	Saver
	Loader
}

// DataStore implements SaveLoader
type DataStore struct {
	Name string
	Data string
}

func (ds *DataStore) Save() error {
	fmt.Printf("Saving %s: %s\n", ds.Name, ds.Data)
	return nil
}

func (ds *DataStore) Load() error {
	fmt.Printf("Loading %s\n", ds.Name)
	ds.Data = "loaded data"
	return nil
}

// Works with Saver, but checks if it's also a Loader
func ProcessData(s Saver) {
	fmt.Println("Processing data...")

	// Check if this Saver also implements Loader
	if sl, ok := s.(SaveLoader); ok {
		fmt.Println("  Full SaveLoader - can load and save!")
		sl.Load()
		sl.Save()
	} else if l, ok := s.(Loader); ok {
		fmt.Println("  Also implements Loader separately")
		l.Load()
		s.Save()
	} else {
		fmt.Println("  Only Saver - can only save")
		s.Save()
	}
}

// ============================================================================
// PRACTICAL EXAMPLE: ERROR TYPE CHECKING
// ============================================================================

// Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

type NotFoundError struct {
	Resource string
	ID       string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

type PermissionError struct {
	Action   string
	Resource string
}

func (e PermissionError) Error() string {
	return fmt.Sprintf("permission denied: cannot %s on %s", e.Action, e.Resource)
}

// HandleError uses type switch to handle different error types
func HandleError(err error) {
	if err == nil {
		fmt.Println("No error")
		return
	}

	switch e := err.(type) {
	case ValidationError:
		fmt.Printf("Fix the %s field: %s\n", e.Field, e.Message)
	case NotFoundError:
		fmt.Printf("The %s you're looking for (ID: %s) doesn't exist\n", e.Resource, e.ID)
	case PermissionError:
		fmt.Printf("You don't have permission to %s this %s\n", e.Action, e.Resource)
	default:
		fmt.Printf("Unexpected error: %v\n", err)
	}
}

func main() {
	fmt.Println("=== Type Assertions and Type Switches ===")
	fmt.Println()

	// Basic type assertion
	fmt.Println("--- Basic Type Assertion ---")
	email := EmailSender{SMTPServer: "smtp.gmail.com", From: "me@gmail.com"}
	sms := SMSSender{PhoneNumber: "+1234567890"}

	var m Messenger = email
	demonstrateTypeAssertion(m)

	fmt.Println()
	m = sms
	demonstrateTypeAssertion(m)

	fmt.Println()

	// Type switch
	fmt.Println("--- Type Switch ---")
	messengers := []Messenger{
		EmailSender{SMTPServer: "smtp.example.com", From: "admin@example.com"},
		SMSSender{PhoneNumber: "+0987654321"},
		&EmailSender{SMTPServer: "mail.company.com", From: "info@company.com"},
	}

	for _, msg := range messengers {
		describeMessenger(msg)
	}

	fmt.Println()

	// Checking for optional interfaces
	fmt.Println("--- Optional Interface Checking ---")
	doc := Document{Title: "Important", Content: "This is the content."}
	note := SimpleNote{Text: "Just a quick note"}

	Print(doc)
	fmt.Println()
	Print(note)

	fmt.Println()

	// Interface to interface assertion
	fmt.Println("--- Interface to Interface ---")
	store := &DataStore{Name: "users.db", Data: "user data"}
	ProcessData(store)

	fmt.Println()

	// Error type handling
	fmt.Println("--- Error Type Handling ---")
	errors := []error{
		nil,
		ValidationError{Field: "email", Message: "invalid format"},
		NotFoundError{Resource: "User", ID: "12345"},
		PermissionError{Action: "delete", Resource: "admin account"},
		fmt.Errorf("generic error"),
	}

	for _, err := range errors {
		HandleError(err)
	}
}

// ============================================================================
// TO RUN:
//   go run day10/05_type_assertions.go
//
// EXPECTED OUTPUT:
//   === Type Assertions and Type Switches ===
//
//   --- Basic Type Assertion ---
//   This is an EmailSender!
//     SMTP Server: smtp.gmail.com
//   Email with attachment doc.pdf to user@example.com: Hello!
//
//   This is NOT an EmailSender
//   This is an SMSSender!
//     Remaining credits: 42
//
//   --- Type Switch ---
//   Email messenger using smtp.example.com
//   SMS messenger from +0987654321
//   Pointer to Email messenger using mail.company.com
//
//   ... (more output)
//
// EXERCISE:
//   1. Create a payment processor with CreditCard, PayPal, and Crypto types
//   2. Each should implement a Payer interface with Pay(amount float64) error
//   3. CreditCard should have an extra method ValidateCard() bool
//   4. Write a ProcessPayment function that checks if the payer is CreditCard
//      and validates the card before processing
//
// KEY POINTS:
//   - value := x.(Type) extracts concrete type (panics if wrong)
//   - value, ok := x.(Type) is the safe version (ok is false if wrong)
//   - Type switch handles multiple possible types cleanly
//   - You can assert from one interface to another
//   - Use type assertions to access methods not in the interface
//   - Common pattern: check for optional/extended interfaces
//   - Error handling often uses type switches for different error types
// ============================================================================
