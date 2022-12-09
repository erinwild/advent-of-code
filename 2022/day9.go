package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moveInDirection(position [2]int, direction string) (newPosition [2]int) {
	newPosition = position
	switch direction {
	case "R":
		newPosition[0] += 1
	case "L":
		newPosition[0] -= 1
	case "U":
		newPosition[1] += 1
	case "D":
		newPosition[1] -= 1
	default:
		fmt.Println("not a valid direction")
	}
	return
}

func updateTail(newHead, tail [2]int, direction string) [2]int {
	// Given the new position of the head knot, return the tail knot's updated position
	diff := [2]int{newHead[0] - tail[0], newHead[1] - tail[1]}
	if (diff[0] == 2) || (diff[0] == -2) || (diff[1] == 2) || (diff[1] == -2) {
		if diff[0] < 0 {
			tail = moveInDirection(tail, "L")
		} else if diff[0] > 0 {
			tail = moveInDirection(tail, "R")
		}
		if diff[1] < 0 {
			tail = moveInDirection(tail, "D")
		} else if diff[1] > 0 {
			tail = moveInDirection(tail, "U")
		}
	}
	return tail
}

func parseMovement(move string) (direction string, distance int) {
	m := strings.Split(move, " ")
	direction = m[0]
	distance, _ = strconv.Atoi(m[1])
	return
}

func main() {
	f, _ := os.Open("input9.txt")
	defer f.Close()

	var knots [10][2]int
	numKnots := len(knots)
	visitedPositions := make(map[[2]int]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction, distance := parseMovement(line)
		for d := 0; d < distance; d++ {
			for i := 0; i < numKnots; i++ {
				if i == 0 {
					// Move head knot
					knots[i] = moveInDirection(knots[i], direction)
				} else {
					knots[i] = updateTail(knots[i-1], knots[i], direction)
				}
			}
			// Track positions of very last knot
			visitedPositions[knots[numKnots-1]] = true
		}
	}
	fmt.Printf("tail visited %d positions\n", len(visitedPositions))
}
