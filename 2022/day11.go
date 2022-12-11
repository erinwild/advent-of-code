package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation func(old int) int

type Monkey struct {
	index            int
	holdingItems     []int
	operation        Operation
	testDivisibleBy  int
	ifTrue           int // index of monkey to throw to
	ifFalse          int // index of monkey to throw to
	countInspections int // sum of number of times monkey has held an item
	commonMultiple   int
}

func (m *Monkey) ThrowItems(monkeys []*Monkey) {
	for _, item := range m.holdingItems {
		// Part 1
		// inspectedItem := getRelief(m.operation(item))
		// Part 2
		inspectedItem := m.operation(item)
		m.countInspections++
		var throwToIndex int
		if isDivisible(inspectedItem, m.testDivisibleBy) {
			throwToIndex = m.ifTrue
		} else {
			throwToIndex = m.ifFalse
		}
		throwTo := getMonkeyByIndex(monkeys, throwToIndex)
		// Part 1
		// throwTo.holdingItems = append(throwTo.holdingItems, inspectedItem)
		// Part 2
		throwTo.holdingItems = append(throwTo.holdingItems, inspectedItem%m.commonMultiple)
	}
	m.holdingItems = []int{}
}

func getMonkeyByIndex(monkeys []*Monkey, index int) *Monkey {
	return monkeys[index]
}

func isDivisible(worry, divisor int) bool {
	return worry%divisor == 0
}

func getRelief(worry int) int {
	return worry / 3
}

func main() {
	f, _ := os.Open("input11.txt")
	defer f.Close()

	var monkeyIndex int
	var lineIndex int
	var monkeys []*Monkey

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineIndex++

		if lineIndex == 1 {
			newMonkey := Monkey{index: monkeyIndex}
			monkeys = append(monkeys, &newMonkey)
		}
		currentMonkey := getMonkeyByIndex(monkeys, monkeyIndex)

		switch lineIndex {
		default:
			continue
		case 2:
			startingItems := strings.Split(strings.TrimPrefix(line, "  Starting items: "), ", ")
			holdingItems := make([]int, len(startingItems))
			for i, si := range startingItems {
				holdingItems[i], _ = strconv.Atoi(si)
			}
			currentMonkey.holdingItems = holdingItems
		case 3:
			operation := strings.Split(strings.TrimPrefix(line, "  Operation: new = old "), " ")
			operator := operation[0]

			var function Operation
			if operation[1] == "old" {
				switch operator {
				default:
					fmt.Println("operator not recognized")
				case "*":
					function = func(old int) int { return old * old }
				case "+":
					function = func(old int) int { return old + old }
				}
			} else {
				operatorValue, _ := strconv.Atoi(operation[1])
				switch operator {
				default:
					fmt.Println("operator not recognized")
				case "*":
					function = func(old int) int { return old * operatorValue }
				case "+":
					function = func(old int) int { return old + operatorValue }
				}
			}
			currentMonkey.operation = function
		case 4:
			test, _ := strconv.Atoi(strings.TrimPrefix(line, "  Test: divisible by "))
			currentMonkey.testDivisibleBy = test
		case 5:
			monkeyIfTrue, _ := strconv.Atoi(strings.TrimPrefix(line, "    If true: throw to monkey "))
			currentMonkey.ifTrue = monkeyIfTrue
		case 6:
			monkeyIfFalse, _ := strconv.Atoi(strings.TrimPrefix(line, "    If false: throw to monkey "))
			currentMonkey.ifFalse = monkeyIfFalse
		case 7:
			// empty line
			monkeyIndex++
			lineIndex = 0
		}

	}
	commonMultiple := 1
	for _, monkey := range monkeys {
		commonMultiple *= monkey.testDivisibleBy
	}
	for _, monkey := range monkeys {
		monkey.commonMultiple = commonMultiple
	}
	for round := 1; round <= 20; round++ {
		for _, monkey := range monkeys {
			monkey.ThrowItems(monkeys)
		}
	}
	allCounts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		allCounts[i] = monkey.countInspections
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allCounts)))
	monkeyBusiness := allCounts[0] * allCounts[1]
	fmt.Println(monkeyBusiness)
}
