package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInput() [][4]int {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var points [][4]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vent := scanner.Text()
		ventPoints := strings.Split(vent, " -> ")
		start := strings.Split(ventPoints[0], ",")
		end := strings.Split(ventPoints[1], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])

		coords := [4]int{startX, startY, endX, endY}

		// Consider only horizontal and vertical lines
		if (coords[0] == coords[2]) || (coords[1] == coords[3]) {
			points = append(points, coords)
		}
	}
	return points
}

func main() {
	points := processInput()

	var oceanFloor [1000][1000]int // Assuming a 1000x1000 map. Should find max from points instead.
	var countOverlap int
	var minPoint int
	var maxPoint int

	for _, point := range points {
		if point[0] == point[2] {
			if point[1] <= point[3] {
				minPoint = point[1]
				maxPoint = point[3]
			} else {
				minPoint = point[3]
				maxPoint = point[1]
			}
			for y := minPoint; y <= maxPoint; y++ {
				oceanFloor[point[0]][y] += 1
				if oceanFloor[point[0]][y] == 2 {
					countOverlap++
				}
			}
		} else if point[1] == point[3] {
			if point[0] <= point[2] {
				minPoint = point[0]
				maxPoint = point[2]
			} else {
				minPoint = point[2]
				maxPoint = point[0]
			}
			for x := minPoint; x <= maxPoint; x++ {
				oceanFloor[x][point[1]] += 1
				if oceanFloor[x][point[1]] == 2 {
					countOverlap++
				}
			}
		} else {
			fmt.Println("line should be horizontal or vertical")
		}
	}

	fmt.Println(countOverlap)
}
