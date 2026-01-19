// Day 6, Exercise 3: Pointer Receivers
//
// Key concepts:
// - Methods can have value receivers or pointer receivers
// - Pointer receivers can modify the struct
// - Value receivers work on a copy
// - Consistency: if one method needs pointer, use pointers for all

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Value Receiver vs Pointer Receiver ===")

	c1 := Counter{value: 0}

	// Value receiver - original NOT modified
	c1.IncrementValue()
	fmt.Println("After IncrementValue:", c1.value) // Still 0!

	// Pointer receiver - original IS modified
	c1.IncrementPointer()
	fmt.Println("After IncrementPointer:", c1.value) // Now 1!

	fmt.Println("\n=== Automatic Dereferencing ===")

	c2 := &Counter{value: 10}

	// Go automatically dereferences for method calls
	c2.IncrementPointer() // Works even though c2 is already a pointer
	fmt.Println("Pointer incremented:", c2.value)

	c3 := Counter{value: 20}
	c3.IncrementPointer() // Go automatically takes address
	fmt.Println("Value incremented via pointer method:", c3.value)

	fmt.Println("\n=== When to Use Pointer Receivers ===")

	// 1. When the method needs to modify the receiver
	account := BankAccount{Balance: 100}
	account.Deposit(50)
	account.Withdraw(30)
	fmt.Println("Account balance:", account.Balance)

	// 2. When the struct is large (avoid copying)
	big := LargeStruct{}
	big.Process() // Would be expensive to copy

	// 3. For consistency with other methods
	point := Point{X: 3, Y: 4}
	fmt.Println("Distance:", point.Distance())
	point.Scale(2)
	fmt.Println("After scale:", point)

	fmt.Println("\n=== Pointer Receiver with Nil ===")

	var nilCounter *Counter
	// This would panic with value access, but we can handle nil
	fmt.Println("Nil counter safe value:", nilCounter.SafeValue())

	fmt.Println("\n=== Practical Example: Stack ===")

	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack:", stack.items)

	val, ok := stack.Pop()
	fmt.Printf("Popped: %d, ok: %v\n", val, ok)
	fmt.Println("Stack after pop:", stack.items)

	fmt.Println("Peek:", stack.Peek())
	fmt.Println("Size:", stack.Size())
	fmt.Println("IsEmpty:", stack.IsEmpty())

	fmt.Println("\n=== Method Chaining with Pointers ===")

	builder := &StringBuilder{}
	result := builder.
		Append("Hello").
		Append(" ").
		Append("World").
		Append("!").
		String()
	fmt.Println("Built string:", result)

	fmt.Println("\n=== Receiver Type Guidelines ===")

	/*
		Use POINTER receiver when:
		1. Method modifies the receiver
		2. Struct is large (avoid copying)
		3. Consistency with other methods that need pointers
		4. The receiver might be nil and you need to handle it

		Use VALUE receiver when:
		1. Method doesn't modify the receiver
		2. Struct is small (few fields of basic types)
		3. The type is immutable by design
		4. It's a map, func, or chan (already reference types)
	*/

	// Example: time.Time uses value receivers (immutable design)
	// Example: bytes.Buffer uses pointer receivers (needs modification)
}

// Counter demonstrates value vs pointer receivers
type Counter struct {
	value int
}

// IncrementValue uses value receiver - gets a copy
func (c Counter) IncrementValue() {
	c.value++ // Only modifies the copy!
}

// IncrementPointer uses pointer receiver - modifies original
func (c *Counter) IncrementPointer() {
	c.value++ // Modifies the actual counter
}

// SafeValue handles nil receiver gracefully
func (c *Counter) SafeValue() int {
	if c == nil {
		return 0
	}
	return c.value
}

// BankAccount demonstrates state modification
type BankAccount struct {
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}
	a.Balance -= amount
	return true
}

// LargeStruct demonstrates avoiding copies
type LargeStruct struct {
	Data [1000]int
	// Other large fields...
}

func (l *LargeStruct) Process() {
	// Uses pointer to avoid copying 1000 ints
	fmt.Println("Processing large struct efficiently")
}

// Point demonstrates consistency in receivers
type Point struct {
	X, Y float64
}

// Distance could use value receiver (doesn't modify)
// but we use pointer for consistency with Scale
func (p *Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Stack demonstrates pointer receivers for data structures
type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := len(s.items) - 1
	val := s.items[last]
	s.items = s.items[:last]
	return val, true
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// StringBuilder demonstrates method chaining
type StringBuilder struct {
	data []byte
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.data = append(sb.data, s...)
	return sb // Return pointer for chaining
}

func (sb *StringBuilder) String() string {
	return string(sb.data)
}

// TO RUN: go run day6/03_pointer_receivers.go
//
// OUTPUT:
// === Value Receiver vs Pointer Receiver ===
// After IncrementValue: 0
// After IncrementPointer: 1
// ...
//
// EXERCISE:
// 1. Create a Rectangle struct with Width and Height
// 2. Add Area() method with value receiver (doesn't modify)
// 3. Add Scale() method with pointer receiver (modifies dimensions)
// 4. Add method chaining: SetWidth() and SetHeight() returning *Rectangle
//
// KEY POINTS:
// - Pointer receivers modify the original struct
// - Value receivers work on a copy
// - Go auto-dereferences when calling methods
// - Be consistent: if one method needs pointer, use for all
// - Return *T for method chaining
