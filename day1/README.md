# Day 1: Go Fundamentals

## What You'll Learn
- Setting up Go
- Package structure
- Variables and types
- Constants and iota
- Formatted printing
- User input
- Operators

## Files in Order

| File | Topic | Run Command |
|------|-------|-------------|
| 01_hello.go | Hello World | `go run 01_hello.go` |
| 02_variables.go | Variables & Types | `go run 02_variables.go` |
| 03_constants.go | Constants & iota | `go run 03_constants.go` |
| 04_printf.go | Formatted Output | `go run 04_printf.go` |
| 05_input.go | User Input | `go run 05_input.go` |
| 06_operators.go | Operators | `go run 06_operators.go` |
| 07_challenge.go | Calculator Challenge | `go run 07_challenge.go` |

## Quick Reference

### Variable Declaration
```go
var name string = "Go"   // explicit type
var age = 25             // type inferred
count := 10              // short declaration (most common)
```

### Basic Types
```go
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
string
bool
byte (alias for uint8)
rune (alias for int32, represents Unicode)
```

### Printf Verbs
```go
%v  - default format
%T  - type
%s  - string
%d  - integer
%f  - float (%.2f for 2 decimals)
%t  - boolean
%p  - pointer
```

## Checklist
- [ ] Run all 7 exercises
- [ ] Complete the TODOs in each file
- [ ] Build the calculator challenge
- [ ] Experiment with your own variables

## Next: Day 2
Control flow: if/else, for loops, switch statements
