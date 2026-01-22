// Day 9, Exercise 5: Builder Pattern
//
// Key concepts:
// - Builder separates construction from representation
// - Fluent interface with method chaining
// - Handle complex object creation with many parameters
// - Useful when struct has many optional fields

package main

import (
	"fmt"
	"strings"
)

// ========================================
// Example 1: HTTP Request Builder
// ========================================

// HTTPRequest represents an HTTP request
type HTTPRequest struct {
	method  string
	url     string
	headers map[string]string
	body    string
	timeout int
}

// HTTPRequestBuilder builds HTTP requests
type HTTPRequestBuilder struct {
	request HTTPRequest
}

// NewHTTPRequestBuilder creates a new builder
func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		request: HTTPRequest{
			method:  "GET",
			headers: make(map[string]string),
			timeout: 30,
		},
	}
}

// Method sets the HTTP method
func (b *HTTPRequestBuilder) Method(method string) *HTTPRequestBuilder {
	b.request.method = method
	return b
}

// URL sets the URL
func (b *HTTPRequestBuilder) URL(url string) *HTTPRequestBuilder {
	b.request.url = url
	return b
}

// Header adds a header
func (b *HTTPRequestBuilder) Header(key, value string) *HTTPRequestBuilder {
	b.request.headers[key] = value
	return b
}

// Body sets the request body
func (b *HTTPRequestBuilder) Body(body string) *HTTPRequestBuilder {
	b.request.body = body
	return b
}

// Timeout sets the timeout in seconds
func (b *HTTPRequestBuilder) Timeout(seconds int) *HTTPRequestBuilder {
	b.request.timeout = seconds
	return b
}

// Build returns the constructed request
func (b *HTTPRequestBuilder) Build() HTTPRequest {
	return b.request
}

// Convenience methods
func (b *HTTPRequestBuilder) GET(url string) *HTTPRequestBuilder {
	return b.Method("GET").URL(url)
}

func (b *HTTPRequestBuilder) POST(url string) *HTTPRequestBuilder {
	return b.Method("POST").URL(url)
}

func (b *HTTPRequestBuilder) JSON(body string) *HTTPRequestBuilder {
	return b.Header("Content-Type", "application/json").Body(body)
}

func (r HTTPRequest) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %s\n", r.method, r.url))
	for k, v := range r.headers {
		sb.WriteString(fmt.Sprintf("  %s: %s\n", k, v))
	}
	if r.body != "" {
		sb.WriteString(fmt.Sprintf("  Body: %s\n", r.body))
	}
	sb.WriteString(fmt.Sprintf("  Timeout: %ds", r.timeout))
	return sb.String()
}

// ========================================
// Example 2: SQL Query Builder
// ========================================

// SQLQuery represents a SQL SELECT query
type SQLQuery struct {
	table      string
	columns    []string
	conditions []string
	orderBy    string
	limit      int
}

// QueryBuilder builds SQL queries
type QueryBuilder struct {
	query SQLQuery
}

// Select starts a new query
func Select(columns ...string) *QueryBuilder {
	return &QueryBuilder{
		query: SQLQuery{
			columns: columns,
		},
	}
}

// From sets the table
func (b *QueryBuilder) From(table string) *QueryBuilder {
	b.query.table = table
	return b
}

// Where adds a condition
func (b *QueryBuilder) Where(condition string) *QueryBuilder {
	b.query.conditions = append(b.query.conditions, condition)
	return b
}

// OrderBy sets the order
func (b *QueryBuilder) OrderBy(column string) *QueryBuilder {
	b.query.orderBy = column
	return b
}

// Limit sets the result limit
func (b *QueryBuilder) Limit(n int) *QueryBuilder {
	b.query.limit = n
	return b
}

// Build generates the SQL string
func (b *QueryBuilder) Build() string {
	var sb strings.Builder

	// SELECT
	if len(b.query.columns) == 0 {
		sb.WriteString("SELECT *")
	} else {
		sb.WriteString("SELECT ")
		sb.WriteString(strings.Join(b.query.columns, ", "))
	}

	// FROM
	sb.WriteString(" FROM ")
	sb.WriteString(b.query.table)

	// WHERE
	if len(b.query.conditions) > 0 {
		sb.WriteString(" WHERE ")
		sb.WriteString(strings.Join(b.query.conditions, " AND "))
	}

	// ORDER BY
	if b.query.orderBy != "" {
		sb.WriteString(" ORDER BY ")
		sb.WriteString(b.query.orderBy)
	}

	// LIMIT
	if b.query.limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", b.query.limit))
	}

	return sb.String()
}

// ========================================
// Example 3: Email Builder
// ========================================

// Email represents an email message
type Email struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	isHTML      bool
	attachments []string
}

// EmailBuilder builds emails
type EmailBuilder struct {
	email Email
}

// NewEmailBuilder creates a builder
func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{
		email: Email{
			to:          []string{},
			cc:          []string{},
			bcc:         []string{},
			attachments: []string{},
		},
	}
}

func (b *EmailBuilder) From(addr string) *EmailBuilder {
	b.email.from = addr
	return b
}

func (b *EmailBuilder) To(addrs ...string) *EmailBuilder {
	b.email.to = append(b.email.to, addrs...)
	return b
}

func (b *EmailBuilder) CC(addrs ...string) *EmailBuilder {
	b.email.cc = append(b.email.cc, addrs...)
	return b
}

func (b *EmailBuilder) BCC(addrs ...string) *EmailBuilder {
	b.email.bcc = append(b.email.bcc, addrs...)
	return b
}

