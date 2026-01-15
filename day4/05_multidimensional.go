// Day 4, Exercise 5: Multidimensional Arrays and Slices
//
// Key concepts:
// - 2D arrays: fixed rows and columns
// - 2D slices: flexible "slice of slices"
// - Jagged arrays (rows of different lengths)
// - Iterating over 2D structures
// - Common use cases: matrices, grids, tables

package main

import "fmt"

func main() {
	fmt.Println("=== 2D Arrays ===")

	// Declare 2D array: [rows][cols]type
	var matrix [3][4]int // 3 rows, 4 columns
	fmt.Println("Zero-value 2D array:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Initialize 2D array
	grid := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\nInitialized 3x3 grid:")
	for _, row := range grid {
		fmt.Println(row)
	}

	// Access elements: matrix[row][col]
	fmt.Printf("\ngrid[0][0] = %d (top-left)\n", grid[0][0])
	fmt.Printf("grid[1][1] = %d (center)\n", grid[1][1])
	fmt.Printf("grid[2][2] = %d (bottom-right)\n", grid[2][2])

	fmt.Println("\n=== 2D Slices (Slice of Slices) ===")

	// Method 1: Literal initialization
	table := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D slice literal:")
	for _, row := range table {
		fmt.Println(row)
	}

	// Method 2: Using make (two-step process)
	rows, cols := 3, 4
	dynamic := make([][]int, rows) // Create outer slice
	for i := range dynamic {
		dynamic[i] = make([]int, cols) // Create each inner slice
	}
	fmt.Printf("\nDynamic %dx%d slice:\n", rows, cols)
	for _, row := range dynamic {
		fmt.Println(row)
	}

	fmt.Println("\n=== Jagged Arrays (Variable Row Lengths) ===")

	// Rows can have different lengths!
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	fmt.Println("Jagged array (triangle):")
	for i, row := range jagged {
		fmt.Printf("Row %d (len=%d): %v\n", i, len(row), row)
	}

	fmt.Println("\n=== Iterating 2D Structures ===")

	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Method 1: Nested range
	fmt.Println("Using nested range:")
	for rowIdx, row := range data {
		for colIdx, val := range row {
			fmt.Printf("[%d][%d]=%d ", rowIdx, colIdx, val)
		}
		fmt.Println()
	}

	// Method 2: Traditional nested loops
	fmt.Println("\nUsing traditional loops:")
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("%d ", data[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\n=== Practical Example: Tic-Tac-Toe Board ===")

	board := [3][3]string{
		{"X", "O", "X"},
		{" ", "X", "O"},
		{"O", " ", "X"},
	}

	fmt.Println("Tic-Tac-Toe:")
	for i, row := range board {
		fmt.Printf(" %s | %s | %s \n", row[0], row[1], row[2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}

	fmt.Println("\n=== Practical Example: Matrix Operations ===")

	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}
	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}

	// Matrix addition
	result := make([][]int, len(matrixA))
	for i := range result {
		result[i] = make([]int, len(matrixA[i]))
		for j := range result[i] {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	fmt.Println("Matrix A:")
	printMatrix(matrixA)
	fmt.Println("Matrix B:")
	printMatrix(matrixB)
	fmt.Println("A + B:")
	printMatrix(result)

	fmt.Println("=== Practical Example: Grid-based Game Map ===")

	// 0=empty, 1=wall, 2=player, 3=treasure
	gameMap := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 2, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 3, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	symbols := map[int]string{0: ".", 1: "#", 2: "@", 3: "$"}

	fmt.Println("Game Map:")
	for _, row := range gameMap {
		for _, cell := range row {
			fmt.Print(symbols[cell], " ")
		}
		fmt.Println()
	}
	fmt.Println("Legend: # = wall, @ = player, $ = treasure, . = empty")

	fmt.Println("\n=== Transposing a Matrix ===")

	original := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// Transpose: swap rows and columns
	transposed := make([][]int, len(original[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(original))
		for j := range transposed[i] {
			transposed[i][j] = original[j][i]
		}
	}

	fmt.Println("Original (2x3):")
	printMatrix(original)
	fmt.Println("Transposed (3x2):")
	printMatrix(transposed)

	fmt.Println("=== Finding Elements in 2D ===")

	searchGrid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	target := 5
	found, row, col := find2D(searchGrid, target)
	if found {
		fmt.Printf("Found %d at position [%d][%d]\n", target, row, col)
	}

	target = 10
	found, _, _ = find2D(searchGrid, target)
	if !found {
		fmt.Printf("%d not found in grid\n", target)
	}
}

// Helper function to print a matrix
func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
}

// Helper function to find element in 2D slice
func find2D(grid [][]int, target int) (bool, int, int) {
	for i, row := range grid {
		for j, val := range row {
			if val == target {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

// TO RUN: go run day4/05_multidimensional.go
//
// OUTPUT:
// === 2D Arrays ===
// Zero-value 2D array:
// [0 0 0 0]
// [0 0 0 0]
// [0 0 0 0]
// ...
//
// EXERCISE:
// 1. Create a 5x5 identity matrix (1s on diagonal, 0s elsewhere)
// 2. Create a multiplication table (10x10)
// 3. Implement a function to rotate a matrix 90 degrees
// 4. Create a simple maze and find a path from start to end
//
// KEY POINTS:
// - 2D arrays: [rows][cols]type, fixed size
// - 2D slices: [][]type, dynamic, can be jagged
// - Access: grid[row][col]
// - Create dynamic 2D slice: make outer, then make each inner
// - Use nested range for iteration
