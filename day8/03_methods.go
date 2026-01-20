// Day 8, Exercise 3: Methods
//
// Key concepts:
// - Methods are functions with a receiver
// - Receiver appears between func and method name
// - Value receiver: method gets a copy
// - Pointer receiver: method can modify the struct
// - Methods make code more readable and organized

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	X, Y   float64
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Counter struct {
	value int
}

func main() {
	fmt.Println("=== What is a Method? ===")

	// A method is a function with a receiver
	// Instead of: circleArea(c Circle) float64
	// We write:   func (c Circle) Area() float64

	c := Circle{X: 0, Y: 0, Radius: 5}

	// Call method with dot notation
	area := c.Area()
	fmt.Printf("Circle with radius %.1f has area %.2f\n", c.Radius, area)

	fmt.Println("\n=== Value Receiver vs Pointer Receiver ===")

	// Value receiver: method gets a COPY
	// Use when: you don't need to modify the struct

	// Pointer receiver: method gets the ORIGINAL
	// Use when: you need to modify the struct

	counter := Counter{value: 0}
	fmt.Println("Initial value:", counter.value)

	// Value receiver - doesn't modify original
	counter.IncrementCopy() // This won't work!
	fmt.Println("After IncrementCopy:", counter.value)

	// Pointer receiver - modifies original
	counter.Increment()
	counter.Increment()
	fmt.Println("After Increment x2:", counter.value)

	fmt.Println("\n=== Methods are More Readable ===")

	// Compare these two approaches:

	// Function style:
	rect1 := Rectangle{Width: 10, Height: 5}
	area1 := rectangleArea(rect1) // Function
	fmt.Printf("Function: area = %.1f\n", area1)

	// Method style:
	rect2 := Rectangle{Width: 10, Height: 5}
	area2 := rect2.Area() // Method - more natural!
	fmt.Printf("Method: area = %.1f\n", area2)

	fmt.Println("\n=== Chaining Methods ===")

	// Methods can be chained when they return the receiver
	rect3 := Rectangle{Width: 10, Height: 5}
	rect3.Double().Double() // Chain calls
	fmt.Printf("After doubling twice: %+v\n", rect3)

	fmt.Println("\n=== Method Sets ===")

	// All methods on a type form its "method set"
	circle := Circle{X: 0, Y: 0, Radius: 5}

	fmt.Println("Circle methods:")
	fmt.Printf("  Area: %.2f\n", circle.Area())
	fmt.Printf("  Circumference: %.2f\n", circle.Circumference())
	fmt.Printf("  Diameter: %.2f\n", circle.Diameter())

	// Move the circle
	circle.Move(10, 20)
	fmt.Printf("  After Move(10, 20): center at (%.1f, %.1f)\n", circle.X, circle.Y)

	// Scale the circle
	circle.Scale(2)
	fmt.Printf("  After Scale(2): radius = %.1f\n", circle.Radius)

	fmt.Println("\n=== When to Use Pointer Receivers ===")

	// Use pointer receiver when:
	// 1. Method needs to modify the receiver
	// 2. Struct is large (avoid copying)
	// 3. Consistency: if one method needs pointer, use pointer for all

	// Rule of thumb: if in doubt, use pointer receiver

	fmt.Println("\n=== Methods Can Have Any Type ===")

	// You can add methods to any type you define
	var temp Temperature = 100
	fmt.Printf("%.1f°C = %.1f°F\n", float64(temp), temp.ToFahrenheit())

	var msg MyString = "hello"
	fmt.Println("Shouting:", msg.Shout())
}

// ============ Circle Methods ============

// Area returns the area (value receiver - doesn't modify)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circumference returns the circumference (value receiver)
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Diameter returns the diameter (value receiver)
func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Move changes the center position (pointer receiver - modifies)
func (c *Circle) Move(dx, dy float64) {
	c.X += dx
	c.Y += dy
}

// Scale multiplies the radius (pointer receiver - modifies)
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// ============ Rectangle Methods ============

// Area returns the area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Double doubles the dimensions (returns pointer for chaining)
func (r *Rectangle) Double() *Rectangle {
	r.Width *= 2
	r.Height *= 2
	return r
}

// Traditional function for comparison
func rectangleArea(r Rectangle) float64 {
	return r.Width * r.Height
}

// ============ Counter Methods ============

// IncrementCopy tries to increment but gets a copy (won't work!)
func (c Counter) IncrementCopy() {
	c.value++ // This modifies the copy, not the original!
}

// Increment actually increments the counter
func (c *Counter) Increment() {
	c.value++
}

// Value returns the current value
func (c Counter) Value() int {
	return c.value
}

// ============ Custom Type Methods ============

// Temperature is a custom type based on float64
type Temperature float64

// ToFahrenheit converts Celsius to Fahrenheit
func (t Temperature) ToFahrenheit() float64 {
	return float64(t)*9/5 + 32
}

// MyString is a custom string type
type MyString string

// Shout returns the string in uppercase with exclamation
func (s MyString) Shout() string {
	return string(s) + "!!!"
}

// TO RUN: go run day8/03_methods.go
//
// KEY POINTS:
// - Methods are functions with a receiver: func (r Type) Name()
// - Value receiver (Type): gets a copy, can't modify original
// - Pointer receiver (*Type): can modify the original
// - Methods improve readability: circle.Area() vs circleArea(circle)
// - When in doubt, use pointer receiver
//
// EXERCISE:
// 1. Add an IsSquare() method to Rectangle
// 2. Add a Contains(x, y float64) method to Circle
// 3. Create a BankAccount struct with Balance field
// 4. Add Deposit and Withdraw methods (which receiver type?)
