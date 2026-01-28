package main

import "fmt"

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 2: Implicit Implementation
// ============================================================================
//
// GO'S UNIQUE APPROACH TO INTERFACES
//
// In Go, interface implementation is IMPLICIT - there's no "implements" keyword!
// A type implements an interface simply by having all the required methods.
//
// This is different from Java/C#:
//   Java:   class Dog implements Speaker { ... }
//   Go:     type Dog struct{} + methods = automatically implements!
//
// This is called "structural typing" or "duck typing":
// "If it walks like a duck and quacks like a duck, it's a duck."
//
// ============================================================================

// Stringer is our interface - anything that can represent itself as a string
type Stringer interface {
	String() string
}

// Book implements Stringer without explicitly declaring it
type Book struct {
	Title  string
	Author string
	Year   int
}

// This method makes Book implement Stringer
func (b Book) String() string {
	return fmt.Sprintf("%s by %s (%d)", b.Title, b.Author, b.Year)
}

// Movie also implements Stringer
type Movie struct {
	Title    string
	Director string
	Rating   float64
}

func (m Movie) String() string {
	return fmt.Sprintf("%s directed by %s - Rating: %.1f", m.Title, m.Director, m.Rating)
}

// PrintItem works with anything that can be stringified
func PrintItem(s Stringer) {
	fmt.Println("Item:", s.String())
}

// ============================================================================
// IMPLEMENTING STANDARD LIBRARY INTERFACES
// ============================================================================
//
// The standard library defines many useful interfaces. Your types can
// implement them to integrate seamlessly with Go's ecosystem.
//
// Common standard library interfaces:
//   - fmt.Stringer: String() string
//   - error: Error() string
//   - io.Reader: Read(p []byte) (n int, err error)
//   - io.Writer: Write(p []byte) (n int, err error)
//   - sort.Interface: Len(), Less(), Swap()
//
// ============================================================================

// Temperature implements fmt.Stringer (works with fmt.Println!)
type Temperature struct {
	Celsius float64
}

// This makes Temperature work beautifully with fmt package
func (t Temperature) String() string {
	fahrenheit := t.Celsius*9/5 + 32
	return fmt.Sprintf("%.1f°C (%.1f°F)", t.Celsius, fahrenheit)
}

// ============================================================================
// VERIFICATION: COMPILE-TIME INTERFACE CHECKS
// ============================================================================
//
// Since implementation is implicit, how do you verify a type implements
// an interface? Use a compile-time check!
//
// var _ InterfaceName = TypeName{}
// var _ InterfaceName = (*TypeName)(nil)
//
// This creates no runtime overhead - it's checked at compile time.
//
// ============================================================================

// Writer interface for demonstration
type Writer interface {
	Write(data string) error
}

// FileWriter implements Writer
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data string) error {
	fmt.Printf("Writing to %s: %s\n", fw.Filename, data)
	return nil
}

// Compile-time verification that FileWriter implements Writer
var _ Writer = FileWriter{}

// ConsoleWriter also implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) error {
	fmt.Println("Console:", data)
	return nil
}

// Compile-time verification
var _ Writer = ConsoleWriter{}

// ============================================================================
// POINTER VS VALUE RECEIVERS WITH INTERFACES
// ============================================================================

// Counter interface
type Counter interface {
	Increment()
	Value() int
}

// ValueCounter uses value receiver - WON'T modify original
type ValueCounter struct {
	count int
}

func (vc ValueCounter) Increment() {
	vc.count++ // This modifies a COPY, not the original!
}

func (vc ValueCounter) Value() int {
	return vc.count
}

// PointerCounter uses pointer receiver - WILL modify original
type PointerCounter struct {
	count int
}

func (pc *PointerCounter) Increment() {
	pc.count++ // This modifies the actual struct
}

func (pc *PointerCounter) Value() int {
	return pc.count
}

// Note: *PointerCounter implements Counter, not PointerCounter

func main() {
	fmt.Println("=== Implicit Implementation ===")
	fmt.Println()

	// Types implement interfaces automatically
	fmt.Println("--- Implicit Interface Implementation ---")
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
		Year:   2015,
	}
	movie := Movie{
		Title:    "Inception",
		Director: "Christopher Nolan",
		Rating:   8.8,
	}

	PrintItem(book)
	PrintItem(movie)

	fmt.Println()

	// Standard library integration
	fmt.Println("--- fmt.Stringer Integration ---")
	temp := Temperature{Celsius: 25}
	// fmt.Println automatically calls String() method!
	fmt.Println("Current temperature:", temp)

	fmt.Println()

	// Using our Writer interface
	fmt.Println("--- Writer Interface ---")
	var w Writer

	w = FileWriter{Filename: "data.txt"}
	w.Write("Hello, file!")

	w = ConsoleWriter{}
	w.Write("Hello, console!")

	fmt.Println()

	// Pointer vs Value receivers
	fmt.Println("--- Pointer vs Value Receiver ---")

	// Value receiver - increment doesn't persist
	vc := ValueCounter{}
	vc.Increment()
	vc.Increment()
	vc.Increment()
	fmt.Printf("ValueCounter after 3 increments: %d (stays 0!)\n", vc.Value())

	// Pointer receiver - increment persists
	pc := &PointerCounter{}
	pc.Increment()
	pc.Increment()
	pc.Increment()
	fmt.Printf("PointerCounter after 3 increments: %d\n", pc.Value())

	fmt.Println()

	// Interface variable assignment
	fmt.Println("--- Interface Variable Types ---")
	var counter Counter

	// This works: *PointerCounter implements Counter
	counter = &PointerCounter{}
	counter.Increment()
	counter.Increment()
	fmt.Printf("Counter value: %d\n", counter.Value())

	// ValueCounter also implements Counter (methods have value receiver)
	counter = ValueCounter{}
	counter.Increment() // But this still won't persist due to value receiver
	fmt.Printf("ValueCounter via interface: %d\n", counter.Value())
}

// ============================================================================
// TO RUN:
//   go run day10/02_implicit_implementation.go
//
// EXPECTED OUTPUT:
//   === Implicit Implementation ===
//
//   --- Implicit Interface Implementation ---
//   Item: The Go Programming Language by Donovan & Kernighan (2015)
//   Item: Inception directed by Christopher Nolan - Rating: 8.8
//
//   --- fmt.Stringer Integration ---
//   Current temperature: 25.0°C (77.0°F)
//
//   --- Writer Interface ---
//   Writing to data.txt: Hello, file!
//   Console: Hello, console!
//
//   --- Pointer vs Value Receiver ---
//   ValueCounter after 3 increments: 0 (stays 0!)
//   PointerCounter after 3 increments: 3
//
//   --- Interface Variable Types ---
//   Counter value: 2
//   ValueCounter via interface: 0
//
// EXERCISE:
//   1. Create a Logger interface with Log(message string) method
//   2. Implement FileLogger and ConsoleLogger
//   3. Add compile-time verification for both implementations
//   4. Create a BankAccount with pointer receiver methods for Deposit/Withdraw
//
// KEY POINTS:
//   - Go uses implicit interface implementation (no "implements" keyword)
//   - Types implement interfaces by having the required methods
//   - Use var _ Interface = Type{} for compile-time verification
//   - Pointer receivers create *T implements Interface, not T
//   - Value receivers allow both T and *T to implement the interface
//   - Implement standard library interfaces for seamless integration
// ============================================================================
