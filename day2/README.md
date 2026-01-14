# Day 2: Control Flow

## What You'll Learn
- If/else statements
- For loops (Go's only loop!)
- Switch statements
- Break, continue, and labels

## Files in Order

| File | Topic | Run Command |
|------|-------|-------------|
| 01_if_else.go | Conditionals | `go run 01_if_else.go` |
| 02_for_loops.go | All loop forms | `go run 02_for_loops.go` |
| 03_switch.go | Switch statements | `go run 03_switch.go` |
| 04_fizzbuzz.go | Classic challenge | `go run 04_fizzbuzz.go` |
| 05_patterns.go | Nested loop practice | `go run 05_patterns.go` |
| 06_guessing_game.go | Mini project | `go run 06_guessing_game.go` |
| 07_challenge.go | Prime checker | `go run 07_challenge.go` |

## Quick Reference

### If/Else
```go
// Basic
if x > 10 {
    fmt.Println("big")
} else if x > 5 {
    fmt.Println("medium")
} else {
    fmt.Println("small")
}

// With initialization (variable scoped to if block)
if num := calculate(); num > 0 {
    fmt.Println("positive")
}
```

### For Loops
```go
// Standard
for i := 0; i < 10; i++ { }

// While-style
for condition { }

// Infinite
for { }

// Range
for index, value := range collection { }

// Skip index
for _, value := range collection { }
```

### Switch
```go
// Basic switch
switch day {
case 1:
    fmt.Println("Monday")
case 6, 7:
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}

// Expressionless (like if-else chain)
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}
```

### Keywords
- `break` - exit loop
- `continue` - skip to next iteration
- `fallthrough` - continue to next case in switch

## Checklist
- [ ] Run all exercises
- [ ] Complete FizzBuzz
- [ ] Play the guessing game
- [ ] Build the prime checker
- [ ] Create your own pattern

## Next: Day 3
Functions, multiple returns, and variadic functions
