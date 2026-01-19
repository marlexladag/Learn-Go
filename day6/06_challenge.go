// Day 6 Challenge: Linked List Implementation
//
// Implement a doubly-linked list using pointers.
// This challenge combines all pointer concepts from today.

package main

import "fmt"

func main() {
	fmt.Println("=== Doubly Linked List Challenge ===")

	list := NewLinkedList()

	// Test empty list
	fmt.Println("Empty list:", list)
	fmt.Println("Size:", list.Size())

	// Add elements
	fmt.Println("\n--- Adding elements ---")
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)
	list.PushFront(5)
	list.PushFront(1)

	fmt.Println("After pushes:", list) // 1 <-> 5 <-> 10 <-> 20 <-> 30
	fmt.Println("Size:", list.Size())

	// Access elements
	fmt.Println("\n--- Accessing elements ---")
	if front, ok := list.Front(); ok {
		fmt.Println("Front:", front)
	}
	if back, ok := list.Back(); ok {
		fmt.Println("Back:", back)
	}

	// Get by index
	if val, ok := list.Get(2); ok {
		fmt.Println("Index 2:", val)
	}

	// Iterate forward
	fmt.Println("\n--- Forward iteration ---")
	list.ForEach(func(val int) {
		fmt.Print(val, " ")
	})
	fmt.Println()

	// Iterate backward
	fmt.Println("\n--- Backward iteration ---")
	list.ForEachReverse(func(val int) {
		fmt.Print(val, " ")
	})
	fmt.Println()

	// Remove elements
	fmt.Println("\n--- Removing elements ---")
	list.PopFront()
	fmt.Println("After PopFront:", list)

	list.PopBack()
	fmt.Println("After PopBack:", list)

	// Remove specific value
	list.Remove(10)
	fmt.Println("After Remove(10):", list)

	// Find element
	fmt.Println("\n--- Finding elements ---")
	fmt.Println("Contains 20:", list.Contains(20))
	fmt.Println("Contains 100:", list.Contains(100))

	if idx := list.IndexOf(20); idx != -1 {
		fmt.Println("Index of 20:", idx)
	}

	// Insert at position
	fmt.Println("\n--- Insert at position ---")
	list.InsertAt(1, 15)
	fmt.Println("After InsertAt(1, 15):", list)

	// Clear list
	fmt.Println("\n--- Clear ---")
	list.Clear()
	fmt.Println("After Clear:", list)
	fmt.Println("IsEmpty:", list.IsEmpty())

	// Bonus: Convert to slice
	fmt.Println("\n--- Bonus operations ---")
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	fmt.Println("As slice:", list.ToSlice())

	// Create from slice
	list2 := LinkedListFromSlice([]int{100, 200, 300})
	fmt.Println("From slice:", list2)
}

// ListNode represents a node in the doubly linked list
type ListNode struct {
	Value int
	Prev  *ListNode
	Next  *ListNode
}

// LinkedList is a doubly linked list
type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

// NewLinkedList creates an empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// LinkedListFromSlice creates a list from a slice
func LinkedListFromSlice(values []int) *LinkedList {
	list := NewLinkedList()
	for _, v := range values {
		list.PushBack(v)
	}
	return list
}

// Size returns the number of elements
func (l *LinkedList) Size() int {
	return l.size
}

// IsEmpty returns true if list has no elements
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// PushFront adds element at the beginning
func (l *LinkedList) PushFront(value int) {
	node := &ListNode{Value: value}

	if l.head == nil {
		// Empty list
		l.head = node
		l.tail = node
	} else {
		// Insert before head
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}
	l.size++
}

// PushBack adds element at the end
func (l *LinkedList) PushBack(value int) {
	node := &ListNode{Value: value}

	if l.tail == nil {
		// Empty list
		l.head = node
		l.tail = node
	} else {
		// Insert after tail
		node.Prev = l.tail
		l.tail.Next = node
		l.tail = node
	}
	l.size++
}

// PopFront removes and returns the first element
func (l *LinkedList) PopFront() (int, bool) {
	if l.head == nil {
		return 0, false
	}

	value := l.head.Value

	if l.head == l.tail {
		// Single element
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.Next
		l.head.Prev = nil
	}
	l.size--
	return value, true
}

