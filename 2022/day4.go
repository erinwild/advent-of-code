package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIfSubset(this, other [2]int) bool {
	// Checks if range 'other' is a subset of range 'this'
	if (other[0] >= this[0]) && (other[1] <= this[1]) {
		return true
	}
	return false
}

func checkOverlap(this, other [2]int) bool {
	// Checks if range 'other' overlaps range 'this'
	if (other[0] > this[1]) || (other[1] < this[0]) {
		return false
	}
	return true
}

func main() {
	f, _ := os.Open("input4.txt")
	defer f.Close()

	var countFullSubset int
	var countOverlap int
	var isSubset bool

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		splitLine := strings.Split(line, ",")[0:2]

		var assignmentPairs [][]string
		assignmentPairs = append(assignmentPairs, strings.Split(splitLine[0], "-"), strings.Split(splitLine[1], "-"))

		aOneStart, _ := strconv.Atoi(assignmentPairs[0][0])
		aOneEnd, _ := strconv.Atoi(assignmentPairs[0][1])
		assignmentOne := [2]int{aOneStart, aOneEnd}
		aTwoStart, _ := strconv.Atoi(assignmentPairs[1][0])
		aTwoEnd, _ := strconv.Atoi(assignmentPairs[1][1])
		assignmentTwo := [2]int{aTwoStart, aTwoEnd}

		// Count full subset
		if (aOneEnd - aOneStart) > (aTwoEnd - aTwoStart) {
			isSubset = checkIfSubset(assignmentOne, assignmentTwo)
		} else {
			isSubset = checkIfSubset(assignmentTwo, assignmentOne)
		}
		if isSubset {
			countFullSubset++
		}
		// Count overlap
		if checkOverlap(assignmentOne, assignmentTwo) {
			countOverlap++
		}
	}
	fmt.Printf("Num times one range fully contain the other: %d\n", countFullSubset)
	fmt.Printf("Num times one range overlaps the other at all: %d\n", countOverlap)
}
