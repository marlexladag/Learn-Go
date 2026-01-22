// Day 9, Exercise 2: Encapsulation in Go
//
// Key concepts:
// - Exported (public) vs unexported (private) via capitalization
// - Uppercase = exported (accessible from other packages)
// - Lowercase = unexported (package-private only)
// - Getters/setters for controlled access
// - Information hiding protects internal state

package main

import "fmt"

// BankAccount demonstrates encapsulation
// Fields are lowercase (unexported) - cannot be accessed directly from outside
type BankAccount struct {
	accountNumber string  // unexported - private
	ownerName     string  // unexported - private
	balance       float64 // unexported - private
	active        bool    // unexported - private
}

// NewBankAccount is the only way to create a valid account
func NewBankAccount(number, owner string, initialDeposit float64) *BankAccount {
	if initialDeposit < 0 {
		initialDeposit = 0
	}
	return &BankAccount{
		accountNumber: number,
		ownerName:     owner,
		balance:       initialDeposit,
		active:        true,
	}
}

// Getters - provide read access to private fields
// Go convention: no "Get" prefix for getters

// AccountNumber returns the account number (read-only)
func (a *BankAccount) AccountNumber() string {
	return a.accountNumber
}

// Owner returns the owner name
func (a *BankAccount) Owner() string {
	return a.ownerName
}

// Balance returns the current balance
func (a *BankAccount) Balance() float64 {
	return a.balance
}

// IsActive returns whether account is active
func (a *BankAccount) IsActive() bool {
	return a.active
}

// Methods that modify state with validation

// Deposit adds money with validation
func (a *BankAccount) Deposit(amount float64) error {
	if !a.active {
		return fmt.Errorf("account is closed")
	}
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw removes money with validation
func (a *BankAccount) Withdraw(amount float64) error {
	if !a.active {
		return fmt.Errorf("account is closed")
	}
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > a.balance {
		return fmt.Errorf("insufficient funds: balance is %.2f", a.balance)
	}
	a.balance -= amount
	return nil
}

// Close deactivates the account
func (a *BankAccount) Close() error {
	if !a.active {
		return fmt.Errorf("account already closed")
	}
	if a.balance > 0 {
		return fmt.Errorf("cannot close: withdraw remaining balance first")
	}
	a.active = false
	return nil
}

// String provides a formatted representation (implements fmt.Stringer)
func (a *BankAccount) String() string {
	status := "Active"
	if !a.active {
		status = "Closed"
	}
	return fmt.Sprintf("Account[%s] Owner: %s, Balance: $%.2f, Status: %s",
		a.accountNumber, a.ownerName, a.balance, status)
}

// Counter demonstrates simple encapsulation with a hidden value
type Counter struct {
	value int
}

// NewCounter creates a counter starting at 0
func NewCounter() *Counter {
	return &Counter{value: 0}
}

// Increment adds 1 to the counter
func (c *Counter) Increment() {
	c.value++
}

// Decrement subtracts 1 (but never goes below 0)
func (c *Counter) Decrement() {
	if c.value > 0 {
		c.value--
	}
}

// Value returns the current count
func (c *Counter) Value() int {
	return c.value
}

// Reset sets counter back to 0
func (c *Counter) Reset() {
	c.value = 0
}

// Temperature demonstrates encapsulation with conversion
type Temperature struct {
	celsius float64 // always stored in Celsius internally
}

// NewTemperature creates from Celsius
func NewTemperature(celsius float64) *Temperature {
	return &Temperature{celsius: celsius}
}

// NewTemperatureFromFahrenheit creates from Fahrenheit
func NewTemperatureFromFahrenheit(f float64) *Temperature {
	return &Temperature{celsius: (f - 32) * 5 / 9}
}

// Celsius returns temperature in Celsius
func (t *Temperature) Celsius() float64 {
	return t.celsius
}

// Fahrenheit returns temperature in Fahrenheit
func (t *Temperature) Fahrenheit() float64 {
	return t.celsius*9/5 + 32
}

// SetCelsius updates the temperature
func (t *Temperature) SetCelsius(c float64) {
	t.celsius = c
}

// SetFahrenheit updates using Fahrenheit
func (t *Temperature) SetFahrenheit(f float64) {
	t.celsius = (f - 32) * 5 / 9
}

func main() {
	fmt.Println("=== Bank Account Encapsulation ===")

	// Create account through constructor (the only way)
	account := NewBankAccount("1234-5678", "Alice Smith", 1000.00)
	fmt.Println(account)

	// Cannot access fields directly:
	// account.balance = 1000000 // COMPILE ERROR if in different package

	// Must use methods
	fmt.Println("\nDepositing $500...")
	account.Deposit(500)
	fmt.Printf("Balance: $%.2f\n", account.Balance())

	fmt.Println("\nWithdrawing $200...")
	account.Withdraw(200)
	fmt.Printf("Balance: $%.2f\n", account.Balance())

	// Validation prevents invalid operations
	fmt.Println("\nTrying to withdraw $5000...")
	err := account.Withdraw(5000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nTrying to deposit -$100...")
	err = account.Deposit(-100)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Counter Example ===")

	counter := NewCounter()
	fmt.Println("Initial:", counter.Value())

	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println("After 3 increments:", counter.Value())

	counter.Decrement()
	fmt.Println("After decrement:", counter.Value())

	// Can't go negative!
	counter.Reset()
	counter.Decrement()
	fmt.Println("After reset and decrement:", counter.Value()) // Still 0

	fmt.Println("\n=== Temperature Example ===")

	// Internal representation is hidden
	temp := NewTemperature(20)
	fmt.Printf("%.1f C = %.1f F\n", temp.Celsius(), temp.Fahrenheit())

	temp2 := NewTemperatureFromFahrenheit(98.6)
	fmt.Printf("Body temp: %.1f C = %.1f F\n", temp2.Celsius(), temp2.Fahrenheit())

	fmt.Println("\n=== Benefits of Encapsulation ===")
	fmt.Println("1. Data integrity - invalid states prevented")
	fmt.Println("2. Flexibility - internal representation can change")
	fmt.Println("3. Abstraction - users don't need to know internals")
	fmt.Println("4. Maintenance - changes don't break external code")
}

// TO RUN: go run day9/02_encapsulation.go
//
// OUTPUT:
// === Bank Account Encapsulation ===
// Account[1234-5678] Owner: Alice Smith, Balance: $1000.00, Status: Active
// ...
//
// EXERCISE:
// 1. Create a Password struct that stores a hashed password
// 2. Only allow setting via SetPassword(plaintext) which hashes it
// 3. Add Verify(plaintext) bool to check if password matches
// 4. The hash should never be directly accessible
//
// KEY POINTS:
// - lowercase = unexported (private to package)
// - UPPERCASE = Exported (public, accessible from other packages)
// - Getters: Balance() not GetBalance()
// - Setters with validation protect invariants
// - Encapsulation allows changing internals without breaking users
