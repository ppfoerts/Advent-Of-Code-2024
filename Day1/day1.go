package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Get file input
	b, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	sliceLeft := []int{}
	sliceRight := []int{}

	scanner := bufio.NewScanner(b)

	// split into two halfs
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Fields(line)
		num0, err := strconv.Atoi(splitLine[0])
		if err != nil {
			log.Fatal(err)
		}
		sliceLeft = append(sliceLeft, num0)
		num1, err := strconv.Atoi(splitLine[1])
		if err != nil {
			log.Fatal(err)
		}
		sliceRight = append(sliceRight, num1)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sort by lowest to greats
	sort.Ints(sliceLeft)
	sort.Ints(sliceRight)

	fmt.Println(sliceLeft)
	fmt.Println(sliceRight)
	// get difference between the comparison of two lists
	differenceSlice := []int{}

	// Using range keyword
	for index, sliceLeftInt := range sliceLeft {
		var difference = absDiffInt(sliceRight[index], sliceLeftInt)
		differenceSlice = append(differenceSlice, difference)
	}
	fmt.Println(differenceSlice)
	// add together differences
	var totalDifference = 0
	for _, differenceSliceInt := range differenceSlice {
		totalDifference = totalDifference + differenceSliceInt
	}
	fmt.Println(totalDifference)

	// get similarit score (how many times the numbers on the left appear on the right)
	// make list of how many times a number appears
	m := make(map[int]int)

	for _, sliceRightInt := range sliceRight {
		value, ok := m[sliceRightInt]
		// If the key exists
		if ok {
			m[sliceRightInt] = value + 1
		} else {
			m[sliceRightInt] = 1
		}
	}

	fmt.Println("map:", m)

	var similarityScore = 0
	// calculate similarity score
	for _, sliceLeftInt := range sliceLeft {
		value, ok := m[sliceLeftInt]
		// If the key exists
		if ok {
			similarityScore = similarityScore + (value * sliceLeftInt)
		}
	}

	fmt.Println(similarityScore)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
