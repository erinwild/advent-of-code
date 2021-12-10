package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// numSegmentsNeeded := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	// segmentsNeeded := [10][7]int{
	// 	{1, 1, 1, 0, 1, 1, 1},
	// 	{0, 0, 1, 0, 0, 1, 0},
	// 	{1, 0, 1, 1, 1, 0, 1},
	// 	{1, 0, 1, 1, 0, 1, 1},
	// 	{0, 1, 1, 1, 0, 1, 0},
	// 	{1, 1, 0, 1, 0, 1, 1},
	// 	{1, 1, 0, 1, 1, 1, 1},
	// 	{1, 0, 1, 0, 0, 1, 0},
	// 	{1, 1, 1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 0, 1, 1},
	// }

	f, _ := os.Open("input.txt")
	defer f.Close()

	var countUnique int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		//pattern := strings.Split(line[0], " ")
		output := strings.Split(line[1], " ")

		for _, o := range output {
			lenOutput := len(o)
			if lenOutput == 2 || lenOutput == 3 || lenOutput == 4 || lenOutput == 7 {
				countUnique++
			}
		}
	}

	fmt.Println(countUnique)
}