// PopBack removes and returns the last element
func (l *LinkedList) PopBack() (int, bool) {
	if l.tail == nil {
		return 0, false
	}

	value := l.tail.Value

	if l.head == l.tail {
		// Single element
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	}
	l.size--
	return value, true
}

// Front returns the first element without removing
func (l *LinkedList) Front() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	return l.head.Value, true
}

// Back returns the last element without removing
func (l *LinkedList) Back() (int, bool) {
	if l.tail == nil {
		return 0, false
	}
	return l.tail.Value, true
}

// Get returns element at index
func (l *LinkedList) Get(index int) (int, bool) {
	node := l.nodeAt(index)
	if node == nil {
		return 0, false
	}
	return node.Value, true
}

// nodeAt returns node at index (internal helper)
func (l *LinkedList) nodeAt(index int) *ListNode {
	if index < 0 || index >= l.size {
		return nil
	}

	// Optimize: start from closer end
	if index < l.size/2 {
		// Start from head
		current := l.head
		for i := 0; i < index; i++ {
			current = current.Next
		}
		return current
	} else {
		// Start from tail
		current := l.tail
		for i := l.size - 1; i > index; i-- {
			current = current.Prev
		}
		return current
	}
}

// InsertAt inserts value at specific index
func (l *LinkedList) InsertAt(index int, value int) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.PushFront(value)
		return true
	}

	if index == l.size {
		l.PushBack(value)
		return true
	}

	// Insert in middle
	nextNode := l.nodeAt(index)
	prevNode := nextNode.Prev

	node := &ListNode{
		Value: value,
		Prev:  prevNode,
		Next:  nextNode,
	}

	prevNode.Next = node
	nextNode.Prev = node
	l.size++
	return true
}

// Remove removes first occurrence of value
func (l *LinkedList) Remove(value int) bool {
	current := l.head

	for current != nil {
		if current.Value == value {
			l.removeNode(current)
			return true
		}
		current = current.Next
	}
	return false
}

// removeNode removes a specific node (internal helper)
func (l *LinkedList) removeNode(node *ListNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.tail = node.Prev
	}

	l.size--
}

// Contains checks if value exists in list
func (l *LinkedList) Contains(value int) bool {
	return l.IndexOf(value) != -1
}

// IndexOf returns index of value, or -1 if not found
func (l *LinkedList) IndexOf(value int) int {
	current := l.head
	index := 0

	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// ForEach iterates through all elements
func (l *LinkedList) ForEach(fn func(int)) {
	current := l.head
	for current != nil {
		fn(current.Value)
		current = current.Next
	}
}

// ForEachReverse iterates in reverse order
func (l *LinkedList) ForEachReverse(fn func(int)) {
	current := l.tail
	for current != nil {
		fn(current.Value)
		current = current.Prev
	}
}

// Clear removes all elements
func (l *LinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// ToSlice converts list to slice
func (l *LinkedList) ToSlice() []int {
	result := make([]int, 0, l.size)
	l.ForEach(func(val int) {
		result = append(result, val)
	})
	return result
}

// String returns string representation
func (l *LinkedList) String() string {
	if l.IsEmpty() {
		return "[]"
	}

	result := "["
	current := l.head
	for current != nil {
		result += fmt.Sprintf("%d", current.Value)
		if current.Next != nil {
			result += " <-> "
		}
		current = current.Next
	}
	result += "]"
	return result
}

// TO RUN: go run day6/06_challenge.go
//
// OUTPUT:
// === Doubly Linked List Challenge ===
// Empty list: []
// Size: 0
// ...
//
// BONUS CHALLENGES:
// 1. Add a Reverse() method that reverses the list in place
// 2. Add a Copy() method that creates a deep copy
// 3. Add a Sort() method (hint: you can convert to slice, sort, recreate)
// 4. Add a Merge() method that merges two sorted lists
// 5. Make it generic using Go generics: LinkedList[T any]
//
// KEY CONCEPTS DEMONSTRATED:
// - Pointer-based data structures
// - Bidirectional node references
// - nil handling for edge cases
// - Helper methods for internal operations
// - Iterator patterns with function callbacks