func (b *EmailBuilder) Subject(s string) *EmailBuilder {
	b.email.subject = s
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func (b *EmailBuilder) HTML(html string) *EmailBuilder {
	b.email.body = html
	b.email.isHTML = true
	return b
}

func (b *EmailBuilder) Attach(files ...string) *EmailBuilder {
	b.email.attachments = append(b.email.attachments, files...)
	return b
}

func (b *EmailBuilder) Build() Email {
	return b.email
}

func (e Email) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("From: %s\n", e.from))
	sb.WriteString(fmt.Sprintf("To: %s\n", strings.Join(e.to, ", ")))
	if len(e.cc) > 0 {
		sb.WriteString(fmt.Sprintf("CC: %s\n", strings.Join(e.cc, ", ")))
	}
	if len(e.bcc) > 0 {
		sb.WriteString(fmt.Sprintf("BCC: %s\n", strings.Join(e.bcc, ", ")))
	}
	sb.WriteString(fmt.Sprintf("Subject: %s\n", e.subject))
	contentType := "text/plain"
	if e.isHTML {
		contentType = "text/html"
	}
	sb.WriteString(fmt.Sprintf("Content-Type: %s\n", contentType))
	if len(e.attachments) > 0 {
		sb.WriteString(fmt.Sprintf("Attachments: %s\n", strings.Join(e.attachments, ", ")))
	}
	sb.WriteString(fmt.Sprintf("Body: %s", e.body))
	return sb.String()
}

// ========================================
// Example 4: Pizza Builder (Classic Example)
// ========================================

// Pizza with many optional toppings
type Pizza struct {
	size     string
	crust    string
	sauce    string
	cheese   string
	toppings []string
}

type PizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder(size string) *PizzaBuilder {
	return &PizzaBuilder{
		pizza: Pizza{
			size:     size,
			crust:    "regular",
			sauce:    "tomato",
			cheese:   "mozzarella",
			toppings: []string{},
		},
	}
}

func (b *PizzaBuilder) Crust(crust string) *PizzaBuilder {
	b.pizza.crust = crust
	return b
}

func (b *PizzaBuilder) Sauce(sauce string) *PizzaBuilder {
	b.pizza.sauce = sauce
	return b
}

func (b *PizzaBuilder) Cheese(cheese string) *PizzaBuilder {
	b.pizza.cheese = cheese
	return b
}

func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.pizza.toppings = append(b.pizza.toppings, topping)
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

func (p Pizza) String() string {
	toppings := "none"
	if len(p.toppings) > 0 {
		toppings = strings.Join(p.toppings, ", ")
	}
	return fmt.Sprintf("%s pizza: %s crust, %s sauce, %s cheese, toppings: %s",
		p.size, p.crust, p.sauce, p.cheese, toppings)
}

func main() {
	fmt.Println("=== HTTP Request Builder ===")

	// Simple GET request
	req1 := NewHTTPRequestBuilder().
		GET("https://api.example.com/users").
		Header("Authorization", "Bearer token123").
		Build()
	fmt.Println(req1)

	fmt.Println()

	// POST with JSON body
	req2 := NewHTTPRequestBuilder().
		POST("https://api.example.com/users").
		JSON(`{"name": "Alice", "email": "alice@example.com"}`).
		Header("X-Request-ID", "abc123").
		Timeout(60).
		Build()
	fmt.Println(req2)

	fmt.Println("\n=== SQL Query Builder ===")

	// Simple query
	query1 := Select("id", "name", "email").
		From("users").
		Build()
	fmt.Println(query1)

	// Complex query
	query2 := Select("*").
		From("orders").
		Where("status = 'pending'").
		Where("amount > 100").
		OrderBy("created_at DESC").
		Limit(10).
		Build()
	fmt.Println(query2)

	fmt.Println("\n=== Email Builder ===")

	email := NewEmailBuilder().
		From("sender@example.com").
		To("user1@example.com", "user2@example.com").
		CC("manager@example.com").
		Subject("Weekly Report").
		Body("Please find the weekly report attached.").
		Attach("report.pdf", "data.xlsx").
		Build()
	fmt.Println(email)

	fmt.Println("\n=== Pizza Builder ===")

	// Default pizza
	pizza1 := NewPizzaBuilder("medium").Build()
	fmt.Println(pizza1)

	// Custom pizza
	pizza2 := NewPizzaBuilder("large").
		Crust("thin").
		Sauce("bbq").
		Cheese("cheddar").
		AddTopping("pepperoni").
		AddTopping("mushrooms").
		AddTopping("olives").
		Build()
	fmt.Println(pizza2)

	fmt.Println("\n=== When to Use Builder Pattern ===")
	fmt.Println("1. Many constructor parameters (>4)")
	fmt.Println("2. Many optional parameters")
	fmt.Println("3. Object needs step-by-step construction")
	fmt.Println("4. Same construction process, different representations")
	fmt.Println("5. Readable, fluent API desired")
}

// TO RUN: go run day9/05_builder_pattern.go
//
// OUTPUT:
// === HTTP Request Builder ===
// GET https://api.example.com/users
//   Authorization: Bearer token123
//   Timeout: 30s
// ...
//
// EXERCISE:
// 1. Create a ConfigBuilder for an application config
// 2. Include: host, port, debug mode, log level, database URL, cache size
// 3. Provide sensible defaults
// 4. Add validation in Build() method
// 5. Add convenience methods like Production() and Development()
//
// KEY POINTS:
// - Builder methods return *Builder for chaining
// - Build() returns the constructed object
// - Each method sets one aspect of the object
// - Provides clean, readable construction syntax
// - Can include validation in Build()
