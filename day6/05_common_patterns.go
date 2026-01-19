// Day 6, Exercise 5: Common Pointer Patterns
//
// Key concepts:
// - Optional values with pointers (nil = not set)
// - In-place modification patterns
// - Pointer slices vs slice of pointers
// - Common mistakes and how to avoid them

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("=== Optional Values Pattern ===")

	// Using pointers for optional fields
	user1 := User{
		Name: "Alice",
		Age:  intPtr(30), // Age is set
	}

	user2 := User{
		Name: "Bob",
		// Age is nil (not provided)
	}

	printUser(user1)
	printUser(user2)

	fmt.Println("\n=== JSON with Optional Fields ===")

	// Pointers allow distinguishing "not set" from "zero value"
	jsonData := []byte(`{"name": "Charlie"}`) // No age field

	var user3 User
	json.Unmarshal(jsonData, &user3)
	fmt.Printf("From JSON: Name=%s, Age=%v\n", user3.Name, user3.Age)

	jsonWithAge := []byte(`{"name": "David", "age": 0}`) // Age explicitly 0
	var user4 User
	json.Unmarshal(jsonWithAge, &user4)
	if user4.Age != nil {
		fmt.Printf("From JSON: Name=%s, Age=%d (explicitly set)\n", user4.Name, *user4.Age)
	}

	fmt.Println("\n=== Update Pattern ===")

	product := Product{
		ID:    1,
		Name:  "Widget",
		Price: 9.99,
		Stock: 100,
	}
	fmt.Printf("Before: %+v\n", product)

	// Partial update - only update provided fields
	update := ProductUpdate{
		Price: floatPtr(12.99), // Only updating price
		// Name and Stock are nil (not updated)
	}
	applyProductUpdate(&product, &update)
	fmt.Printf("After: %+v\n", product)

	fmt.Println("\n=== Slice of Pointers vs Pointer to Slice ===")

	// Slice of pointers: []*T
	// Each element is a pointer, useful for:
	// - Large structs (avoid copying)
	// - Shared references
	// - nil elements

	people := []*Person{
		{Name: "Alice"},
		{Name: "Bob"},
		nil, // Can have nil elements
		{Name: "Charlie"},
	}

	for i, p := range people {
		if p != nil {
			fmt.Printf("Person %d: %s\n", i, p.Name)
		} else {
			fmt.Printf("Person %d: <nil>\n", i)
		}
	}

	// Pointer to slice: *[]T
	// Used when function needs to modify the slice itself
	numbers := []int{1, 2, 3}
	appendIfEven(&numbers, 4)
	appendIfEven(&numbers, 5) // Won't append (odd)
	fmt.Println("Numbers:", numbers)

	fmt.Println("\n=== Common Mistake: Loop Variable Pointer ===")

	items := []string{"a", "b", "c"}
	var pointers []*string

	// WRONG: All pointers point to same variable!
	for _, item := range items {
		pointers = append(pointers, &item) // Bug!
	}
	fmt.Println("Wrong way (all same):")
	for _, p := range pointers {
		fmt.Printf("  %s\n", *p) // All print "c"!
	}

	// CORRECT: Create new variable in each iteration
	pointers = nil
	for _, item := range items {
		item := item // Shadow with new variable
		pointers = append(pointers, &item)
	}
	fmt.Println("Correct way:")
	for _, p := range pointers {
		fmt.Printf("  %s\n", *p)
	}

	// ALSO CORRECT: Use index
	pointers = nil
	for i := range items {
		pointers = append(pointers, &items[i])
	}
	fmt.Println("Using index:")
	for _, p := range pointers {
		fmt.Printf("  %s\n", *p)
	}

	fmt.Println("\n=== Safe Dereference Pattern ===")

	var maybeNil *int
	fmt.Println("Safe deref nil:", safeDeref(maybeNil, -1))

	value := 42
	maybeNil = &value
	fmt.Println("Safe deref value:", safeDeref(maybeNil, -1))

	fmt.Println("\n=== Builder Pattern with Pointers ===")

	server := NewServerConfig().
		WithHost("localhost").
		WithPort(8080).
		WithTimeout(30).
		Build()

	fmt.Printf("Server config: %+v\n", server)

	fmt.Println("\n=== Interface and Pointers ===")

	// Both value and pointer can implement interfaces
	// But pointer receivers only work with pointers

	var counter Counter

	// Value type works with value receiver methods
	counter = MyCounter{count: 5}
	fmt.Println("Value counter:", counter.Count())

	// Pointer type works with both
	counter = &MyCounter{count: 10}
	fmt.Println("Pointer counter:", counter.Count())

	// Increment requires pointer receiver
	mc := &MyCounter{count: 0}
	mc.Increment()
	mc.Increment()
	fmt.Println("After increments:", mc.Count())
}

// User with optional Age field
type User struct {
	Name string `json:"name"`
	Age  *int   `json:"age,omitempty"`
}

func printUser(u User) {
	if u.Age != nil {
		fmt.Printf("User: %s, Age: %d\n", u.Name, *u.Age)
	} else {
		fmt.Printf("User: %s, Age: not specified\n", u.Name)
	}
}

// Helper to create int pointer
func intPtr(i int) *int {
	return &i
}

func floatPtr(f float64) *float64 {
	return &f
}

// Product for update pattern
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type ProductUpdate struct {
	Name  *string
	Price *float64
	Stock *int
}

func applyProductUpdate(p *Product, u *ProductUpdate) {
	if u.Name != nil {
		p.Name = *u.Name
	}
	if u.Price != nil {
		p.Price = *u.Price
	}
	if u.Stock != nil {
		p.Stock = *u.Stock
	}
}

// Person for slice examples
type Person struct {
	Name string
}

func appendIfEven(s *[]int, val int) {
	if val%2 == 0 {
		*s = append(*s, val)
	}
}

// Safe dereference with default
func safeDeref(ptr *int, defaultVal int) int {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}

// ServerConfig with builder pattern
type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

type ServerBuilder struct {
	config ServerConfig
}

func NewServerConfig() *ServerBuilder {
	return &ServerBuilder{
		config: ServerConfig{
			Host:    "0.0.0.0",
			Port:    80,
			Timeout: 60,
		},
	}
}

func (b *ServerBuilder) WithHost(host string) *ServerBuilder {
	b.config.Host = host
	return b
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.config.Port = port
	return b
}

func (b *ServerBuilder) WithTimeout(timeout int) *ServerBuilder {
	b.config.Timeout = timeout
	return b
}

func (b *ServerBuilder) Build() ServerConfig {
	return b.config
}

// Counter interface example
type Counter interface {
	Count() int
}

type MyCounter struct {
	count int
}

func (c MyCounter) Count() int {
	return c.count
}

func (c *MyCounter) Increment() {
	c.count++
}

// TO RUN: go run day6/05_common_patterns.go
//
// OUTPUT:
// === Optional Values Pattern ===
// User: Alice, Age: 30
// User: Bob, Age: not specified
// ...
//
// EXERCISE:
// 1. Create a Settings struct with optional fields for Theme, FontSize, Language
// 2. Write an UpdateSettings function that only updates non-nil fields
// 3. Create a helper function strPtr(s string) *string
// 4. Test distinguishing between "not provided" and "empty string"
//
// KEY POINTS:
// - Use *T for optional fields (nil = not set)
// - Watch out for loop variable pointer capture bug
// - Use safe dereference pattern when nil is possible
// - Builder pattern works well with method chaining
// - Consider whether value or pointer semantics fit your use case
