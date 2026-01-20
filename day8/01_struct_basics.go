// Day 8, Exercise 1: Struct Basics
//
// Key concepts:
// - Structs group related data together
// - Define with type Name struct { fields }
// - Access fields with dot notation
// - Zero value initializes all fields to their zero values

package main

import "fmt"

// Person is a struct that groups related data
type Person struct {
	Name string
	Age  int
	City string
}

// Point represents a 2D coordinate
type Point struct {
	X, Y float64 // Multiple fields of same type
}

// Rectangle uses composition
type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

func main() {
	fmt.Println("=== What is a Struct? ===")

	// A struct is a collection of fields
	// Think of it as a blueprint for grouping related data

	// Creating a struct - Method 1: Field names (recommended)
	alice := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Println("Alice:", alice)

	// Creating a struct - Method 2: Positional (not recommended)
	// Must provide ALL fields in order
	bob := Person{"Bob", 25, "Boston"}
	fmt.Println("Bob:", bob)

	// Creating a struct - Method 3: Partial initialization
	// Unspecified fields get zero values
	charlie := Person{Name: "Charlie"}
	fmt.Println("Charlie:", charlie) // Age=0, City=""

	fmt.Println("\n=== Accessing Fields ===")

	// Use dot notation to access fields
	fmt.Println("Alice's name:", alice.Name)
	fmt.Println("Alice's age:", alice.Age)

	// Modify fields
	alice.Age = 31
	alice.City = "San Francisco"
	fmt.Println("Alice updated:", alice)

	fmt.Println("\n=== Zero Value of Structs ===")

	// Zero value: all fields set to their zero values
	var empty Person
	fmt.Println("Empty person:", empty)
	fmt.Printf("Name: %q, Age: %d, City: %q\n", empty.Name, empty.Age, empty.City)

	fmt.Println("\n=== Struct Literals ===")

	// Point struct
	origin := Point{0, 0}
	p1 := Point{X: 3, Y: 4}
	fmt.Println("Origin:", origin)
	fmt.Println("Point 1:", p1)

	fmt.Println("\n=== Nested Structs ===")

	// Structs can contain other structs
	rect := Rectangle{
		TopLeft:     Point{X: 0, Y: 10},
		BottomRight: Point{X: 10, Y: 0},
	}
	fmt.Println("Rectangle:", rect)

	// Access nested fields with chained dots
	fmt.Println("Top-left X:", rect.TopLeft.X)
	fmt.Println("Bottom-right Y:", rect.BottomRight.Y)

	// Calculate width and height
	width := rect.BottomRight.X - rect.TopLeft.X
	height := rect.TopLeft.Y - rect.BottomRight.Y
	fmt.Printf("Width: %.1f, Height: %.1f\n", width, height)

	fmt.Println("\n=== Comparing Structs ===")

	// Structs are comparable if all fields are comparable
	p2 := Point{X: 3, Y: 4}
	p3 := Point{X: 3, Y: 4}
	p4 := Point{X: 5, Y: 6}

	fmt.Println("p2 == p3:", p2 == p3) // true
	fmt.Println("p2 == p4:", p2 == p4) // false

	fmt.Println("\n=== Copying Structs ===")

	// Structs are VALUE types - assignment copies all fields
	original := Person{Name: "Dave", Age: 40, City: "Denver"}
	copied := original // This is a COPY

	copied.Name = "David"
	copied.Age = 41

	fmt.Println("Original:", original) // Unchanged!
	fmt.Println("Copied:", copied)     // Changed

	fmt.Println("\n=== Anonymous Structs ===")

	// Quick one-off struct without defining a type
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Server: %s:%d\n", config.Host, config.Port)
}

// TO RUN: go run day8/01_struct_basics.go
//
// OUTPUT:
// === What is a Struct? ===
// Alice: {Alice 30 New York}
// Bob: {Bob 25 Boston}
// Charlie: {Charlie 0 }
// ...
//
// EXERCISE:
// 1. Create a Book struct with Title, Author, Pages, Price fields
// 2. Create 3 books using different initialization methods
// 3. Modify one book's price
// 4. Compare two books with the same values
//
// KEY POINTS:
// - Structs group related data together
// - Use field names for clarity: Person{Name: "Alice"}
// - Zero value initializes all fields
// - Structs are value types (assignment copies)
// - Use dot notation to access/modify fields
