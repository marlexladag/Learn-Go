// Day 9, Exercise 4: Struct Validation
//
// Key concepts:
// - Validate data when creating/modifying structs
// - Return errors for invalid data
// - Keep structs in valid states
// - Validation in constructors vs separate Validate() method

package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ValidationError holds multiple validation errors
type ValidationError struct {
	Errors []string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s", strings.Join(v.Errors, "; "))
}

func (v *ValidationError) Add(err string) {
	v.Errors = append(v.Errors, err)
}

func (v *ValidationError) HasErrors() bool {
	return len(v.Errors) > 0
}

// ========================================
// Pattern 1: Validation in Constructor
// ========================================

// Email wraps a string but guarantees it's valid
type Email struct {
	address string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NewEmail validates and creates an Email
func NewEmail(address string) (*Email, error) {
	address = strings.TrimSpace(address)
	address = strings.ToLower(address)

	if address == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if !emailRegex.MatchString(address) {
		return nil, fmt.Errorf("invalid email format: %s", address)
	}

	return &Email{address: address}, nil
}

func (e *Email) String() string {
	return e.address
}

// ========================================
// Pattern 2: Separate Validate Method
// ========================================

// User with a Validate() method
type User struct {
	Username string
	Email    string
	Age      int
	Password string
}

// Validate checks all fields and returns all errors
func (u *User) Validate() error {
	errs := &ValidationError{}

	// Username validation
	if u.Username == "" {
		errs.Add("username is required")
	} else if len(u.Username) < 3 {
		errs.Add("username must be at least 3 characters")
	} else if len(u.Username) > 20 {
		errs.Add("username must be at most 20 characters")
	}

	// Email validation
	if u.Email == "" {
		errs.Add("email is required")
	} else if !emailRegex.MatchString(u.Email) {
		errs.Add("email format is invalid")
	}

	// Age validation
	if u.Age < 0 {
		errs.Add("age cannot be negative")
	} else if u.Age < 13 {
		errs.Add("must be at least 13 years old")
	} else if u.Age > 150 {
		errs.Add("age seems unrealistic")
	}

	// Password validation
	if u.Password == "" {
		errs.Add("password is required")
	} else if len(u.Password) < 8 {
		errs.Add("password must be at least 8 characters")
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

// ========================================
// Pattern 3: Builder with Validation
// ========================================

// Product is immutable after creation
type Product struct {
	name        string
	price       float64
	quantity    int
	description string
}

// ProductBuilder collects data and validates on Build()
type ProductBuilder struct {
	product Product
	errors  []string
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (b *ProductBuilder) Name(name string) *ProductBuilder {
	if name == "" {
		b.errors = append(b.errors, "name is required")
	}
	b.product.name = name
	return b
}

func (b *ProductBuilder) Price(price float64) *ProductBuilder {
	if price < 0 {
		b.errors = append(b.errors, "price cannot be negative")
	}
	b.product.price = price
	return b
}

func (b *ProductBuilder) Quantity(qty int) *ProductBuilder {
	if qty < 0 {
		b.errors = append(b.errors, "quantity cannot be negative")
	}
	b.product.quantity = qty
	return b
}

func (b *ProductBuilder) Description(desc string) *ProductBuilder {
	b.product.description = desc
	return b
}

func (b *ProductBuilder) Build() (*Product, error) {
	// Additional cross-field validation
	if b.product.price == 0 && b.product.quantity > 0 {
		b.errors = append(b.errors, "cannot have inventory of free product")
	}

	if len(b.errors) > 0 {
		return nil, fmt.Errorf("validation failed: %s", strings.Join(b.errors, "; "))
	}

	return &b.product, nil
}

// Getters for Product
func (p *Product) Name() string        { return p.name }
func (p *Product) Price() float64      { return p.price }
func (p *Product) Quantity() int       { return p.quantity }
func (p *Product) Description() string { return p.description }

// ========================================
// Pattern 4: Domain-Specific Types
// ========================================

// Money represents a monetary amount (avoids float issues)
type Money struct {
	cents    int64
	currency string
}

func NewMoney(amount float64, currency string) (*Money, error) {
	if currency == "" {
		return nil, fmt.Errorf("currency is required")
	}
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency must be 3-letter code (e.g., USD)")
	}
	if amount < 0 {
		return nil, fmt.Errorf("amount cannot be negative")
	}

	// Convert to cents to avoid floating point issues
	cents := int64(amount*100 + 0.5)

	return &Money{cents: cents, currency: strings.ToUpper(currency)}, nil
}

func (m *Money) Amount() float64 {
	return float64(m.cents) / 100
}

func (m *Money) String() string {
	return fmt.Sprintf("%.2f %s", m.Amount(), m.currency)
}

// DateRange ensures start is before end
type DateRange struct {
	start time.Time
	end   time.Time
}

func NewDateRange(start, end time.Time) (*DateRange, error) {
	if start.IsZero() {
		return nil, fmt.Errorf("start date is required")
	}
	if end.IsZero() {
		return nil, fmt.Errorf("end date is required")
	}
	if end.Before(start) {
		return nil, fmt.Errorf("end date must be after start date")
	}

	return &DateRange{start: start, end: end}, nil
}

func (d *DateRange) Duration() time.Duration {
	return d.end.Sub(d.start)
}

func (d *DateRange) Contains(t time.Time) bool {
	return !t.Before(d.start) && !t.After(d.end)
}

func main() {
	fmt.Println("=== Pattern 1: Constructor Validation ===")

	// Valid email
	email, err := NewEmail("Alice@Example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid email:", email) // normalized to lowercase
	}

	// Invalid emails
	_, err = NewEmail("")
	fmt.Println("Empty email:", err)

	_, err = NewEmail("not-an-email")
	fmt.Println("Bad format:", err)

	fmt.Println("\n=== Pattern 2: Validate Method ===")

	// Valid user
	user1 := &User{
		Username: "alice",
		Email:    "alice@example.com",
		Age:      25,
		Password: "secretpass123",
	}
	if err := user1.Validate(); err != nil {
		fmt.Println("User1 errors:", err)
	} else {
		fmt.Println("User1 is valid!")
	}

	// Invalid user - multiple errors
	user2 := &User{
		Username: "ab",            // too short
		Email:    "not-an-email",  // invalid format
		Age:      10,              // too young
		Password: "short",         // too short
	}
	if err := user2.Validate(); err != nil {
		fmt.Println("User2 errors:", err)
	}

	fmt.Println("\n=== Pattern 3: Builder with Validation ===")

	// Valid product
	product, err := NewProductBuilder().
		Name("Widget").
		Price(19.99).
		Quantity(100).
		Description("A useful widget").
		Build()

	if err != nil {
		fmt.Println("Product error:", err)
	} else {
		fmt.Printf("Product: %s at $%.2f (qty: %d)\n",
			product.Name(), product.Price(), product.Quantity())
	}

	// Invalid product
	_, err = NewProductBuilder().
		Name("").           // missing name
		Price(-10).         // negative price
		Quantity(-5).       // negative quantity
		Build()
	fmt.Println("Invalid product:", err)

	fmt.Println("\n=== Pattern 4: Domain Types ===")

	// Money type
	price, _ := NewMoney(99.99, "USD")
	fmt.Println("Price:", price)

	// Date range
	start := time.Now()
	end := start.Add(7 * 24 * time.Hour)
	vacation, _ := NewDateRange(start, end)
	fmt.Printf("Vacation: %v days\n", vacation.Duration().Hours()/24)

	// Invalid date range
	_, err = NewDateRange(end, start)
	fmt.Println("Invalid range:", err)

	fmt.Println("\n=== When to Use Each Pattern ===")
	fmt.Println("1. Constructor validation: Simple types, always valid")
	fmt.Println("2. Validate() method: Complex types, multiple error reporting")
	fmt.Println("3. Builder + validation: Many optional fields, step-by-step")
	fmt.Println("4. Domain types: Replace primitives with validated types")
}

// TO RUN: go run day9/04_validation.go
//
// OUTPUT:
// === Pattern 1: Constructor Validation ===
// Valid email: alice@example.com
// ...
//
// EXERCISE:
// 1. Create a PhoneNumber type with validation (format: +1-XXX-XXX-XXXX)
// 2. Create a CreditCard struct with Number, ExpiryMonth, ExpiryYear, CVV
// 3. Add validation: 16 digits, valid expiry, 3-4 digit CVV
// 4. Return all validation errors at once
//
// KEY POINTS:
// - Validate early, fail fast
// - Return errors, don't panic
// - Consider returning all errors, not just the first
// - Domain types (Email, Money) prevent invalid values at the type level
// - Constructors guarantee valid state from the start
