// Day 9, Challenge: Build an E-Commerce Order System
//
// Apply all patterns learned today to build an order management system.
//
// Requirements:
// 1. Use constructor patterns for creating valid objects
// 2. Encapsulate internal state with getters/setters
// 3. Use composition to build complex types
// 4. Validate all inputs
// 5. Use builder pattern for complex order creation

package main

import (
	"fmt"
	"strings"
	"time"
)

// ========================================
// Value Objects (validated on creation)
// ========================================

// Money represents a monetary value
type Money struct {
	cents    int64
	currency string
}

func NewMoney(amount float64, currency string) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("amount cannot be negative")
	}
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency must be 3-letter code")
	}
	return &Money{
		cents:    int64(amount*100 + 0.5),
		currency: strings.ToUpper(currency),
	}, nil
}

func (m *Money) Amount() float64   { return float64(m.cents) / 100 }
func (m *Money) Currency() string  { return m.currency }
func (m *Money) Cents() int64      { return m.cents }
func (m *Money) String() string    { return fmt.Sprintf("%.2f %s", m.Amount(), m.currency) }

func (m *Money) Add(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, fmt.Errorf("cannot add different currencies")
	}
	return &Money{cents: m.cents + other.cents, currency: m.currency}, nil
}

func (m *Money) Multiply(qty int) *Money {
	return &Money{cents: m.cents * int64(qty), currency: m.currency}
}

// Email is a validated email address
type Email struct {
	address string
}

func NewEmail(addr string) (*Email, error) {
	addr = strings.TrimSpace(strings.ToLower(addr))
	if !strings.Contains(addr, "@") || !strings.Contains(addr, ".") {
		return nil, fmt.Errorf("invalid email format")
	}
	return &Email{address: addr}, nil
}

func (e *Email) String() string { return e.address }

// ========================================
// Domain Objects (encapsulated)
// ========================================

// Address with encapsulation
type Address struct {
	street  string
	city    string
	state   string
	zipCode string
	country string
}

func NewAddress(street, city, state, zipCode, country string) (*Address, error) {
	if street == "" || city == "" || country == "" {
		return nil, fmt.Errorf("street, city, and country are required")
	}
	return &Address{
		street:  street,
		city:    city,
		state:   state,
		zipCode: zipCode,
		country: country,
	}, nil
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s, %s", a.street, a.city, a.state, a.zipCode, a.country)
}

// Customer with encapsulated fields
type Customer struct {
	id        string
	name      string
	email     *Email
	addresses []*Address
	createdAt time.Time
}

func NewCustomer(id, name string, email *Email) *Customer {
	return &Customer{
		id:        id,
		name:      name,
		email:     email,
		addresses: []*Address{},
		createdAt: time.Now(),
	}
}

func (c *Customer) ID() string         { return c.id }
func (c *Customer) Name() string       { return c.name }
func (c *Customer) Email() string      { return c.email.String() }
func (c *Customer) CreatedAt() time.Time { return c.createdAt }

func (c *Customer) AddAddress(addr *Address) {
	c.addresses = append(c.addresses, addr)
}

func (c *Customer) Addresses() []*Address {
	// Return a copy to prevent external modification
	result := make([]*Address, len(c.addresses))
	copy(result, c.addresses)
	return result
}

// Product with encapsulation
type Product struct {
	sku         string
	name        string
	description string
	price       *Money
	inStock     int
}

func NewProduct(sku, name string, price *Money) *Product {
	return &Product{
		sku:     sku,
		name:    name,
		price:   price,
		inStock: 0,
	}
}

func (p *Product) SKU() string         { return p.sku }
func (p *Product) Name() string        { return p.name }
func (p *Product) Price() *Money       { return p.price }
func (p *Product) InStock() int        { return p.inStock }
func (p *Product) Description() string { return p.description }

func (p *Product) SetDescription(desc string) { p.description = desc }
func (p *Product) AddStock(qty int)           { p.inStock += qty }

func (p *Product) RemoveStock(qty int) error {
	if qty > p.inStock {
		return fmt.Errorf("insufficient stock: have %d, need %d", p.inStock, qty)
	}
	p.inStock -= qty
	return nil
}

// ========================================
// Order (composition + builder)
// ========================================

// OrderItem composes Product with quantity
type OrderItem struct {
	product  *Product
	quantity int
	subtotal *Money
}

func NewOrderItem(product *Product, quantity int) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, fmt.Errorf("quantity must be positive")
	}
	return &OrderItem{
		product:  product,
		quantity: quantity,
		subtotal: product.Price().Multiply(quantity),
	}, nil
}

func (i *OrderItem) Product() *Product { return i.product }
func (i *OrderItem) Quantity() int     { return i.quantity }
func (i *OrderItem) Subtotal() *Money  { return i.subtotal }

// OrderStatus represents order lifecycle
type OrderStatus string

