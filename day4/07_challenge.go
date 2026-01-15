// Day 4, Exercise 7: Challenge - Student Grade Manager
//
// This challenge combines all Day 4 concepts:
// - Arrays for fixed-size data
// - Slices for dynamic lists
// - Slice operations (append, remove, sort)
// - 2D slices for grade tables
// - Strings and runes for name handling
//
// Build a complete grade management system!

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Student represents a student with their grades
type Student struct {
	Name   string
	Grades []float64
}

// GradeManager holds all students
type GradeManager struct {
	Students []Student
	Subjects []string
}

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     STUDENT GRADE MANAGER v1.0         ║")
	fmt.Println("║     Day 4 Challenge: Arrays & Slices   ║")
	fmt.Println("╚════════════════════════════════════════╝")

	// Initialize with default subjects
	manager := GradeManager{
		Subjects: []string{"Math", "Science", "English", "History"},
	}

	// Add some sample students
	manager.AddStudent("Alice Johnson", []float64{95, 88, 92, 85})
	manager.AddStudent("Bob Smith", []float64{78, 82, 75, 88})
	manager.AddStudent("Charlie Brown", []float64{88, 91, 84, 79})
	manager.AddStudent("Diana Ross", []float64{92, 95, 98, 94})
	manager.AddStudent("Eve Wilson", []float64{70, 68, 72, 75})

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. View all students")
		fmt.Println("2. View grade table")
		fmt.Println("3. Add new student")
		fmt.Println("4. Update grade")
		fmt.Println("5. Remove student")
		fmt.Println("6. View statistics")
		fmt.Println("7. Sort students")
		fmt.Println("8. Search student")
		fmt.Println("9. Exit")
		fmt.Print("\nChoice: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			manager.ViewAllStudents()
		case "2":
			manager.ViewGradeTable()
		case "3":
			manager.InteractiveAddStudent(reader)
		case "4":
			manager.InteractiveUpdateGrade(reader)
		case "5":
			manager.InteractiveRemoveStudent(reader)
		case "6":
			manager.ViewStatistics()
		case "7":
			manager.InteractiveSortStudents(reader)
		case "8":
			manager.InteractiveSearchStudent(reader)
		case "9":
			fmt.Println("\nGoodbye! Keep studying! 📚")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// AddStudent adds a new student to the manager
func (gm *GradeManager) AddStudent(name string, grades []float64) {
	// Ensure grades slice matches subjects
	if len(grades) < len(gm.Subjects) {
		// Pad with zeros
		for len(grades) < len(gm.Subjects) {
			grades = append(grades, 0)
		}
	} else if len(grades) > len(gm.Subjects) {
		// Trim to match
		grades = grades[:len(gm.Subjects)]
	}

	student := Student{
		Name:   name,
		Grades: grades,
	}
	gm.Students = append(gm.Students, student)
}

// RemoveStudent removes a student by index
func (gm *GradeManager) RemoveStudent(index int) bool {
	if index < 0 || index >= len(gm.Students) {
		return false
	}
	gm.Students = append(gm.Students[:index], gm.Students[index+1:]...)
	return true
}

// FindStudent searches for a student by name (case-insensitive)
func (gm *GradeManager) FindStudent(name string) (int, *Student) {
	nameLower := strings.ToLower(name)
	for i := range gm.Students {
		if strings.Contains(strings.ToLower(gm.Students[i].Name), nameLower) {
			return i, &gm.Students[i]
		}
	}
	return -1, nil
}

// ViewAllStudents displays all students with their averages
func (gm *GradeManager) ViewAllStudents() {
	if len(gm.Students) == 0 {
		fmt.Println("\nNo students registered.")
		return
	}

	fmt.Println("\n┌─────────────────────────────────────────────────┐")
	fmt.Println("│              ALL STUDENTS                       │")
	fmt.Println("├────┬──────────────────────┬──────────┬──────────┤")
	fmt.Println("│ #  │ Name                 │ Average  │ Grade    │")
	fmt.Println("├────┼──────────────────────┼──────────┼──────────┤")

	for i, student := range gm.Students {
		avg := calculateAverage(student.Grades)
		grade := getLetterGrade(avg)
		// Truncate name if too long (handle Unicode properly)
		displayName := truncateString(student.Name, 20)
		fmt.Printf("│ %2d │ %-20s │ %6.2f   │    %s     │\n",
			i+1, displayName, avg, grade)
	}
	fmt.Println("└────┴──────────────────────┴──────────┴──────────┘")
}

// ViewGradeTable shows a full 2D table of all grades
func (gm *GradeManager) ViewGradeTable() {
	if len(gm.Students) == 0 {
		fmt.Println("\nNo students registered.")
		return
	}

	fmt.Println("\n=== GRADE TABLE ===")

	// Print header
	fmt.Printf("%-18s", "Student")
	for _, subject := range gm.Subjects {
		fmt.Printf(" │ %8s", subject)
	}
	fmt.Printf(" │ %8s\n", "Avg")

	// Print separator
	fmt.Print(strings.Repeat("-", 18))
	for range gm.Subjects {
		fmt.Print("-┼----------")
	}
	fmt.Println("-┼----------")

	// Print each student's grades
	for _, student := range gm.Students {
		displayName := truncateString(student.Name, 17)
		fmt.Printf("%-18s", displayName)
		for _, grade := range student.Grades {
			fmt.Printf(" │ %8.1f", grade)
		}
		avg := calculateAverage(student.Grades)
		fmt.Printf(" │ %8.2f\n", avg)
	}

	// Print subject averages
	fmt.Print(strings.Repeat("-", 18))
	for range gm.Subjects {
		fmt.Print("-┼----------")
	}
	fmt.Println("-┼----------")

	fmt.Printf("%-18s", "Class Average")
	for i := range gm.Subjects {
		var sum float64
		for _, student := range gm.Students {
			sum += student.Grades[i]
		}
		avg := sum / float64(len(gm.Students))
		fmt.Printf(" │ %8.1f", avg)
	}
	// Overall class average
	var totalSum float64
	var totalCount int
	for _, student := range gm.Students {
		for _, grade := range student.Grades {
			totalSum += grade
			totalCount++
		}
	}
	fmt.Printf(" │ %8.2f\n", totalSum/float64(totalCount))
}

// ViewStatistics shows various statistics
func (gm *GradeManager) ViewStatistics() {
	if len(gm.Students) == 0 {
		fmt.Println("\nNo students registered.")
		return
	}

	fmt.Println("\n=== STATISTICS ===")

	// Collect all averages
	var averages []float64
	var topStudent Student
	var lowStudent Student
	topAvg := 0.0
	lowAvg := 100.0

	for _, student := range gm.Students {
		avg := calculateAverage(student.Grades)
		averages = append(averages, avg)

		if avg > topAvg {
			topAvg = avg
			topStudent = student
		}
		if avg < lowAvg {
			lowAvg = avg
			lowStudent = student
		}
	}

	// Overall class average
	classAvg := calculateAverage(averages)

	fmt.Printf("\nTotal Students: %d\n", len(gm.Students))
	fmt.Printf("Total Subjects: %d\n", len(gm.Subjects))
	fmt.Printf("Class Average: %.2f\n", classAvg)
	fmt.Printf("\nTop Performer: %s (%.2f)\n", topStudent.Name, topAvg)
	fmt.Printf("Needs Improvement: %s (%.2f)\n", lowStudent.Name, lowAvg)

	// Grade distribution
	fmt.Println("\n--- Grade Distribution ---")
	gradeCount := make(map[string]int)
	for _, avg := range averages {
		grade := getLetterGrade(avg)
		gradeCount[grade]++
	}
	for _, grade := range []string{"A", "B", "C", "D", "F"} {
		count := gradeCount[grade]
		bar := strings.Repeat("█", count*3)
		fmt.Printf("%s: %s %d\n", grade, bar, count)
	}

	// Subject analysis
	fmt.Println("\n--- Subject Analysis ---")
	for i, subject := range gm.Subjects {
		var sum float64
		var min, max float64 = 100, 0
		for _, student := range gm.Students {
			grade := student.Grades[i]
			sum += grade
			if grade < min {
				min = grade
			}
			if grade > max {
				max = grade
			}
		}
		avg := sum / float64(len(gm.Students))
		fmt.Printf("%s: Avg=%.1f, Min=%.1f, Max=%.1f\n",
			subject, avg, min, max)
	}
}

// InteractiveAddStudent adds a student with user input
func (gm *GradeManager) InteractiveAddStudent(reader *bufio.Reader) {
	fmt.Print("\nEnter student name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty.")
		return
	}

	grades := make([]float64, len(gm.Subjects))
	for i, subject := range gm.Subjects {
		fmt.Printf("Enter %s grade (0-100): ", subject)
		input, _ := reader.ReadString('\n')
		grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil || grade < 0 || grade > 100 {
			fmt.Println("Invalid grade, using 0.")
			grade = 0
		}
		grades[i] = grade
	}

	gm.AddStudent(name, grades)
	fmt.Printf("\n✓ Added %s successfully!\n", name)
}

