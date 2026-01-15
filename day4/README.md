# Day 4: Arrays, Slices & Strings

## Quick Reference

### Arrays (Fixed Size)
```go
// Declaration
var arr [5]int                    // Zero-valued
arr := [5]int{1, 2, 3, 4, 5}     // Initialized
arr := [...]int{1, 2, 3}          // Compiler counts
arr := [5]int{0: 10, 4: 50}       // Sparse

// Access
arr[0]                            // First element
arr[len(arr)-1]                   // Last element

// Iteration
for i, v := range arr { }         // Index and value
for i := range arr { }            // Index only
for _, v := range arr { }         // Value only
```

### Slices (Dynamic)
```go
// Creation
s := []int{1, 2, 3}               // Literal
s := make([]int, len, cap)        // With make
var s []int                       // Nil slice (preferred)
s := []int{}                      // Empty slice

// Properties
len(s)                            // Current length
cap(s)                            // Capacity

// Slicing
s[start:end]                      // Elements start to end-1
s[:end]                           // From beginning
s[start:]                         // To end
s[:]                              // Full slice
s[low:high:max]                   // Full slice expression (limits capacity)

// Operations
s = append(s, elem)               // Append one
s = append(s, 1, 2, 3)            // Append multiple
s = append(s, other...)           // Append slice
copy(dst, src)                    // Copy elements
```

### Slice Manipulation
```go
// Remove element at index i
s = append(s[:i], s[i+1:]...)

// Insert at index i
s = append(s[:i], append([]T{val}, s[i:]...)...)
// Or use slices.Insert(s, i, val)

// Remove first
s = s[1:]

// Remove last
s = s[:len(s)-1]

// Copy (independent)
dst := make([]int, len(src))
copy(dst, src)
```

### slices Package (Go 1.21+)
```go
import "slices"

slices.Sort(s)                    // Sort in place
slices.Reverse(s)                 // Reverse in place
slices.Contains(s, v)             // Check if contains
slices.Index(s, v)                // Find index (-1 if not found)
slices.Insert(s, i, v)            // Insert at index
slices.Delete(s, i, j)            // Delete range [i:j]
slices.Compact(s)                 // Remove consecutive duplicates
slices.IsSorted(s)                // Check if sorted
```

### 2D Slices
```go
// Literal
grid := [][]int{
    {1, 2, 3},
    {4, 5, 6},
}

// Dynamic creation
rows, cols := 3, 4
grid := make([][]int, rows)
for i := range grid {
    grid[i] = make([]int, cols)
}

// Access
grid[row][col]
```

### Strings
```go
// Strings are immutable byte slices
s := "Hello, 世界"

len(s)                            // Byte count
utf8.RuneCountInString(s)         // Character count

// Iteration
for i, r := range s { }           // By rune (correct for Unicode)
for i := 0; i < len(s); i++ { }   // By byte (ASCII only)

// Conversion
[]byte(s)                         // String to bytes
[]rune(s)                         // String to runes
string(bytes)                     // Bytes to string
string(runes)                     // Runes to string

// strings package
strings.TrimSpace(s)
strings.ToUpper(s) / ToLower(s)
strings.Contains(s, sub)
strings.Index(s, sub)
strings.Split(s, sep)
strings.Join(slice, sep)
strings.Replace(s, old, new, n)

// Efficient string building
var b strings.Builder
b.WriteString("Hello")
b.WriteRune('!')
result := b.String()
```

## Key Concepts

| Concept | Array | Slice |
|---------|-------|-------|
| Size | Fixed (part of type) | Dynamic |
| Type | `[n]T` | `[]T` |
| Zero value | All elements zero | `nil` |
| Comparison | `==` works | Cannot use `==` |
| Pass to function | Copied (value) | Reference |

## Common Gotchas

1. **Shared backing array**: Sub-slices share memory with original
   ```go
   a := []int{1, 2, 3, 4, 5}
   b := a[1:3]
   b[0] = 100  // Also changes a[1]!
   ```

2. **Append may reallocate**: Always reassign
   ```go
   s = append(s, v)  // Correct
   append(s, v)      // WRONG - result lost
   ```

3. **len(string) returns bytes, not characters**
   ```go
   s := "世界"
   len(s)                        // 6 (bytes)
   utf8.RuneCountInString(s)     // 2 (characters)
   ```

4. **String indexing returns bytes**
   ```go
   s := "世界"
   s[0]            // Returns a byte, not '世'
   []rune(s)[0]    // Returns '世'
   ```

## Files in This Day

| File | Topic |
|------|-------|
| `01_arrays.go` | Array basics, declaration, iteration |
| `02_slices_basics.go` | Slice creation, append, copy |
| `03_slice_operations.go` | Remove, insert, filter, sort |
| `04_slice_internals.go` | Capacity, memory, gotchas |
| `05_multidimensional.go` | 2D arrays and slices |
| `06_strings_runes.go` | Strings, bytes, runes, Unicode |
| `07_challenge.go` | Student Grade Manager |

## Run Commands

```bash
go run day4/01_arrays.go
go run day4/02_slices_basics.go
go run day4/03_slice_operations.go
go run day4/04_slice_internals.go
go run day4/05_multidimensional.go
go run day4/06_strings_runes.go
go run day4/07_challenge.go
```