const (
	StatusPending    OrderStatus = "pending"
	StatusConfirmed  OrderStatus = "confirmed"
	StatusShipped    OrderStatus = "shipped"
	StatusDelivered  OrderStatus = "delivered"
	StatusCancelled  OrderStatus = "cancelled"
)

// Order composes multiple types
type Order struct {
	id              string
	customer        *Customer
	items           []*OrderItem
	shippingAddress *Address
	billingAddress  *Address
	status          OrderStatus
	total           *Money
	createdAt       time.Time
	updatedAt       time.Time
	notes           string
}

// Getters
func (o *Order) ID() string              { return o.id }
func (o *Order) Customer() *Customer     { return o.customer }
func (o *Order) Status() OrderStatus     { return o.status }
func (o *Order) Total() *Money           { return o.total }
func (o *Order) CreatedAt() time.Time    { return o.createdAt }
func (o *Order) Notes() string           { return o.notes }
func (o *Order) ShippingAddress() *Address { return o.shippingAddress }

func (o *Order) Items() []*OrderItem {
	result := make([]*OrderItem, len(o.items))
	copy(result, o.items)
	return result
}

// Status transitions with validation
func (o *Order) Confirm() error {
	if o.status != StatusPending {
		return fmt.Errorf("can only confirm pending orders")
	}
	o.status = StatusConfirmed
	o.updatedAt = time.Now()
	return nil
}

func (o *Order) Ship() error {
	if o.status != StatusConfirmed {
		return fmt.Errorf("can only ship confirmed orders")
	}
	o.status = StatusShipped
	o.updatedAt = time.Now()
	return nil
}

func (o *Order) Deliver() error {
	if o.status != StatusShipped {
		return fmt.Errorf("can only deliver shipped orders")
	}
	o.status = StatusDelivered
	o.updatedAt = time.Now()
	return nil
}

func (o *Order) Cancel() error {
	if o.status == StatusDelivered {
		return fmt.Errorf("cannot cancel delivered orders")
	}
	if o.status == StatusCancelled {
		return fmt.Errorf("order already cancelled")
	}
	o.status = StatusCancelled
	o.updatedAt = time.Now()
	return nil
}

func (o *Order) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Order #%s\n", o.id))
	sb.WriteString(fmt.Sprintf("Customer: %s (%s)\n", o.customer.Name(), o.customer.Email()))
	sb.WriteString(fmt.Sprintf("Status: %s\n", o.status))
	sb.WriteString("Items:\n")
	for _, item := range o.items {
		sb.WriteString(fmt.Sprintf("  - %s x%d = %s\n",
			item.Product().Name(), item.Quantity(), item.Subtotal()))
	}
	sb.WriteString(fmt.Sprintf("Total: %s\n", o.total))
	sb.WriteString(fmt.Sprintf("Ship to: %s\n", o.shippingAddress))
	if o.notes != "" {
		sb.WriteString(fmt.Sprintf("Notes: %s\n", o.notes))
	}
	return sb.String()
}

// ========================================
// Order Builder
// ========================================

type OrderBuilder struct {
	order  Order
	errors []string
}