// InteractiveUpdateGrade updates a specific grade
func (gm *GradeManager) InteractiveUpdateGrade(reader *bufio.Reader) {
	fmt.Print("\nEnter student name to update: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	idx, student := gm.FindStudent(name)
	if idx == -1 {
		fmt.Println("Student not found.")
		return
	}

	fmt.Printf("Found: %s\n", student.Name)
	fmt.Println("Subjects:")
	for i, subject := range gm.Subjects {
		fmt.Printf("%d. %s (current: %.1f)\n", i+1, subject, student.Grades[i])
	}

	fmt.Print("Enter subject number: ")
	input, _ := reader.ReadString('\n')
	subjectIdx, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || subjectIdx < 1 || subjectIdx > len(gm.Subjects) {
		fmt.Println("Invalid subject number.")
		return
	}
	subjectIdx-- // Convert to 0-based

	fmt.Print("Enter new grade (0-100): ")
	input, _ = reader.ReadString('\n')
	newGrade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil || newGrade < 0 || newGrade > 100 {
		fmt.Println("Invalid grade.")
		return
	}

	oldGrade := student.Grades[subjectIdx]
	student.Grades[subjectIdx] = newGrade
	fmt.Printf("✓ Updated %s's %s: %.1f → %.1f\n",
		student.Name, gm.Subjects[subjectIdx], oldGrade, newGrade)
}

// InteractiveRemoveStudent removes a student
func (gm *GradeManager) InteractiveRemoveStudent(reader *bufio.Reader) {
	fmt.Print("\nEnter student name to remove: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	idx, student := gm.FindStudent(name)
	if idx == -1 {
		fmt.Println("Student not found.")
		return
	}

	fmt.Printf("Are you sure you want to remove %s? (y/n): ", student.Name)
	confirm, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
		gm.RemoveStudent(idx)
		fmt.Printf("✓ Removed %s\n", student.Name)
	} else {
		fmt.Println("Cancelled.")
	}
}

// InteractiveSortStudents sorts students by various criteria
func (gm *GradeManager) InteractiveSortStudents(reader *bufio.Reader) {
	fmt.Println("\nSort by:")
	fmt.Println("1. Name (A-Z)")
	fmt.Println("2. Name (Z-A)")
	fmt.Println("3. Average (Highest first)")
	fmt.Println("4. Average (Lowest first)")
	fmt.Print("Choice: ")

	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		slices.SortFunc(gm.Students, func(a, b Student) int {
			return strings.Compare(a.Name, b.Name)
		})
		fmt.Println("✓ Sorted by name (A-Z)")
	case "2":
		slices.SortFunc(gm.Students, func(a, b Student) int {
			return strings.Compare(b.Name, a.Name)
		})
		fmt.Println("✓ Sorted by name (Z-A)")
	case "3":
		slices.SortFunc(gm.Students, func(a, b Student) int {
			avgA := calculateAverage(a.Grades)
			avgB := calculateAverage(b.Grades)
			if avgB > avgA {
				return 1
			} else if avgB < avgA {
				return -1
			}
			return 0
		})
		fmt.Println("✓ Sorted by average (highest first)")
	case "4":
		slices.SortFunc(gm.Students, func(a, b Student) int {
			avgA := calculateAverage(a.Grades)
			avgB := calculateAverage(b.Grades)
			if avgA > avgB {
				return 1
			} else if avgA < avgB {
				return -1
			}
			return 0
		})
		fmt.Println("✓ Sorted by average (lowest first)")
	default:
		fmt.Println("Invalid choice.")
	}
}

