package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var leftList, rightList []int

	leftList, rightList, err = scanInputFile(file)
	if err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	fmt.Printf("Read %d pairs of numbers\n", len(leftList))

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	leftCounter, rightCounter := getDataStructures(leftList, rightList)

	totalDistance := calculateTotalDistance(leftList, rightList)

	similarityScore := calculateSimilarityScore(leftList, leftCounter, rightCounter)

	fmt.Printf("Total distance between the lists: %d\n", totalDistance)
	fmt.Printf("Similarity Score: %d\n", similarityScore)
}

func getDataStructures(leftList []int, rightList []int) (map[int]int, map[int]int) {
	leftCounter := make(map[int]int)
	rightCounter := make(map[int]int)

	for i := 0; i < len(leftList); i++ {
		leftCounter[leftList[i]]++
		rightCounter[rightList[i]]++
	}

	return leftCounter, rightCounter
}

func calculateSimilarityScore(leftList []int, leftCounter map[int]int, rightCounter map[int]int) int {
	similarityScore := 0

	for leftNum := range leftCounter {
		if rightCount, exists := rightCounter[leftNum]; exists {
			similarityScore += leftNum * rightCount
		}
	}

	return similarityScore
}

func calculateTotalDistance(leftList []int, rightList []int) int {
	// Calculate total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}
	return totalDistance
}

// scanInputFile scans the input file and returns two slices of integers
func scanInputFile(file *os.File) ([]int, []int, error) {
	var leftList, rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Split the line by whitespace
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		// Parse left number
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error parsing left number: %v\n", err)
			continue
		}

		// Parse right number
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error parsing right number: %v\n", err)
			continue
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return leftList, rightList, nil
}

// openInputFile opens and returns a file handle for the given filename
func openInputFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return file, nil
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
