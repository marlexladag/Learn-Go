// Day 5, Exercise 3: Maps with Structs
//
// Key concepts:
// - Maps can have structs as values (very common pattern)
// - Structs can be used as keys if all fields are comparable
// - Maps of structs are great for storing records/entities
// - Cannot directly modify struct fields in map (must reassign)

package main

import "fmt"

// Person represents a person record
type Person struct {
	Name    string
	Age     int
	Email   string
	Country string
}

// Point represents a 2D coordinate (can be used as map key)
type Point struct {
	X, Y int
}

func main() {
	fmt.Println("=== Maps with Struct Values ===")

	// Map from ID to Person
	users := make(map[int]Person)

	users[1] = Person{Name: "Alice", Age: 30, Email: "alice@example.com", Country: "USA"}
	users[2] = Person{Name: "Bob", Age: 25, Email: "bob@example.com", Country: "UK"}
	users[3] = Person{Name: "Charlie", Age: 35, Email: "charlie@example.com", Country: "Canada"}

	fmt.Println("User 1:", users[1])
	fmt.Println("User 2's name:", users[2].Name)
	fmt.Println("User 3's email:", users[3].Email)

	fmt.Println("\n=== Map Literal with Structs ===")

	employees := map[string]Person{
		"E001": {Name: "Diana", Age: 28, Email: "diana@corp.com", Country: "Germany"},
		"E002": {Name: "Eve", Age: 32, Email: "eve@corp.com", Country: "France"},
		"E003": {Name: "Frank", Age: 45, Email: "frank@corp.com", Country: "Japan"},
	}

	for id, emp := range employees {
		fmt.Printf("%s: %s (%d years old) from %s\n", id, emp.Name, emp.Age, emp.Country)
	}

	fmt.Println("\n=== Updating Struct in Map ===")

	// IMPORTANT: Cannot directly modify struct field in map
	// users[1].Age = 31  // This does NOT work!

	// Must retrieve, modify, and reassign
	user := users[1]
	user.Age = 31
	users[1] = user

	fmt.Println("Updated user 1:", users[1])

	// Or use map of pointers for direct modification
	userPtrs := make(map[int]*Person)
	userPtrs[1] = &Person{Name: "Grace", Age: 29, Email: "grace@example.com", Country: "Australia"}

	// Now can modify directly
	userPtrs[1].Age = 30
	fmt.Println("User pointer 1:", *userPtrs[1])

	fmt.Println("\n=== Struct as Map Key ===")

	// Structs with comparable fields can be map keys
	grid := make(map[Point]string)

	grid[Point{0, 0}] = "Origin"
	grid[Point{1, 0}] = "East"
	grid[Point{0, 1}] = "North"
	grid[Point{1, 1}] = "Northeast"

	fmt.Println("At (0,0):", grid[Point{0, 0}])
	fmt.Println("At (1,1):", grid[Point{1, 1}])

	// Check if position exists
	if label, exists := grid[Point{2, 2}]; exists {
		fmt.Println("Found:", label)
	} else {
		fmt.Println("Position (2,2) not mapped")
	}

	fmt.Println("\n=== Practical Example: Student Database ===")

	type Student struct {
		Name   string
		Grade  float64
		Passed bool
	}

	students := map[string]Student{
		"STU001": {Name: "Alice", Grade: 92.5, Passed: true},
		"STU002": {Name: "Bob", Grade: 58.0, Passed: false},
		"STU003": {Name: "Charlie", Grade: 75.5, Passed: true},
		"STU004": {Name: "Diana", Grade: 88.0, Passed: true},
	}

	// Find all passing students
	fmt.Println("Passing students:")
	for id, student := range students {
		if student.Passed {
			fmt.Printf("  %s: %s (%.1f%%)\n", id, student.Name, student.Grade)
		}
	}

	// Calculate class average
	var totalGrade float64
	for _, student := range students {
		totalGrade += student.Grade
	}
	average := totalGrade / float64(len(students))
	fmt.Printf("Class average: %.2f%%\n", average)

	// Find highest grade
	var topStudent Student
	var topID string
	for id, student := range students {
		if student.Grade > topStudent.Grade {
			topStudent = student
			topID = id
		}
	}
	fmt.Printf("Top student: %s (%s) with %.1f%%\n", topStudent.Name, topID, topStudent.Grade)

	fmt.Println("\n=== Nested Maps ===")

	// Map of maps: department -> employee name -> salary
	salaries := map[string]map[string]int{
		"Engineering": {
			"Alice": 90000,
			"Bob":   85000,
		},
		"Marketing": {
			"Charlie": 75000,
			"Diana":   80000,
		},
	}

	fmt.Println("Alice's salary:", salaries["Engineering"]["Alice"])

	// Adding to nested map - must check if inner map exists
	if salaries["HR"] == nil {
		salaries["HR"] = make(map[string]int)
	}
	salaries["HR"]["Eve"] = 70000

	fmt.Println("HR department:", salaries["HR"])

	// Print all salaries
	fmt.Println("\nAll salaries:")
	for dept, employees := range salaries {
		fmt.Printf("  %s:\n", dept)
		for name, salary := range employees {
			fmt.Printf("    %s: $%d\n", name, salary)
		}
	}
}

// TO RUN: go run day5/03_maps_with_structs.go
//
// OUTPUT:
// === Maps with Struct Values ===
// User 1: {Alice 30 alice@example.com USA}
// User 2's name: Bob
// ...
//
// EXERCISE:
// 1. Create a Product struct with Name, Price, and Quantity fields
// 2. Create a map of product IDs to Products
// 3. Add 5 products to the map
// 4. Calculate the total inventory value (Price * Quantity for each)
// 5. Find the most expensive product
//
// KEY POINTS:
// - Struct values are common (map[string]Person)
// - Cannot modify struct fields directly in map - must reassign
// - Use map of pointers (*Person) for direct modification
// - Structs with comparable fields can be keys
// - Nested maps require nil checks before adding