func NewOrderBuilder(id string) *OrderBuilder {
	return &OrderBuilder{
		order: Order{
			id:        id,
			items:     []*OrderItem{},
			status:    StatusPending,
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
	}
}

func (b *OrderBuilder) Customer(c *Customer) *OrderBuilder {
	if c == nil {
		b.errors = append(b.errors, "customer is required")
	}
	b.order.customer = c
	return b
}

func (b *OrderBuilder) AddItem(product *Product, quantity int) *OrderBuilder {
	item, err := NewOrderItem(product, quantity)
	if err != nil {
		b.errors = append(b.errors, err.Error())
		return b
	}
	b.order.items = append(b.order.items, item)
	return b
}

func (b *OrderBuilder) ShipTo(addr *Address) *OrderBuilder {
	if addr == nil {
		b.errors = append(b.errors, "shipping address is required")
	}
	b.order.shippingAddress = addr
	return b
}

func (b *OrderBuilder) BillTo(addr *Address) *OrderBuilder {
	b.order.billingAddress = addr
	return b
}

func (b *OrderBuilder) Notes(notes string) *OrderBuilder {
	b.order.notes = notes
	return b
}

func (b *OrderBuilder) Build() (*Order, error) {
	// Validate required fields
	if b.order.customer == nil {
		b.errors = append(b.errors, "customer is required")
	}
	if len(b.order.items) == 0 {
		b.errors = append(b.errors, "order must have at least one item")
	}
	if b.order.shippingAddress == nil {
		b.errors = append(b.errors, "shipping address is required")
	}

	if len(b.errors) > 0 {
		return nil, fmt.Errorf("order validation failed: %s", strings.Join(b.errors, "; "))
	}

	// Calculate total
	var totalCents int64
	currency := ""
	for _, item := range b.order.items {
		if currency == "" {
			currency = item.Subtotal().Currency()
		}
		totalCents += item.Subtotal().Cents()
	}
	b.order.total = &Money{cents: totalCents, currency: currency}

	// Use shipping as billing if not specified
	if b.order.billingAddress == nil {
		b.order.billingAddress = b.order.shippingAddress
	}

	return &b.order, nil
}

// ========================================
// Main - Demo Everything
// ========================================

func main() {
	fmt.Println("=== E-Commerce Order System ===\n")

	// Create products
	fmt.Println("Creating products...")
	laptopPrice, _ := NewMoney(999.99, "USD")
	laptop := NewProduct("SKU-001", "Gaming Laptop", laptopPrice)
	laptop.SetDescription("High-performance gaming laptop")
	laptop.AddStock(50)

	mousePrice, _ := NewMoney(49.99, "USD")
	mouse := NewProduct("SKU-002", "Wireless Mouse", mousePrice)
	mouse.AddStock(200)

	keyboardPrice, _ := NewMoney(129.99, "USD")
	keyboard := NewProduct("SKU-003", "Mechanical Keyboard", keyboardPrice)
	keyboard.AddStock(100)

	fmt.Printf("  %s: %s (%d in stock)\n", laptop.SKU(), laptop.Name(), laptop.InStock())
	fmt.Printf("  %s: %s (%d in stock)\n", mouse.SKU(), mouse.Name(), mouse.InStock())
	fmt.Printf("  %s: %s (%d in stock)\n", keyboard.SKU(), keyboard.Name(), keyboard.InStock())

	// Create customer
	fmt.Println("\nCreating customer...")
	email, _ := NewEmail("alice@example.com")
	customer := NewCustomer("CUST-001", "Alice Johnson", email)

	homeAddr, _ := NewAddress("123 Main St", "Boston", "MA", "02101", "USA")
	customer.AddAddress(homeAddr)

	workAddr, _ := NewAddress("456 Office Park", "Cambridge", "MA", "02139", "USA")
	customer.AddAddress(workAddr)

	fmt.Printf("  Customer: %s (%s)\n", customer.Name(), customer.Email())
	fmt.Printf("  Addresses: %d registered\n", len(customer.Addresses()))

	// Build an order using the builder
	fmt.Println("\nBuilding order...")
	order, err := NewOrderBuilder("ORD-2024-001").
		Customer(customer).
		AddItem(laptop, 1).
		AddItem(mouse, 2).
		AddItem(keyboard, 1).
		ShipTo(homeAddr).
		Notes("Please leave at front door").
		Build()

	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	fmt.Println("\n" + order.String())

	// Process order through lifecycle
	fmt.Println("=== Order Lifecycle ===")

	fmt.Printf("Current status: %s\n", order.Status())

	fmt.Println("\nConfirming order...")
	if err := order.Confirm(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Status: %s\n", order.Status())
	}

	fmt.Println("\nShipping order...")
	if err := order.Ship(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Status: %s\n", order.Status())
	}

	fmt.Println("\nDelivering order...")
	if err := order.Deliver(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Status: %s\n", order.Status())
	}

	// Try invalid transition
	fmt.Println("\nTrying to cancel delivered order...")
	if err := order.Cancel(); err != nil {
		fmt.Println("Error:", err)
	}

	// Try building invalid order
	fmt.Println("\n=== Validation Demo ===")

	_, err = NewOrderBuilder("ORD-BAD").
		Build() // Missing customer, items, address

	if err != nil {
		fmt.Println("Invalid order rejected:", err)
	}

	fmt.Println("\n=== Concepts Applied ===")
	fmt.Println("1. Constructor patterns: NewMoney, NewEmail, NewAddress, NewCustomer")
	fmt.Println("2. Encapsulation: Private fields with getter/setter methods")
	fmt.Println("3. Composition: Order contains Customer, Items, Address")
	fmt.Println("4. Validation: Checked at creation and state transitions")
	fmt.Println("5. Builder pattern: OrderBuilder for complex order creation")
}

// TO RUN: go run day9/06_challenge.go
//
// OUTPUT:
// === E-Commerce Order System ===
// Creating products...
// ...
//
// EXTENSIONS TO TRY:
// 1. Add discount/coupon support to orders
// 2. Implement inventory management (reserve stock on order)
// 3. Add payment processing with multiple payment methods
// 4. Create an order history for customers
// 5. Add email notifications for status changes
//
// PATTERNS USED:
// - Value Objects: Money, Email (immutable, validated)
// - Entity: Customer, Product, Order (identity, mutable)
// - Builder: OrderBuilder (complex construction)
// - Encapsulation: All structs hide internal state
// - Composition: Order composes Customer, Items, Address
// - State Machine: Order status transitions
