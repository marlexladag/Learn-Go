// Day 8, Exercise 2: Structs and Functions
//
// Key concepts:
// - Pass structs to functions by value (copy)
// - Pass structs by pointer to modify them
// - Return structs from functions
// - Constructor pattern: NewType() functions

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	X, Y   float64 // Center coordinates
	Radius float64
}

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== Passing Structs by Value ===")

	c1 := Circle{X: 0, Y: 0, Radius: 5}
	fmt.Println("Before:", c1)

	// This function receives a COPY
	tryToModify(c1)
	fmt.Println("After tryToModify:", c1) // Unchanged!

	fmt.Println("\n=== Passing Structs by Pointer ===")

	c2 := Circle{X: 0, Y: 0, Radius: 5}
	fmt.Println("Before:", c2)

	// Pass address to modify the original
	modifyRadius(&c2, 10)
	fmt.Println("After modifyRadius:", c2) // Changed!

	fmt.Println("\n=== Why Use Pointers with Structs? ===")

	// 1. To modify the original struct
	person := Person{Name: "Alice", Age: 30}
	birthday(&person)
	fmt.Println("After birthday:", person)

	// 2. To avoid copying large structs (performance)
	largeStruct := createLargeData()
	processData(&largeStruct) // Efficient - no copy

	fmt.Println("\n=== Returning Structs from Functions ===")

	// Return by value (common for small structs)
	origin := makePoint(0, 0)
	fmt.Println("Origin:", origin)

	// Return pointer (common for larger structs)
	newCircle := NewCircle(5, 5, 3)
	fmt.Println("New circle:", *newCircle)

	fmt.Println("\n=== Constructor Pattern ===")

	// "Constructor" functions initialize structs properly
	p1 := NewPerson("Bob", 25)
	fmt.Println("Person 1:", *p1)

	// With validation
	p2, err := NewPersonValidated("", -5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	p3, err := NewPersonValidated("Charlie", 30)
	if err == nil {
		fmt.Println("Person 3:", *p3)
	}

	fmt.Println("\n=== Working with Struct Fields ===")

	// Calculate area and circumference
	circle := Circle{X: 0, Y: 0, Radius: 5}
	area := circleArea(circle)
	circumference := circleCircumference(circle)

	fmt.Printf("Circle with radius %.1f:\n", circle.Radius)
	fmt.Printf("  Area: %.2f\n", area)
	fmt.Printf("  Circumference: %.2f\n", circumference)

	fmt.Println("\n=== Multiple Structs as Parameters ===")

	// Calculate distance between two circles
	c3 := Circle{X: 0, Y: 0, Radius: 5}
	c4 := Circle{X: 10, Y: 0, Radius: 3}

	dist := distanceBetweenCircles(c3, c4)
	fmt.Printf("Distance between circle centers: %.2f\n", dist)

	touching := circlesOverlap(c3, c4)
	fmt.Printf("Circles overlap: %v\n", touching)
}

// tryToModify receives a COPY - changes don't affect original
func tryToModify(c Circle) {
	c.Radius = 100
	fmt.Println("Inside function:", c)
}

// modifyRadius receives a POINTER - changes affect original
func modifyRadius(c *Circle, newRadius float64) {
	c.Radius = newRadius // Go automatically dereferences
	// Same as: (*c).Radius = newRadius
}

// birthday increments age through pointer
func birthday(p *Person) {
	p.Age++
}

// makePoint returns a struct by value
func makePoint(x, y float64) struct{ X, Y float64 } {
	return struct{ X, Y float64 }{x, y}
}

// NewCircle is a constructor that returns a pointer
// Convention: functions starting with New return pointers
func NewCircle(x, y, radius float64) *Circle {
	return &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
}

// NewPerson is a simple constructor
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// NewPersonValidated is a constructor with validation
func NewPersonValidated(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	return &Person{Name: name, Age: age}, nil
}

// circleArea calculates the area
func circleArea(c Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}

// circleCircumference calculates the circumference
func circleCircumference(c Circle) float64 {
	return 2 * math.Pi * c.Radius
}

// distanceBetweenCircles calculates distance between centers
func distanceBetweenCircles(c1, c2 Circle) float64 {
	dx := c2.X - c1.X
	dy := c2.Y - c1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// circlesOverlap checks if two circles overlap
func circlesOverlap(c1, c2 Circle) bool {
	dist := distanceBetweenCircles(c1, c2)
	return dist < (c1.Radius + c2.Radius)
}

// LargeData simulates a large struct
type LargeData struct {
	Data [1000]int
}

func createLargeData() LargeData {
	return LargeData{}
}

func processData(d *LargeData) {
	// Process without copying
	d.Data[0] = 42
}

// TO RUN: go run day8/02_structs_and_functions.go
//
// KEY POINTS:
// - Pass by value for small structs you don't need to modify
// - Pass by pointer to modify the original or avoid copying large structs
// - Constructor pattern: NewType() returns *Type
// - Go auto-dereferences struct pointers: p.Field works for *Type
//
// EXERCISE:
// 1. Create a Rectangle struct with Width, Height
// 2. Write a function to calculate area (pass by value)
// 3. Write a function to scale the rectangle (pass by pointer)
// 4. Write a NewRectangle constructor with validation
