package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countSequentiallyShorter(treeHeight int, heightsToCompare []int) (int, bool) {
	var count int
	visible := true
	for _, otherTreeHeight := range heightsToCompare {
		if treeHeight > otherTreeHeight {
			count++
		} else {
			visible = false
			count++
			break
		}
	}
	return count, visible
}

func getSurroundingTrees(trees [][]int, row, col int) (left, right, top, bottom []int) {
	// Check left
	for otherCol := col - 1; otherCol >= 0; otherCol-- {
		left = append(left, trees[row][otherCol])
	}
	// Check right
	for otherCol := col + 1; otherCol < len(trees[0]); otherCol++ {
		right = append(right, trees[row][otherCol])
	}
	// Check top
	for otherRow := row - 1; otherRow >= 0; otherRow-- {
		top = append(top, trees[otherRow][col])
	}
	// Check bottom
	for otherRow := row + 1; otherRow < len(trees); otherRow++ {
		bottom = append(bottom, trees[otherRow][col])
	}

	return
}

func main() {
	f, _ := os.Open("input8.txt")
	defer f.Close()

	var trees [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		heights := make([]int, len(line))
		for col, val := range line {
			heights[col], _ = strconv.Atoi(string(val))
		}
		trees = append(trees, heights)
	}

	numRows := len(trees)
	numCols := len(trees[0])
	countVisible := (numRows + numCols - 2) * 2 // Outside edge
	highestScenicScore := 0
	for row := 1; row < numRows-1; row++ {
		for col := 1; col < numCols-1; col++ {
			currentTree := trees[row][col]
			treesLeft, treesRight, treesTop, treesBottom := getSurroundingTrees(trees, row, col)
			treesViewableLeft, visibleFromLeft := countSequentiallyShorter(currentTree, treesLeft)
			treesViewableRight, visibleFromRight := countSequentiallyShorter(currentTree, treesRight)
			treesViewableTop, visibleFromTop := countSequentiallyShorter(currentTree, treesTop)
			treesViewableBottom, visibleFromBottom := countSequentiallyShorter(currentTree, treesBottom)

			if visibleFromLeft || visibleFromRight || visibleFromTop || visibleFromBottom {
				countVisible++
			}

			scenicScore := treesViewableLeft * treesViewableRight * treesViewableTop * treesViewableBottom
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	fmt.Printf("%d trees are visible from outside the grid\n", countVisible)
	fmt.Println("The highest scenic score is", highestScenicScore)
}
