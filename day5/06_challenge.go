// Day 5, Challenge: Build a Simple In-Memory Cache
//
// In this challenge, you'll build a simple cache system that:
// 1. Stores key-value pairs
// 2. Tracks access counts for each key
// 3. Can report statistics about cache usage
// 4. Can find the most/least accessed keys
//
// This combines many map concepts from today's lessons!

package main

import (
	"fmt"
	"sort"
)

// Cache represents a simple in-memory cache with access tracking
type Cache struct {
	data        map[string]string
	accessCount map[string]int
	hits        int
	misses      int
}

// NewCache creates a new empty cache
func NewCache() *Cache {
	return &Cache{
		data:        make(map[string]string),
		accessCount: make(map[string]int),
		hits:        0,
		misses:      0,
	}
}

// Set adds or updates a key-value pair in the cache
func (c *Cache) Set(key, value string) {
	c.data[key] = value
	// Initialize access count if new key
	if _, exists := c.accessCount[key]; !exists {
		c.accessCount[key] = 0
	}
}

// Get retrieves a value from the cache
// Returns the value and whether it was found
func (c *Cache) Get(key string) (string, bool) {
	value, exists := c.data[key]
	if exists {
		c.hits++
		c.accessCount[key]++
		return value, true
	}
	c.misses++
	return "", false
}

// Delete removes a key from the cache
func (c *Cache) Delete(key string) bool {
	if _, exists := c.data[key]; exists {
		delete(c.data, key)
		delete(c.accessCount, key)
		return true
	}
	return false
}

// Size returns the number of items in the cache
func (c *Cache) Size() int {
	return len(c.data)
}

// Keys returns all keys in the cache
func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Stats returns cache statistics
func (c *Cache) Stats() (hits, misses int, hitRate float64) {
	total := c.hits + c.misses
	if total == 0 {
		return c.hits, c.misses, 0.0
	}
	return c.hits, c.misses, float64(c.hits) / float64(total) * 100
}

// MostAccessed returns the n most accessed keys
func (c *Cache) MostAccessed(n int) []string {
	type keyCount struct {
		key   string
		count int
	}

	var items []keyCount
	for k, count := range c.accessCount {
		items = append(items, keyCount{k, count})
	}

	// Sort by count descending
	sort.Slice(items, func(i, j int) bool {
		return items[i].count > items[j].count
	})

	// Get top n
	result := make([]string, 0, n)
	for i := 0; i < n && i < len(items); i++ {
		result = append(result, items[i].key)
	}
	return result
}

// LeastAccessed returns the n least accessed keys
func (c *Cache) LeastAccessed(n int) []string {
	type keyCount struct {
		key   string
		count int
	}

	var items []keyCount
	for k, count := range c.accessCount {
		items = append(items, keyCount{k, count})
	}

	// Sort by count ascending
	sort.Slice(items, func(i, j int) bool {
		return items[i].count < items[j].count
	})

	// Get bottom n
	result := make([]string, 0, n)
	for i := 0; i < n && i < len(items); i++ {
		result = append(result, items[i].key)
	}
	return result
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.data = make(map[string]string)
	c.accessCount = make(map[string]int)
	c.hits = 0
	c.misses = 0
}

func main() {
	fmt.Println("=== Simple Cache Demo ===\n")

	// Create a new cache
	cache := NewCache()

	// Add some data
	fmt.Println("Adding items to cache...")
	cache.Set("user:1", "Alice")
	cache.Set("user:2", "Bob")
	cache.Set("user:3", "Charlie")
	cache.Set("config:theme", "dark")
	cache.Set("config:language", "en")

	fmt.Printf("Cache size: %d items\n", cache.Size())
	fmt.Printf("Keys: %v\n\n", cache.Keys())

	// Simulate access patterns
	fmt.Println("Simulating access patterns...")

	// user:1 is accessed frequently
	for i := 0; i < 10; i++ {
		cache.Get("user:1")
	}

	// user:2 is accessed sometimes
	for i := 0; i < 5; i++ {
		cache.Get("user:2")
	}

	// user:3 is accessed rarely
	cache.Get("user:3")

	// config values accessed a few times
	cache.Get("config:theme")
	cache.Get("config:theme")
	cache.Get("config:language")

	// Some misses
	cache.Get("nonexistent:1")
	cache.Get("nonexistent:2")

	// Report statistics
	hits, misses, hitRate := cache.Stats()
	fmt.Printf("\n=== Cache Statistics ===\n")
	fmt.Printf("Hits: %d\n", hits)
	fmt.Printf("Misses: %d\n", misses)
	fmt.Printf("Hit Rate: %.1f%%\n", hitRate)

	// Show access patterns
	fmt.Printf("\n=== Access Patterns ===\n")
	fmt.Printf("Most accessed (top 3): %v\n", cache.MostAccessed(3))
	fmt.Printf("Least accessed (bottom 3): %v\n", cache.LeastAccessed(3))

	// Demonstrate retrieval
	fmt.Printf("\n=== Data Retrieval ===\n")
	if value, found := cache.Get("user:1"); found {
		fmt.Printf("user:1 = %s\n", value)
	}

	// Delete an item
	fmt.Printf("\n=== Deletion ===\n")
	fmt.Printf("Before delete - size: %d\n", cache.Size())
	cache.Delete("user:3")
	fmt.Printf("After delete - size: %d\n", cache.Size())
	fmt.Printf("Keys: %v\n", cache.Keys())

	fmt.Println("\n=== Challenge Complete! ===")
}

// TO RUN: go run day5/06_challenge.go
//
// OUTPUT:
// === Simple Cache Demo ===
//
// Adding items to cache...
// Cache size: 5 items
// Keys: [config:language config:theme user:1 user:2 user:3]
//
// Simulating access patterns...
//
// === Cache Statistics ===
// Hits: 20
// Misses: 2
// Hit Rate: 90.9%
// ...
//
// BONUS CHALLENGES:
// 1. Add expiration time for cache entries
// 2. Implement a max size limit with LRU (Least Recently Used) eviction
// 3. Add a GetOrSet method that sets a default if key doesn't exist
// 4. Make the cache thread-safe using sync.Mutex (we'll cover this later!)
// 5. Add support for any value type using interface{} or generics
//
// KEY CONCEPTS USED:
// - Maps with different value types
// - Structs with map fields
// - Methods on structs
// - Iteration and sorting
// - The "comma ok" idiom
// - Reference semantics with pointers
// - Helper functions for common operations
