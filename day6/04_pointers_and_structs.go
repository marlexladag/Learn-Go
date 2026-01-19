// Day 6, Exercise 4: Pointers and Structs
//
// Key concepts:
// - Creating struct pointers with & and new()
// - Automatic dereferencing for struct fields
// - Struct literals return values, not pointers (use &)
// - Nested struct pointers

package main

import "fmt"

func main() {
	fmt.Println("=== Creating Struct Pointers ===")

	// Method 1: Using & with struct literal
	p1 := &Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("p1: %+v (type: %T)\n", p1, p1)

	// Method 2: Using new() - returns pointer to zero value
	p2 := new(Person)
	p2.Name = "Bob"
	p2.Age = 25
	fmt.Printf("p2: %+v\n", p2)

	// Method 3: Two-step (less common)
	p3 := Person{Name: "Charlie", Age: 35}
	p3Ptr := &p3
	fmt.Printf("p3Ptr: %+v\n", p3Ptr)

	fmt.Println("\n=== Automatic Dereferencing ===")

	user := &User{
		ID:    1,
		Email: "user@example.com",
	}

	// Go automatically dereferences struct pointers
	// These are equivalent:
	fmt.Println("Email (auto):", user.Email)
	fmt.Println("Email (explicit):", (*user).Email)

	// Modification also auto-dereferences
	user.Email = "new@example.com"
	fmt.Println("Updated email:", user.Email)

	fmt.Println("\n=== Pointer vs Value Semantics ===")

	// Value: creates independent copy
	original := Person{Name: "Original", Age: 20}
	copyVal := original
	copyVal.Age = 99
	fmt.Println("Original (unchanged):", original.Age)
	fmt.Println("Copy:", copyVal.Age)

	// Pointer: shares the same data
	original2 := &Person{Name: "Original2", Age: 20}
	copyPtr := original2 // Both point to same struct
	copyPtr.Age = 99
	fmt.Println("Original2 (changed!):", original2.Age)
	fmt.Println("CopyPtr:", copyPtr.Age)

	fmt.Println("\n=== Nested Struct Pointers ===")

	company := &Company{
		Name: "TechCorp",
		CEO: &Person{
			Name: "Jane",
			Age:  45,
		},
		Employees: []*Person{
			{Name: "Worker1", Age: 28},
			{Name: "Worker2", Age: 32},
		},
	}

	fmt.Println("Company:", company.Name)
	fmt.Println("CEO:", company.CEO.Name)
	fmt.Println("Employees:")
	for _, emp := range company.Employees {
		fmt.Printf("  - %s (%d)\n", emp.Name, emp.Age)
	}

	// Modifying nested pointer affects original
	company.CEO.Age = 46
	fmt.Println("CEO age updated:", company.CEO.Age)

	fmt.Println("\n=== Nil Struct Fields ===")

	car := &Car{
		Brand: "Toyota",
		// Engine is nil
	}

	fmt.Println("Car brand:", car.Brand)
	fmt.Println("Engine is nil:", car.Engine == nil)

	// Check before accessing
	if car.Engine != nil {
		fmt.Println("Horsepower:", car.Engine.Horsepower)
	} else {
		fmt.Println("No engine installed")
	}

	// Add engine
	car.Engine = &Engine{Horsepower: 200}
	fmt.Println("After adding engine:", car.Engine.Horsepower, "hp")

	fmt.Println("\n=== Constructor Pattern ===")

	// Use constructor functions to ensure proper initialization
	config := NewConfig("production")
	fmt.Printf("Config: %+v\n", config)

	// With options
	configWithOptions := NewConfigWithOptions("staging", &ConfigOptions{
		Timeout: 60,
		Debug:   true,
	})
	fmt.Printf("Config with options: %+v\n", configWithOptions)

	fmt.Println("\n=== Linked Structures ===")

	// Pointers enable linked data structures
	list := &Node{Value: 1}
	list.Next = &Node{Value: 2}
	list.Next.Next = &Node{Value: 3}

	// Traverse the list
	fmt.Print("Linked list: ")
	current := list
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()

	fmt.Println("\n=== Self-Referential Structures ===")

	// Tree structure
	root := &TreeNode{
		Value: 10,
		Left: &TreeNode{
			Value: 5,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 7},
		},
		Right: &TreeNode{
			Value: 15,
			Right: &TreeNode{Value: 20},
		},
	}

	fmt.Println("Tree root:", root.Value)
	fmt.Println("Left child:", root.Left.Value)
	fmt.Println("Right child:", root.Right.Value)
}

// Person is a simple struct
type Person struct {
	Name string
	Age  int
}

// User has an ID and email
type User struct {
	ID    int
	Email string
}

// Company contains pointer fields
type Company struct {
	Name      string
	CEO       *Person
	Employees []*Person
}

// Car demonstrates optional pointer fields
type Car struct {
	Brand  string
	Engine *Engine
}

type Engine struct {
	Horsepower int
}

// Config with constructor pattern
type Config struct {
	Environment string
	Timeout     int
	Debug       bool
}

type ConfigOptions struct {
	Timeout int
	Debug   bool
}

// NewConfig creates a Config with defaults
func NewConfig(env string) *Config {
	return &Config{
		Environment: env,
		Timeout:     30, // Default
		Debug:       false,
	}
}

// NewConfigWithOptions creates a Config with custom options
func NewConfigWithOptions(env string, opts *ConfigOptions) *Config {
	config := NewConfig(env)
	if opts != nil {
		if opts.Timeout > 0 {
			config.Timeout = opts.Timeout
		}
		config.Debug = opts.Debug
	}
	return config
}

// Node for linked list
type Node struct {
	Value int
	Next  *Node
}

// TreeNode for binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// TO RUN: go run day6/04_pointers_and_structs.go
//
// OUTPUT:
// === Creating Struct Pointers ===
// p1: &{Name:Alice Age:30} (type: *main.Person)
// ...
//
// EXERCISE:
// 1. Create a Book struct with Title, Author (*Person), and Pages
// 2. Write a function that takes *Book and updates the page count
// 3. Create a Library struct containing []*Book
// 4. Write methods to add/remove books from the library
//
// KEY POINTS:
// - Use & with struct literal or new() for pointers
// - Go auto-dereferences: ptr.Field works without (*ptr).Field
// - Pointer fields can be nil - always check before access
// - Constructor functions ensure proper initialization
// - Pointers enable linked/tree structures
