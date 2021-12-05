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
		points = append(points, coords)
	}
	return points
}

func main() {
	points := processInput()

	var oceanFloor [1000][1000]int // Assuming a 1000x1000 map. Should find max from points instead.
	var countOverlap int
	var xMin int
	var xMax int
	var yMin int
	var yMax int

	for _, point := range points {
		if point[0] == point[2] {
			if point[1] <= point[3] {
				yMin = point[1]
				yMax = point[3]
			} else {
				yMin = point[3]
				yMax = point[1]
			}
			for y := yMin; y <= yMax; y++ {
				oceanFloor[point[0]][y] += 1
				if oceanFloor[point[0]][y] == 2 {
					countOverlap++
				}
			}
		} else if point[1] == point[3] {
			if point[0] <= point[2] {
				xMin = point[0]
				xMax = point[2]
			} else {
				xMin = point[2]
				xMax = point[0]
			}
			for x := xMin; x <= xMax; x++ {
				oceanFloor[x][point[1]] += 1
				if oceanFloor[x][point[1]] == 2 {
					countOverlap++
				}
			}
		} else {
			if point[0] <= point[2] {
				xMin = point[0]
				xMax = point[2]
				yMin = point[1]
				yMax = point[3]
			} else {
				xMin = point[2]
				xMax = point[0]
				yMin = point[3]
				yMax = point[1]
			}
			if yMin <= yMax {
				for x, y := xMin, yMin; (x <= xMax) && (y <= yMax); x, y = x+1, y+1 {
					oceanFloor[x][y] += 1
					if oceanFloor[x][y] == 2 {
						countOverlap++
					}
				}
			} else {
				for x, y := xMin, yMin; (x <= xMax) && (y >= yMax); x, y = x+1, y-1 {
					oceanFloor[x][y] += 1
					if oceanFloor[x][y] == 2 {
						countOverlap++
					}
				}
			}
		}
	}

	fmt.Println(countOverlap)
}
