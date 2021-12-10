package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateFuel(crabs []int, position int) int {
	var totalFuel int
	for _, crab := range crabs {
		diff := crab - position
		if diff < 0 {
			totalFuel -= diff
		} else if diff > 0 {
			totalFuel += diff
		}
	}
	return totalFuel
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var inputText []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputText = strings.Split(scanner.Text(), ",")
	}
	var crabPositions []int
	var maxPosition int
	for _, c := range inputText {
		crab, _ := strconv.Atoi(c)
		crabPositions = append(crabPositions, crab)
		if crab > maxPosition {
			maxPosition = crab
		}
	}

	var hPos []int
	var bestPos int
	var lowestFuel int
	for i := 0; i <= maxPosition; i++ {
		fuel := calculateFuel(crabPositions, i)
		hPos = append(hPos, fuel)
		if fuel < hPos[bestPos] {
			bestPos = i
			lowestFuel = fuel
		}
	}
	fmt.Println(bestPos, lowestFuel)
}
