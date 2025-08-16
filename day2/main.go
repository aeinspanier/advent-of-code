package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	file, err := openInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read and parse the reports
	reports, err := scanInputFile(file)
	if err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Printf("Read %d reports\n", len(reports))

	// Count safe reports
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func isDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if diff >= 1 && diff <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff >= 1 && diff <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func isSafe(report []int) bool {
	// First check if the report is safe without any modifications
	if isIncreasing(report) || isDecreasing(report) {
		return true
	}

	// If not safe, find the problematic index and try removing it
	problemIndex := findProblemIndex(report)
	if problemIndex >= 0 {
		// Create a new slice with the problematic level removed
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:problemIndex]...)
		modifiedReport = append(modifiedReport, report[problemIndex+1:]...)

		// Check if removing the problematic level makes it safe
		return isIncreasing(modifiedReport) || isDecreasing(modifiedReport)
	}

	return false
}

// findProblemIndex finds the index of the level that's causing the report to be unsafe
func findProblemIndex(report []int) int {
	// Try removing each level one by one to see which one makes it safe
	for i := 0; i < len(report); i++ {
		// Create a slice without the i-th element
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		// Check if removing this level makes it safe
		if isIncreasing(modifiedReport) || isDecreasing(modifiedReport) {
			return i // This level can be removed to make it safe
		}
	}
	return -1 // No single level removal makes it safe
}

// scanInputFile scans the input file and returns a slice of reports
// Each report is a slice of integers representing levels
func scanInputFile(file *os.File) ([][]int, error) {
	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line by whitespace to get individual numbers
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// Parse each number in the report
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Error parsing number '%s': %v\n", part, err)
				continue
			}
			report = append(report, num)
		}

		if len(report) > 0 {
			reports = append(reports, report)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return reports, nil
}

// openInputFile opens and returns a file handle for the given filename
func openInputFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return file, nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
