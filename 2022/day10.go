package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDesiredCycle(cycle int) bool {
	switch cycle {
	default:
		return false
	case 20, 60, 100, 140, 180, 220:
		return true
	}
}

func getSprite(registerVal map[int]int, cycle int) (sprite [3]int) {
	middle := registerVal[cycle]
	return [3]int{middle - 1, middle, middle + 1}
}

func printImage(registerVal map[int]int) {
	col := 0
	for pixel := 0; pixel < len(registerVal); pixel++ {
		sprite := getSprite(registerVal, pixel+1)
		if (col == sprite[0]) || (col == sprite[1]) || (col == sprite[2]) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		col++
		if col == 40 {
			fmt.Print("\n")
			col = 0
		}
	}
}

func main() {
	f, _ := os.Open("input10.txt")
	defer f.Close()

	cycle := 1
	registerVal := map[int]int{1: 1}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "addx") {
			addx, _ := strconv.Atoi(strings.TrimPrefix(line, "addx "))

			cycle++
			registerVal[cycle] = registerVal[cycle-1]

			cycle++
			registerVal[cycle] = registerVal[cycle-1] + addx
		} else {
			cycle++
			registerVal[cycle] = registerVal[cycle-1]
		}
	}

	sum := 0
	for c, rv := range registerVal {
		if isDesiredCycle(c) {
			fmt.Println("cycle", c, "*", "val", rv)
			sum += c * rv
		}
	}
	fmt.Println("The sum of the signal strengths is:", sum)
	printImage(registerVal)
}