// InteractiveSearchStudent searches for students
func (gm *GradeManager) InteractiveSearchStudent(reader *bufio.Reader) {
	fmt.Print("\nEnter name to search: ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	queryLower := strings.ToLower(query)
	var found []Student

	for _, student := range gm.Students {
		if strings.Contains(strings.ToLower(student.Name), queryLower) {
			found = append(found, student)
		}
	}

	if len(found) == 0 {
		fmt.Printf("No students found matching '%s'\n", query)
		return
	}

	fmt.Printf("\nFound %d student(s):\n", len(found))
	for _, student := range found {
		avg := calculateAverage(student.Grades)
		fmt.Printf("  - %s (Average: %.2f, Grade: %s)\n",
			student.Name, avg, getLetterGrade(avg))
	}
}

// Helper function: calculate average of a slice
func calculateAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Helper function: convert score to letter grade
func getLetterGrade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

// Helper function: truncate string safely (Unicode-aware)
func truncateString(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen-1]) + "…"
}

// TO RUN: go run day4/07_challenge.go
//
// This program demonstrates:
// - Slices for storing students (dynamic array)
// - 2D structure: each student has a slice of grades
// - Slice operations: append, remove, sort, search
// - String handling: case-insensitive search, truncation
// - Using the slices package for sorting
//
// BONUS CHALLENGES:
// 1. Add ability to add/remove subjects
// 2. Save/load data to/from a file
// 3. Add a "curve grades" feature that adjusts all grades
// 4. Implement grade weighting (e.g., finals worth more)
// 5. Add input validation to prevent duplicate student names
//
// CONCEPTS PRACTICED:
// - Arrays vs Slices
// - Slice capacity and growth
// - Slice manipulation (append, remove)
// - 2D slices
// - String operations
// - Rune-safe string handling
// - Sorting with custom comparators
