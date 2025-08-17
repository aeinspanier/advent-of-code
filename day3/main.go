package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	// Read and parse the input
	input, err := scanInputFile(file)
	if err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Printf("Read input with %d characters\n", len(input))

	regexStr := `mul\(\s*\d+\s*,\s*\d+\s*\)`
	mulStrings := getAllMulStrings(input, regexStr)

	fmt.Printf("Extracted %d mul(x,y) statements\n", len(mulStrings))

	total := 0

	for _, mulStmt := range mulStrings {
		result := multiply(mulStmt)
		total += result
	}

	fmt.Printf("Total of operations: %d\n", total)

	mulStringsConditional := getAllMulStrings(input, `(do\(\)|don't\(\)|mul\(\s*\d+\s*,\s*\d+\s*\))`)

	total2 := 0
	tally := true

	for _, stmt := range mulStringsConditional {
		if stmt == "do()" {
			tally = true
		} else if stmt == "don't()" {
			tally = false
		} else if tally {
			total2 += multiply(stmt)
		} else {
			//fmt.Printf("excluded stmt: %s\n", stmt)
		}
	}

	fmt.Printf("Total of do() operations: %d\n", total2)

}

func multiply(mulStmt string) int {
	pattern := regexp.MustCompile(`\d+`)
	intStrs := pattern.FindAllString(mulStmt, -1)

	num1, err := strconv.Atoi(intStrs[0])
	if err != nil {
		return 0 // or handle error
	}

	num2, err := strconv.Atoi(intStrs[1])
	if err != nil {
		return 0 // or handle error
	}

	return num1 * num2
}

func getAllMulStrings(input string, rgxString string) []string {
	pattern := regexp.MustCompile(rgxString)
	return pattern.FindAllString(input, -1)
}

// scanInputFile scans the input file and returns the content as a string
func scanInputFile(file *os.File) (string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Join all lines into a single string
	return strings.Join(lines, ""), nil
}

// openInputFile opens and returns a file handle for the given filename
func openInputFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return file, nil
}
