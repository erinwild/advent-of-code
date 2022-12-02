package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var lenInput int
	positionCounts := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lenInput++
		number := scanner.Text()
		for position, value := range strings.Split(number, "") {
			if value == "1" {
				positionCounts[position] += 1
			}
		}
	}

	var mostCommon int
	var leastCommon int

	for position := 0; position < len(positionCounts); position++ {
		value := positionCounts[position]
		if value > (lenInput - value) {
			mostCommon += 1
		} else {
			leastCommon += 1
		}
		mostCommon = mostCommon << 1
		leastCommon = leastCommon << 1
	}
	mostCommon = mostCommon >> 1
	leastCommon = leastCommon >> 1

	fmt.Println(mostCommon)
	fmt.Println(leastCommon)
	fmt.Println(mostCommon * leastCommon)
}
