package main

import (
	"bufio"
	"fmt"
	"os"
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

	// Read and parse the input
	grid, err := scanInputFile(file)
	if err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Printf("Read grid with %d rows and %d columns\n", len(grid), len(grid[0]))

	// Find all instances of XMAS
	totalXMAS := findXMAS(grid)
	totalBigXMas := findBigXMas(grid)
	fmt.Printf("Total XMAS instances found: %d\n", totalXMAS)
	fmt.Printf("Total Big XMAS instances found: %d\n", totalBigXMas)
}

func findBigXMas(grid [][]string) int {
	total := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check each position as a potential starting point
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == "A" &&
				row-1 > -1 &&
				row+1 < rows &&
				col-1 > -1 &&
				col+1 < cols {
				total += checkBigXMas(grid, row, col)
			}
		}
	}
	return total
}

func checkBigXMas(grid [][]string, row int, col int) int {
	tl, tr := grid[row-1][col-1], grid[row-1][col+1]
	bl, br := grid[row+1][col-1], grid[row+1][col+1]

	diag1 := (tl == "M" && br == "S") || (tl == "S" && br == "M")
	diag2 := (tr == "M" && bl == "S") || (tr == "S" && bl == "M")

	if grid[row][col] == "A" &&
		diag1 && diag2 {
		return 1
	}
	return 0
}

// findXMAS finds all instances of "XMAS" in the grid
func findXMAS(grid [][]string) int {
	total := 0
	rows := len(grid)
	cols := len(grid[0])

	// Check each position as a potential starting point
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check all 8 directions from this position
			total += checkDirection(grid, row, col, -1, -1) // diagonal up-left
			total += checkDirection(grid, row, col, -1, 0)  // up
			total += checkDirection(grid, row, col, -1, 1)  // diagonal up-right
			total += checkDirection(grid, row, col, 0, -1)  // left
			total += checkDirection(grid, row, col, 0, 1)   // right
			total += checkDirection(grid, row, col, 1, -1)  // diagonal down-left
			total += checkDirection(grid, row, col, 1, 0)   // down
			total += checkDirection(grid, row, col, 1, 1)   // diagonal down-right
		}
	}

	return total
}

// checkDirection checks if "XMAS" appears starting from (row, col) in the given direction
func checkDirection(grid [][]string, startRow, startCol, dRow, dCol int) int {
	rows := len(grid)
	cols := len(grid[0])

	// Check if we can fit "XMAS" (4 characters) in this direction
	endRow := startRow + 3*dRow
	endCol := startCol + 3*dCol

	// Check bounds
	if endRow < 0 || endRow >= rows || endCol < 0 || endCol >= cols {
		return 0
	}

	// Check if the sequence spells "XMAS"
	if grid[startRow][startCol] == "X" &&
		grid[startRow+dRow][startCol+dCol] == "M" &&
		grid[startRow+2*dRow][startCol+2*dCol] == "A" &&
		grid[startRow+3*dRow][startCol+3*dCol] == "S" {
		return 1
	}

	return 0
}

// scanInputFile scans the input file and returns a 2D grid
func scanInputFile(file *os.File) ([][]string, error) {
	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line into individual characters
		var row []string
		for _, char := range line {
			row = append(row, string(char))
		}

		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return grid, nil
}

// openInputFile opens and returns a file handle for the given filename
func openInputFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return file, nil
}
