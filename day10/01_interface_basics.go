package main

import "fmt"

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 1: Interface Basics
// ============================================================================
//
// WHAT IS AN INTERFACE?
// An interface is a collection of method signatures. It defines behavior
// without specifying how that behavior is implemented.
//
// Think of an interface as a contract: "Any type that has these methods
// can be used wherever this interface is expected."
//
// KEY INSIGHT: Interfaces focus on WHAT a type can do, not WHAT it is.
// ============================================================================

// Speaker defines the behavior of anything that can speak
type Speaker interface {
	Speak() string
}

// Dog is a concrete type that implements Speaker
type Dog struct {
	Name string
}

// Speak makes Dog implement the Speaker interface
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Cat is another concrete type that implements Speaker
type Cat struct {
	Name string
}

// Speak makes Cat implement the Speaker interface
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

// Robot also implements Speaker
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: Beep boop!", r.Model)
}

// MakeSpeak accepts ANY type that implements Speaker
// This is the power of interfaces - polymorphism!
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// ============================================================================
// INTERFACES WITH MULTIPLE METHODS
// ============================================================================

// Shape defines behavior for geometric shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// PrintShapeInfo works with ANY shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ============================================================================
// WHY USE INTERFACES?
// ============================================================================
//
// 1. ABSTRACTION: Hide implementation details
// 2. POLYMORPHISM: Same function works with different types
// 3. TESTABILITY: Easy to create mock implementations for testing
// 4. FLEXIBILITY: Add new types without changing existing code
// 5. DECOUPLING: Code depends on behavior, not concrete types
//
// ============================================================================

func main() {
	fmt.Println("=== Interface Basics ===")
	fmt.Println()

	// Different types, same interface
	fmt.Println("--- Speaker Interface ---")
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	robot := Robot{Model: "T-1000"}

	// All can be passed to MakeSpeak because they implement Speaker
	MakeSpeak(dog)
	MakeSpeak(cat)
	MakeSpeak(robot)

	fmt.Println()

	// Store different types in a slice of interfaces
	fmt.Println("--- Slice of Speakers ---")
	speakers := []Speaker{dog, cat, robot}
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}

	fmt.Println()

	// Shape interface with multiple methods
	fmt.Println("--- Shape Interface ---")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Print("Rectangle - ")
	PrintShapeInfo(rect)

	fmt.Print("Circle - ")
	PrintShapeInfo(circle)

	fmt.Println()

	// Slice of shapes
	fmt.Println("--- All Shapes ---")
	shapes := []Shape{rect, circle}
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
}

// ============================================================================
// TO RUN:
//   go run day10/01_interface_basics.go
//
// EXPECTED OUTPUT:
//   === Interface Basics ===
//
//   --- Speaker Interface ---
//   Buddy says: Woof!
//   Whiskers says: Meow!
//   Robot T-1000 says: Beep boop!
//
//   --- Slice of Speakers ---
//   Buddy says: Woof!
//   Whiskers says: Meow!
//   Robot T-1000 says: Beep boop!
//
//   --- Shape Interface ---
//   Rectangle - Area: 50.00, Perimeter: 30.00
//   Circle - Area: 153.94, Perimeter: 43.98
//
//   --- All Shapes ---
//   Total area of all shapes: 203.94
//
// EXERCISE:
//   1. Add a Triangle struct that implements Shape
//   2. Create a Person struct that implements Speaker
//   3. Write a function that finds the largest shape by area
//
// KEY POINTS:
//   - An interface is a set of method signatures
//   - Any type with matching methods implements the interface
//   - Interfaces enable polymorphism in Go
//   - Functions can accept interface types for flexibility
//   - Slices can hold different types via interface
// ============================================================================
