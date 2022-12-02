package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var grid [][]string
	var riskLevel int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid = append(grid, line)
	}

	for r, row := range grid {
		for c, value := range row {
			height, _ := strconv.Atoi(value)

			up, down, left, right := -1, -1, -1, -1

			if r > 0 {
				up, _ = strconv.Atoi(grid[r-1][c])
			}
			if r < len(grid)-1 {
				down, _ = strconv.Atoi(grid[r+1][c])
			}
			if c > 0 {
				left, _ = strconv.Atoi(grid[r][c-1])
			}
			if c < len(row)-1 {
				right, _ = strconv.Atoi(grid[r][c+1])
			}
			if (up >= 0 && up <= height) ||
				(down >= 0 && down <= height) ||
				(left >= 0 && left <= height) ||
				(right >= 0 && right <= height) {
			} else {
				riskLevel += height + 1
			}
		}
	}
	fmt.Println(riskLevel)
}
