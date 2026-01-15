// Day 4, Exercise 6: Strings and Runes
//
// Key concepts:
// - Strings are immutable byte slices
// - Bytes vs Runes (Unicode code points)
// - String indexing returns bytes, not characters
// - Converting between strings, bytes, and runes
// - Common string operations with slices

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== Strings as Byte Slices ===")

	s := "Hello"
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Length (bytes): %d\n", len(s))
	fmt.Printf("Type: %T\n", s)

	// Strings are essentially []byte
	fmt.Print("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i]) // ASCII values
	}
	fmt.Println()

	// Convert to byte slice
	bytes := []byte(s)
	fmt.Printf("As []byte: %v\n", bytes)

	fmt.Println("\n=== The Unicode Challenge ===")

	// ASCII characters: 1 byte each
	ascii := "ABC"
	fmt.Printf("'%s': %d bytes, %d runes\n", ascii, len(ascii), utf8.RuneCountInString(ascii))

	// Unicode characters: multiple bytes each
	unicode := "Hello, 世界"
	fmt.Printf("'%s': %d bytes, %d runes\n", unicode, len(unicode), utf8.RuneCountInString(unicode))

	emoji := "Go🚀"
	fmt.Printf("'%s': %d bytes, %d runes\n", emoji, len(emoji), utf8.RuneCountInString(emoji))

	fmt.Println("\n=== Bytes vs Runes ===")

	text := "Hello, 世界"

	// Indexing gives bytes (WRONG for Unicode!)
	fmt.Println("Byte-by-byte (incorrect for Unicode):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%d:%c ", i, text[i])
	}
	fmt.Println()

	// Range over string gives runes (CORRECT!)
	fmt.Println("\nRune-by-rune (correct):")
	for i, r := range text {
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println()

	fmt.Println("\n=== Rune Type ===")

	// rune is an alias for int32
	var r rune = 'A'
	fmt.Printf("Rune 'A': value=%d, type=%T\n", r, r)

	r = '世'
	fmt.Printf("Rune '世': value=%d, type=%T\n", r, r)

	r = '🚀'
	fmt.Printf("Rune '🚀': value=%d, type=%T\n", r, r)

	fmt.Println("\n=== Converting Between Types ===")

	original := "Hello, 世界!"

	// String to []byte
	byteSlice := []byte(original)
	fmt.Printf("[]byte: %v\n", byteSlice)

	// String to []rune
	runeSlice := []rune(original)
	fmt.Printf("[]rune: %v\n", runeSlice)

	// []byte back to string
	fromBytes := string(byteSlice)
	fmt.Printf("From bytes: %s\n", fromBytes)

	// []rune back to string
	fromRunes := string(runeSlice)
	fmt.Printf("From runes: %s\n", fromRunes)

	// Single rune to string
	singleRune := '世'
	runeAsString := string(singleRune)
	fmt.Printf("Rune to string: %s\n", runeAsString)

	fmt.Println("\n=== String Slicing ===")

	str := "Hello, World!"

	// Slicing works on bytes!
	fmt.Println("str[0:5]:", str[0:5])   // "Hello"
	fmt.Println("str[7:]:", str[7:])     // "World!"
	fmt.Println("str[:5]:", str[:5])     // "Hello"

	// WARNING: Slicing Unicode strings by bytes can corrupt!
	unicodeStr := "Hello, 世界"
	// unicodeStr[7:9] would cut '世' in half - BAD!

	// Safe way: convert to runes first
	runes := []rune(unicodeStr)
	first5Runes := string(runes[:5])
	last2Runes := string(runes[len(runes)-2:])
	fmt.Printf("First 5 runes: %s\n", first5Runes)
	fmt.Printf("Last 2 runes: %s\n", last2Runes)

	fmt.Println("\n=== Strings Are Immutable ===")

	immutable := "Hello"
	// immutable[0] = 'h'  // ERROR: cannot assign

	// To modify, convert to []byte or []rune
	mutableBytes := []byte(immutable)
	mutableBytes[0] = 'h'
	modified := string(mutableBytes)
	fmt.Printf("Original: %s, Modified: %s\n", immutable, modified)

	fmt.Println("\n=== Common String Operations ===")

	sample := "  Hello, Go World!  "

	// Using strings package
	fmt.Printf("Original: '%s'\n", sample)
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(sample))
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(sample))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(sample))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(sample, "Go"))
	fmt.Printf("Index of 'Go': %d\n", strings.Index(sample, "Go"))
	fmt.Printf("Replace: '%s'\n", strings.Replace(sample, "Go", "Golang", 1))
	fmt.Printf("Split: %v\n", strings.Split(strings.TrimSpace(sample), " "))

	fmt.Println("\n=== Building Strings Efficiently ===")

	// Inefficient: string concatenation creates new strings
	// var result string
	// for i := 0; i < 1000; i++ {
	//     result += "x"  // Creates new string each time!
	// }

	// Efficient: use strings.Builder
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		builder.WriteString("Go")
		builder.WriteRune('!')
	}
	result := builder.String()
	fmt.Println("Builder result:", result)

	// Or use []byte and append
	var buffer []byte
	for i := 0; i < 5; i++ {
		buffer = append(buffer, "Hi "...)
	}
	fmt.Println("Buffer result:", string(buffer))

	fmt.Println("\n=== Practical Example: Reverse a String ===")

	forward := "Hello, 世界!"
	fmt.Printf("Original: %s\n", forward)

	// Convert to runes, reverse, convert back
	runesRev := []rune(forward)
	for i, j := 0, len(runesRev)-1; i < j; i, j = i+1, j-1 {
		runesRev[i], runesRev[j] = runesRev[j], runesRev[i]
	}
	reversed := string(runesRev)
	fmt.Printf("Reversed: %s\n", reversed)

	fmt.Println("\n=== Practical Example: Count Characters ===")

	phrase := "Hello, 世界! 🚀"
	byteCount := len(phrase)
	runeCount := utf8.RuneCountInString(phrase)
	wordCount := len(strings.Fields(phrase))

	fmt.Printf("Phrase: %s\n", phrase)
	fmt.Printf("Bytes: %d\n", byteCount)
	fmt.Printf("Characters (runes): %d\n", runeCount)
	fmt.Printf("Words: %d\n", wordCount)

	fmt.Println("\n=== Practical Example: Is Palindrome? ===")

	words := []string{"radar", "hello", "level", "世界世"}

	for _, word := range words {
		fmt.Printf("'%s' is palindrome: %t\n", word, isPalindrome(word))
	}
}

// Check if a string is a palindrome (works with Unicode)
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// TO RUN: go run day4/06_strings_runes.go
//
// OUTPUT:
// === Strings as Byte Slices ===
// String: Hello
// Length (bytes): 5
// ...
//
// EXERCISE:
// 1. Write a function that counts vowels in any Unicode string
// 2. Create a function to capitalize the first letter of each word
// 3. Implement a simple Caesar cipher that works with any alphabet
// 4. Write a function that removes all non-alphanumeric characters
//
// KEY POINTS:
// - Strings are immutable byte slices
// - len(s) returns bytes, not characters
// - Use utf8.RuneCountInString() for character count
// - Range over strings iterates by runes
// - Convert to []rune for safe Unicode manipulation
// - Use strings.Builder for efficient string building
