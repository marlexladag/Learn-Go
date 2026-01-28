package main

import "fmt"

// ============================================================================
// DAY 10: INTERFACES IN GO
// File 4: The Empty Interface
// ============================================================================
//
// THE EMPTY INTERFACE: interface{}
//
// An empty interface has zero methods. Since every type has at least
// zero methods, EVERY type implements the empty interface!
//
// This makes interface{} (or 'any' in Go 1.18+) a way to hold any value.
//
// Common uses:
//   - Generic containers before Go 1.18 generics
//   - JSON parsing (map[string]interface{})
//   - Printf-style variadic functions
//   - Plugin systems
//
// WARNING: Empty interfaces sacrifice type safety. Use sparingly!
//
// ============================================================================

// Container can hold any type
type Container struct {
	items []interface{}
}

func NewContainer() *Container {
	return &Container{items: make([]interface{}, 0)}
}

func (c *Container) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *Container) Get(index int) interface{} {
	if index < 0 || index >= len(c.items) {
		return nil
	}
	return c.items[index]
}

func (c *Container) Size() int {
	return len(c.items)
}

// ============================================================================
// 'any' IS AN ALIAS FOR interface{} (Go 1.18+)
// ============================================================================
//
// In modern Go, you can use 'any' instead of 'interface{}':
//   var x any = "hello"     // same as interface{}
//   var y interface{} = 42  // same as any
//
// 'any' is more readable and preferred in new code.
//
// ============================================================================

// PrintAnything accepts any type
func PrintAnything(value any) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// AcceptMultiple is like fmt.Println - accepts any number of any type
func AcceptMultiple(values ...any) {
	for i, v := range values {
		fmt.Printf("  [%d] %v (%T)\n", i, v, v)
	}
}

// ============================================================================
// EMPTY INTERFACE IN MAPS (JSON-like data)
// ============================================================================

// Simulate JSON-like dynamic data
func createDynamicData() map[string]any {
	return map[string]any{
		"name":     "Alice",
		"age":      30,
		"active":   true,
		"balance":  1234.56,
		"tags":     []string{"admin", "user"},
		"metadata": map[string]any{"level": 5, "premium": true},
	}
}

// ============================================================================
// RETRIEVING VALUES FROM EMPTY INTERFACE
// ============================================================================
//
// When you have an interface{}, you need to extract the actual value.
// This requires TYPE ASSERTION (covered in detail in file 05).
//
// Basic syntax:
//   value := x.(Type)        // panics if wrong type
//   value, ok := x.(Type)    // safe - ok is false if wrong type
//
// ============================================================================

// SafeGetString safely extracts a string from any
func SafeGetString(value any) (string, bool) {
	str, ok := value.(string)
	return str, ok
}

// SafeGetInt safely extracts an int from any
func SafeGetInt(value any) (int, bool) {
	num, ok := value.(int)
	return num, ok
}

