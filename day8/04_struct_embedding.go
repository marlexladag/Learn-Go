// Day 8, Exercise 4: Struct Embedding (Composition)
//
// Key concepts:
// - Embedding promotes fields and methods
// - Go uses composition over inheritance
// - Embedded fields can be accessed directly
// - Multiple embeddings are allowed

package main

import "fmt"

// Base types
type Address struct {
	Street  string
	City    string
	Country string
}

type ContactInfo struct {
	Email string
	Phone string
}

// Person embeds Address - composition
type Person struct {
	Name string
	Age  int
	Address // Embedded (no field name)
}

// Employee embeds Person and ContactInfo
type Employee struct {
	Person // Embedded
	ContactInfo
	EmployeeID string
	Department string
}

// Traditional composition (with field name)
type Company struct {
	Name    string
	Address Address // Named field, not embedded
}

func main() {
	fmt.Println("=== Basic Embedding ===")

	// Create a person with embedded address
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
	}

	// Access embedded fields DIRECTLY
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.City)    // Promoted from Address!
	fmt.Println("Street:", p.Street) // Also promoted

	// Can still access via embedded type name
	fmt.Println("Full address:", p.Address)

	fmt.Println("\n=== Embedding vs Named Fields ===")

	// With embedding: fields are "promoted"
	fmt.Printf("Person's city: %s\n", p.City) // Direct access

	// With named field: must use field name
	company := Company{
		Name: "Acme Corp",
		Address: Address{
			Street:  "456 Corp Ave",
			City:    "New York",
			Country: "USA",
		},
	}
	fmt.Printf("Company's city: %s\n", company.Address.City) // Must use .Address

	fmt.Println("\n=== Multiple Embedding ===")

	// Employee embeds both Person and ContactInfo
	emp := Employee{
		Person: Person{
			Name: "Bob",
			Age:  35,
			Address: Address{
				Street:  "789 Work St",
				City:    "Seattle",
				Country: "USA",
			},
		},
		ContactInfo: ContactInfo{
			Email: "bob@company.com",
			Phone: "555-1234",
		},
		EmployeeID: "E001",
		Department: "Engineering",
	}

	// All fields are accessible directly!
	fmt.Println("Employee:", emp.Name)        // From Person
	fmt.Println("City:", emp.City)            // From Person.Address
	fmt.Println("Email:", emp.Email)          // From ContactInfo
	fmt.Println("Department:", emp.Department) // Own field

	fmt.Println("\n=== Method Promotion ===")

	// Methods from embedded types are also promoted
	addr := Address{
		Street:  "100 Example St",
		City:    "Chicago",
		Country: "USA",
	}
	fmt.Println("Address formatted:", addr.Format())

	// Person can use Address's methods directly
	person := Person{
		Name: "Charlie",
		Age:  25,
		Address: Address{
			Street:  "200 Demo Ave",
			City:    "Denver",
			Country: "USA",
		},
	}
	fmt.Println("Person's address:", person.Format()) // Promoted method!

	fmt.Println("\n=== Overriding Promoted Methods ===")

	// Employee can override Person's methods
	fmt.Println("Person's intro:", person.Introduce())

	emp2 := Employee{
		Person: Person{
			Name: "Diana",
			Age:  40,
		},
		EmployeeID: "E002",
		Department: "Sales",
	}
	fmt.Println("Employee's intro:", emp2.Introduce()) // Overridden!

	// Can still access the embedded method explicitly
	fmt.Println("As person:", emp2.Person.Introduce())

	fmt.Println("\n=== Practical Example: Game Entity ===")

	// Base position type
	player := Player{
		Position: Position{X: 100, Y: 200},
		Name:     "Hero",
		Health:   100,
	}

	enemy := Enemy{
		Position: Position{X: 50, Y: 75},
		Type:     "Goblin",
		Damage:   10,
	}

	fmt.Printf("Player at (%d, %d)\n", player.X, player.Y)
	fmt.Printf("Enemy at (%d, %d)\n", enemy.X, enemy.Y)

	// Both can use Position's methods
	player.Move(10, 5)
	enemy.Move(-5, 10)

	fmt.Printf("After move - Player at (%d, %d)\n", player.X, player.Y)
	fmt.Printf("After move - Enemy at (%d, %d)\n", enemy.X, enemy.Y)
}

// ============ Methods ============

// Format returns a formatted address string
func (a Address) Format() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

// Introduce returns an introduction
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s, %d years old", p.Name, p.Age)
}

// Introduce overrides Person's Introduce for Employee
func (e Employee) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s from %s (ID: %s)", e.Name, e.Department, e.EmployeeID)
}

// ============ Game Example ============

type Position struct {
	X, Y int
}

func (p *Position) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

type Player struct {
	Position // Embedded
	Name     string
	Health   int
}

type Enemy struct {
	Position // Embedded
	Type     string
	Damage   int
}

// TO RUN: go run day8/04_struct_embedding.go
//
// KEY POINTS:
// - Embedding promotes fields: person.City instead of person.Address.City
// - Methods are also promoted
// - Use embedding for "has-a" relationships
// - Go favors composition over inheritance
// - Can override promoted methods
//
// EXERCISE:
// 1. Create a Vehicle struct with Speed and MaxSpeed fields
// 2. Create Car and Boat that embed Vehicle
// 3. Add an Accelerate(amount int) method to Vehicle
// 4. Add specific methods to Car and Boat
