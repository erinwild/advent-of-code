package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/RyanCarrier/dijkstra"
)

func getNeighbours(r, c, numRows, numCols int) (neighbours [][2]int) {
	if r > 0 && r < numRows-1 {
		neighbours = append(neighbours, [2]int{r - 1, c}, [2]int{r + 1, c})
	} else if r == 0 {
		neighbours = append(neighbours, [2]int{r + 1, c})
	} else if r == numRows-1 {
		neighbours = append(neighbours, [2]int{r - 1, c})
	}
	if c > 0 && c < numCols-1 {
		neighbours = append(neighbours, [2]int{r, c - 1}, [2]int{r, c + 1})
	} else if c == 0 {
		neighbours = append(neighbours, [2]int{r, c + 1})
	} else if c == numCols-1 {
		neighbours = append(neighbours, [2]int{r, c - 1})
	}
	return
}

func isValidMove(this, that rune) bool {
	if that == rune('E') {
		if this == rune('z') {
			return true
		} else {
			return false
		}
	} else if this == rune('S') {
		if that == rune('a') {
			return true
		} else {
			return false
		}
	} else if that > this {
		// check if one step away
		return (this+1 == that)
	} else {
		return true
	}
}

func main() {
	f, _ := os.Open("input12.txt")
	defer f.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		var row []rune
		for _, square := range line {
			row = append(row, square)
		}
		grid = append(grid, row)
	}

	numRows := len(grid)
	numCols := len(grid[0])
	g := dijkstra.NewGraph()
	for v := 0; v < numRows*numCols; v++ {
		g.AddVertex(v)
	}
	var start, end int
	var possibleStarts []int
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			vertex := row*numCols + col
			elevation := grid[row][col]
			if elevation == rune('S') {
				start = vertex
			} else if elevation == rune('E') {
				end = vertex
			} else if elevation == rune('a') {
				possibleStarts = append(possibleStarts, vertex)
			}
			neighbours := getNeighbours(row, col, numRows, numCols)
			for _, n := range neighbours {
				nRow, nCol := n[0], n[1]
				neighbourElevation := grid[nRow][nCol]
				if isValidMove(elevation, neighbourElevation) {
					neighbourVertex := nRow*numCols + nCol
					g.AddArc(vertex, neighbourVertex, 1)
				}
			}
		}
	}
	// Part 1
	path, _ := g.Shortest(start, end)
	originalPath := path
	// Part 2
	shortestPath := path.Distance
	for _, s := range possibleStarts {
		newPath, _ := g.Shortest(s, end)
		newPathDistance := newPath.Distance
		if (newPathDistance < shortestPath) && (newPathDistance > 0) {
			shortestPath = newPathDistance
			path = newPath
		}
	}
	fmt.Println("The distance from the original start is", originalPath.Distance)
	fmt.Println("The best distance from elevation a is", path.Distance)
}
