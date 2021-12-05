package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func determineMostCommonBit(position int, numbers []string) (int, int) {
	var count int
	var mostCommonBit int
	var leastCommonBit int

	for _, number := range numbers {
		value := strings.Split(number, "")[position]
		if value == "1" {
			count++
		}
	}
	if count >= (len(numbers) - count) {
		mostCommonBit = 1
	} else {
		leastCommonBit = 1
	}
	return mostCommonBit, leastCommonBit
}

func filterByBit(position int, bit int, numbers []string) []string {
	var filtered []string
	for _, number := range numbers {
		value := strings.Split(number, "")[position]
		if value == fmt.Sprint(bit) {
			filtered = append(filtered, number)
		}
	}
	return filtered
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var binaryNumbers []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		number := scanner.Text()
		binaryNumbers = append(binaryNumbers, number)
	}

	var mostCommon int64
	var leastCommon int64
	var mostCommonBitInPos int
	var leastCommonBitInPos int
	mostCommonNumbers := make([]string, len(binaryNumbers))
	leastCommonNumbers := make([]string, len(binaryNumbers))
	copy(mostCommonNumbers, binaryNumbers)
	copy(leastCommonNumbers, binaryNumbers)

	for position := 0; position < len(binaryNumbers[0]); position++ {
		if len(mostCommonNumbers) <= 1 {
			continue
		} else {
			mostCommonBitInPos, _ = determineMostCommonBit(position, mostCommonNumbers)
			mostCommonNumbers = filterByBit(position, mostCommonBitInPos, mostCommonNumbers)
		}
		if len(leastCommonNumbers) == 1 {
			continue
		} else {
			_, leastCommonBitInPos = determineMostCommonBit(position, leastCommonNumbers)
			leastCommonNumbers = filterByBit(position, leastCommonBitInPos, leastCommonNumbers)
		}
	}

	mostCommon, _ = strconv.ParseInt(mostCommonNumbers[0], 2, 64)
	leastCommon, _ = strconv.ParseInt(leastCommonNumbers[0], 2, 64)

	fmt.Println(mostCommon)
	fmt.Println(leastCommon)
	fmt.Println(mostCommon * leastCommon)
}
