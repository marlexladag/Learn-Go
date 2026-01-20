// Day 8, Exercise 5: Struct Tags
//
// Key concepts:
// - Struct tags are metadata attached to fields
// - Used by encoding packages (json, xml, etc.)
// - Format: `key:"value" key2:"value2"`
// - Access via reflect package (advanced)

package main

import (
	"encoding/json"
	"fmt"
)

// Person with JSON struct tags
type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email,omitempty"` // Omit if empty
	Password  string `json:"-"`               // Never include in JSON
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at,omitempty"`
}

// Config demonstrates multiple tag types
type Config struct {
	Host     string `json:"host" env:"SERVER_HOST" default:"localhost"`
	Port     int    `json:"port" env:"SERVER_PORT" default:"8080"`
	Debug    bool   `json:"debug" env:"DEBUG_MODE"`
	APIKey   string `json:"-" env:"API_KEY"` // Hidden from JSON
}

// APIResponse for nested JSON
type APIResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    *User  `json:"data,omitempty"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func main() {
	fmt.Println("=== JSON Struct Tags ===")

	// Create a person
	person := Person{
		Name:     "Alice",
		Age:      30,
		Email:    "alice@example.com",
		Password: "secret123", // Won't appear in JSON
		IsActive: true,
	}

	// Convert to JSON
	jsonData, _ := json.MarshalIndent(person, "", "  ")
	fmt.Println("Person as JSON:")
	fmt.Println(string(jsonData))

	fmt.Println("\n=== omitempty Tag ===")

	// omitempty: field is excluded if it has zero value
	emptyPerson := Person{
		Name: "Bob",
		Age:  25,
		// Email is empty string - will be omitted
		// CreatedAt is empty string - will be omitted
	}

	jsonData, _ = json.MarshalIndent(emptyPerson, "", "  ")
	fmt.Println("Person with empty fields:")
	fmt.Println(string(jsonData))

	fmt.Println("\n=== Parsing JSON ===")

	// Parse JSON into struct
	jsonString := `{
		"name": "Charlie",
		"age": 35,
		"email": "charlie@example.com",
		"is_active": false,
		"password": "will_be_ignored"
	}`

	var parsed Person
	json.Unmarshal([]byte(jsonString), &parsed)

	fmt.Println("Parsed person:")
	fmt.Printf("  Name: %s\n", parsed.Name)
	fmt.Printf("  Age: %d\n", parsed.Age)
	fmt.Printf("  Email: %s\n", parsed.Email)
	fmt.Printf("  Password: %q (json:\"-\" means ignored)\n", parsed.Password)

	fmt.Println("\n=== Nested Structs in JSON ===")

	// API response with nested user
	response := APIResponse{
		Status: "success",
		Code:   200,
		Data: &User{
			ID:       1,
			Username: "alice",
			Role:     "admin",
		},
	}

	jsonData, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println("API Response:")
	fmt.Println(string(jsonData))

	// Error response (no data)
	errorResponse := APIResponse{
		Status:  "error",
		Code:    404,
		Message: "User not found",
		// Data is nil - will be omitted due to omitempty
	}

	jsonData, _ = json.MarshalIndent(errorResponse, "", "  ")
	fmt.Println("\nError Response:")
	fmt.Println(string(jsonData))

	fmt.Println("\n=== Common JSON Tags ===")

	fmt.Println(`
Common struct tag patterns:

json:"name"           - Field appears as "name" in JSON
json:"name,omitempty" - Omit if zero value
json:"-"              - Never include in JSON
json:",omitempty"     - Keep field name, omit if empty

Example:
type User struct {
    ID        int    ` + "`json:\"id\"`" + `              // "ID" -> "id"
    FirstName string ` + "`json:\"first_name\"`" + `      // "FirstName" -> "first_name"
    Password  string ` + "`json:\"-\"`" + `               // Never in JSON
    Email     string ` + "`json:\"email,omitempty\"`" + ` // Omit if ""
}
`)

	fmt.Println("\n=== Working with Unknown JSON ===")

	// When structure is unknown, use map
	unknownJSON := `{"foo": "bar", "count": 42, "active": true}`
	var data map[string]interface{}
	json.Unmarshal([]byte(unknownJSON), &data)

	fmt.Println("Unknown JSON as map:")
	for key, value := range data {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
}

// TO RUN: go run day8/05_struct_tags.go
//
// KEY POINTS:
// - Struct tags are enclosed in backticks: `json:"name"`
// - Multiple tags separated by space: `json:"name" xml:"name"`
// - json:"-" excludes field from JSON
// - json:",omitempty" omits zero values
// - Tags are just strings - packages interpret them
//
// COMMON TAG PACKAGES:
// - encoding/json: JSON serialization
// - encoding/xml: XML serialization
// - database/sql: Database mapping
// - github.com/go-playground/validator: Validation
//
// EXERCISE:
// 1. Create a Product struct with json tags
// 2. Include fields: ID, Name, Price, Description (optional)
// 3. Add an InternalCode field that shouldn't be in JSON
// 4. Marshal/unmarshal a product to/from JSON
