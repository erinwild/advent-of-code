package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (stack *Stack) Push(crates ...string) {
	for _, crate := range crates {
		*stack = append(*stack, crate)
	}
}

func (stack *Stack) Pop() string {
	lastIndex := len(*stack) - 1
	crate := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]
	return crate
}

func (stack *Stack) Peek() string {
	return (*stack)[len(*stack)-1]
}

func CrateMover9000(stacks *[]Stack, numToMove, from, to int) {
	// Part 1
	for i := 0; i < numToMove; i++ {
		crateToMove := (*stacks)[from-1].Pop()
		(*stacks)[to-1].Push(crateToMove)
	}
}

func CrateMover9001(stacks *[]Stack, numToMove, from, to int) {
	// Part 2
	var cratesToMove []string
	for i := 0; i < numToMove; i++ {
		cratesToMove = append(cratesToMove, (*stacks)[from-1].Pop())
	}
	for i := numToMove - 1; i >= 0; i-- {
		(*stacks)[to-1].Push(cratesToMove[i])
	}
}

func main() {
	f, _ := os.Open("input5.txt")
	defer f.Close()

	var isStartOfMoves bool
	var stacks []Stack

	// Manually initialize starting stacks
	startingCrates := [][]string{
		{"T", "P", "Z", "C", "S", "L", "Q", "N"},
		{"L", "P", "T", "V", "H", "C", "G"},
		{"D", "C", "Z", "F"},
		{"G", "W", "T", "D", "L", "M", "V", "C"},
		{"P", "W", "C"},
		{"P", "F", "J", "D", "C", "T", "S", "Z"},
		{"V", "W", "G", "B", "D"},
		{"N", "J", "S", "Q", "H", "W"},
		{"R", "C", "Q", "F", "S", "L", "V"},
	}
	for _, crates := range startingCrates {
		var newStack Stack
		newStack.Push(crates...)
		stacks = append(stacks, newStack)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			isStartOfMoves = true
			continue
		}
		if isStartOfMoves {
			// Move crates between stacks
			tokens := strings.Split(line, " ")
			numCratesToMove, _ := strconv.Atoi(tokens[1])
			fromStackIndex, _ := strconv.Atoi(tokens[3])
			toStackIndex, _ := strconv.Atoi(tokens[5])
			CrateMover9001(&stacks, numCratesToMove, fromStackIndex, toStackIndex)
		} else {
			// Skip parsing starting crates configuration because that looks hard :/
			continue
		}
	}
	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
	fmt.Println("")
}