// Describe uses type switch to handle different types
func Describe(value any) string {
	switch v := value.(type) {
	case nil:
		return "nil value"
	case int:
		return fmt.Sprintf("integer: %d", v)
	case float64:
		return fmt.Sprintf("float: %.2f", v)
	case string:
		return fmt.Sprintf("string: %q (length %d)", v, len(v))
	case bool:
		return fmt.Sprintf("boolean: %t", v)
	case []int:
		return fmt.Sprintf("int slice with %d elements", len(v))
	case []string:
		return fmt.Sprintf("string slice: %v", v)
	case map[string]any:
		return fmt.Sprintf("map with %d keys", len(v))
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

// ============================================================================
// PRACTICAL USE CASE: CONFIGURATION SYSTEM
// ============================================================================

type Config struct {
	settings map[string]any
}

func NewConfig() *Config {
	return &Config{settings: make(map[string]any)}
}

func (c *Config) Set(key string, value any) {
	c.settings[key] = value
}

func (c *Config) Get(key string) any {
	return c.settings[key]
}

func (c *Config) GetString(key string) string {
	if val, ok := c.settings[key].(string); ok {
		return val
	}
	return ""
}

func (c *Config) GetInt(key string) int {
	if val, ok := c.settings[key].(int); ok {
		return val
	}
	return 0
}

func (c *Config) GetBool(key string) bool {
	if val, ok := c.settings[key].(bool); ok {
		return val
	}
	return false
}

func main() {
	fmt.Println("=== The Empty Interface ===")
	fmt.Println()

	// Container with mixed types
	fmt.Println("--- Generic Container ---")
	container := NewContainer()
	container.Add("hello")
	container.Add(42)
	container.Add(3.14)
	container.Add(true)
	container.Add([]int{1, 2, 3})

	fmt.Printf("Container has %d items\n", container.Size())
	for i := 0; i < container.Size(); i++ {
		fmt.Printf("  [%d] %v (%T)\n", i, container.Get(i), container.Get(i))
	}

	fmt.Println()

	// PrintAnything function
	fmt.Println("--- PrintAnything Function ---")
	PrintAnything("a string")
	PrintAnything(123)
	PrintAnything(45.67)
	PrintAnything([]string{"a", "b", "c"})

	fmt.Println()

	// Variadic with any
	fmt.Println("--- AcceptMultiple (variadic any) ---")
	AcceptMultiple("hello", 42, true, 3.14)

	fmt.Println()

	// Dynamic data (JSON-like)
	fmt.Println("--- Dynamic Data (JSON-like) ---")
	data := createDynamicData()
	for key, value := range data {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}

	fmt.Println()

	// Safe extraction
	fmt.Println("--- Safe Value Extraction ---")
	if name, ok := SafeGetString(data["name"]); ok {
		fmt.Println("Name:", name)
	}
	if age, ok := SafeGetInt(data["age"]); ok {
		fmt.Println("Age:", age)
	}
	// This will fail safely
	if _, ok := SafeGetString(data["age"]); !ok {
		fmt.Println("Age is not a string (expected)")
	}

	fmt.Println()

	// Describe function with type switch
	fmt.Println("--- Describe Function (Type Switch) ---")
	testValues := []any{
		nil,
		42,
		3.14159,
		"hello world",
		true,
		[]int{1, 2, 3, 4, 5},
		[]string{"go", "rust", "python"},
		map[string]any{"a": 1, "b": 2},
		struct{ X int }{X: 10},
	}

	for _, val := range testValues {
		fmt.Println(" ", Describe(val))
	}

	fmt.Println()

	// Configuration system
	fmt.Println("--- Configuration System ---")
	config := NewConfig()
	config.Set("host", "localhost")
	config.Set("port", 8080)
	config.Set("debug", true)
	config.Set("timeout", 30)

	fmt.Println("Host:", config.GetString("host"))
	fmt.Println("Port:", config.GetInt("port"))
	fmt.Println("Debug:", config.GetBool("debug"))
	fmt.Println("Timeout:", config.GetInt("timeout"))

	// Getting wrong type returns zero value
	fmt.Println("Port as string:", config.GetString("port"), "(empty - wrong type)")
}

// ============================================================================
// TO RUN:
//   go run day10/04_empty_interface.go
//
// EXPECTED OUTPUT:
//   === The Empty Interface ===
//
//   --- Generic Container ---
//   Container has 5 items
//     [0] hello (string)
//     [1] 42 (int)
//     [2] 3.14 (float64)
//     [3] true (bool)
//     [4] [1 2 3] ([]int)
//
//   --- PrintAnything Function ---
//   Value: a string, Type: string
//   Value: 123, Type: int
//   Value: 45.67, Type: float64
//   Value: [a b c], Type: []string
//
//   --- AcceptMultiple (variadic any) ---
//     [0] hello (string)
//     [1] 42 (int)
//     [2] true (bool)
//     [3] 3.14 (float64)
//
//   ... (more output)
//
// EXERCISE:
//   1. Add GetFloat64 method to Config
//   2. Create a Stack using interface{} that has Push, Pop, and Peek
//   3. Write a function that flattens map[string]any to dot notation
//      e.g., {"a": {"b": 1}} -> {"a.b": 1}
//
// KEY POINTS:
//   - interface{} (or 'any') can hold any type
//   - Every type implements the empty interface
//   - Use type assertion to extract values: v.(Type)
//   - Use type switch for handling multiple types
//   - Empty interface sacrifices type safety - use sparingly
//   - Prefer generics (Go 1.18+) over interface{} when possible
//   - Common in JSON parsing, printf-style functions, configs
// ============================================================================
