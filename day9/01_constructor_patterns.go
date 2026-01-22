// Day 9, Exercise 1: Constructor Patterns
//
// Key concepts:
// - Go doesn't have constructors, but has patterns for initialization
// - New functions return pointers (NewType pattern)
// - Factory functions can return interfaces or handle complex setup
// - Option functions for flexible configuration

package main

import (
	"fmt"
	"time"
)

// User represents a user account
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
	Active    bool
}

// NewUser is a constructor function (returns a pointer)
// This is the most common pattern in Go
func NewUser(id int, username, email string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		Active:    true,
	}
}

// Server represents a server configuration
type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
	TLS     bool
}

// NewServer with default values
func NewServer(host string) *Server {
	return &Server{
		Host:    host,
		Port:    8080,         // sensible default
		Timeout: time.Second * 30,
		MaxConn: 100,
		TLS:     false,
	}
}

// ServerOption is a function type for configuration
type ServerOption func(*Server)

// WithPort sets the port
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

// WithTimeout sets the timeout
func WithTimeout(t time.Duration) ServerOption {
	return func(s *Server) {
		s.Timeout = t
	}
}

// WithTLS enables TLS
func WithTLS() ServerOption {
	return func(s *Server) {
		s.TLS = true
	}
}

// WithMaxConnections sets max connections
func WithMaxConnections(n int) ServerOption {
	return func(s *Server) {
		s.MaxConn = n
	}
}

// NewServerWithOptions uses the functional options pattern
func NewServerWithOptions(host string, opts ...ServerOption) *Server {
	// Start with defaults
	s := &Server{
		Host:    host,
		Port:    8080,
		Timeout: time.Second * 30,
		MaxConn: 100,
		TLS:     false,
	}

	// Apply all options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Database connection example with required and optional params
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

// NewDBConfig with required parameters
func NewDBConfig(host, user, password, database string) *DBConfig {
	return &DBConfig{
		Host:     host,
		Port:     5432, // default PostgreSQL port
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  "disable",
	}
}

func main() {
	fmt.Println("=== Basic Constructor Pattern ===")

	// Using constructor function
	user := NewUser(1, "alice", "alice@example.com")
	fmt.Printf("User: %+v\n", user)

	// Compare with direct initialization
	user2 := &User{
		ID:       2,
		Username: "bob",
		Email:    "bob@example.com",
		// Oops! Forgot CreatedAt and Active
	}
	fmt.Printf("User2 (incomplete): %+v\n", user2)

	fmt.Println("\n=== Constructor with Defaults ===")

	// Server with defaults
	srv := NewServer("localhost")
	fmt.Printf("Default server: %+v\n", srv)

	fmt.Println("\n=== Functional Options Pattern ===")

	// No options - use all defaults
	srv1 := NewServerWithOptions("api.example.com")
	fmt.Printf("Default options: %+v\n", srv1)

	// Customize just what you need
	srv2 := NewServerWithOptions("api.example.com",
		WithPort(443),
		WithTLS(),
	)
	fmt.Printf("With TLS: %+v\n", srv2)

	// Full customization
	srv3 := NewServerWithOptions("api.example.com",
		WithPort(8443),
		WithTimeout(time.Minute),
		WithMaxConnections(1000),
		WithTLS(),
	)
	fmt.Printf("Fully configured: %+v\n", srv3)

	fmt.Println("\n=== When to Use Each Pattern ===")

	// Pattern 1: Simple constructor (most common)
	// Use when: few fields, clear required params
	fmt.Println("1. Simple constructor: NewUser(id, name, email)")

	// Pattern 2: Constructor with defaults
	// Use when: many fields, but most have sensible defaults
	fmt.Println("2. With defaults: NewServer(host) // other fields default")

	// Pattern 3: Functional options
	// Use when: many optional config options, flexibility needed
	fmt.Println("3. Functional options: NewServer(host, WithPort(443), WithTLS())")

	fmt.Println("\n=== Database Config Example ===")

	// Required params enforced, sensible defaults for rest
	db := NewDBConfig("localhost", "admin", "secret", "myapp")
	fmt.Printf("DB Config: %+v\n", db)
}

// TO RUN: go run day9/01_constructor_patterns.go
//
// OUTPUT:
// === Basic Constructor Pattern ===
// User: &{ID:1 Username:alice Email:alice@example.com CreatedAt:... Active:true}
// User2 (incomplete): &{ID:2 Username:bob Email:bob@example.com CreatedAt:0001-01-01... Active:false}
// ...
//
// EXERCISE:
// 1. Create a Logger struct with Level, Output, Prefix, Timestamp fields
// 2. Create NewLogger() with sensible defaults
// 3. Create functional options: WithLevel(), WithPrefix(), WithTimestamp()
// 4. Create a logger using the functional options pattern
//
// KEY POINTS:
// - Go uses function constructors instead of special syntax
// - NewType() conventionally returns *Type
// - Functional options provide flexible, readable configuration
// - Constructors ensure required fields are set and defaults applied
