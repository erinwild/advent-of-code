package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func priority(item rune) int {
	// Returns "priority" score of character
	// a-z returns 1-26; A-Z returns 27-52
	uppercaseItem := unicode.ToUpper(item)
	lowercasePriority := int(uppercaseItem) - 64
	if item == uppercaseItem {
		return lowercasePriority + 26
	} else {
		return lowercasePriority
	}
}

func findCommonItem(rucksacks ...string) (rune, bool) {
	// Returns the common character across multiple strings
	// If only one string is provided, it is split into two equal halves for comparison
	numRucksacks := len(rucksacks)
	if numRucksacks == 1 {
		// Split single rucksack into two compartments for comparison
		rucksack := rucksacks[0]
		amountInRucksack := len(rucksack) / 2
		rucksacks = []string{rucksack[:amountInRucksack], rucksack[amountInRucksack:]}
		numRucksacks++
	}

	items := make(map[rune]int)
	for index, rucksack := range rucksacks {
		if index == 0 {
			for _, item := range rucksack {
				items[item] = 1
			}
		} else {
			for _, item := range rucksack {
				if items[item] == index {
					items[item] += 1
					if items[item] == numRucksacks {
						return item, true
					}
				}
			}
		}
	}
	return rune('0'), false
}

func main() {
	f, _ := os.Open("input3.txt")
	defer f.Close()

	rucksacks := []string{}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())

		// Change desired number of rucksacks to compare below
		if len(rucksacks) == 3 {
			if commonItem, ok := findCommonItem(rucksacks...); ok {
				sum += priority(commonItem)
			}
			// Reset rucksacks
			rucksacks = []string{}
		}
	}
	fmt.Println(sum)
}
