package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newDay(fishes []int) []int {
	newFishes := fishes
	for i, fish := range fishes {
		if fish >= 1 {
			newFishes[i]--
		} else if fish == 0 {
			newFishes[i] = 6
			newFishes = append(newFishes, 8)
		}
	}
	return newFishes
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var inputFish []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputFish = strings.Split(scanner.Text(), ",")
	}
	var lanternFish []int
	for _, fish := range inputFish {
		fish, _ := strconv.Atoi(fish)
		lanternFish = append(lanternFish, fish)
	}

	numDays := 80
	for n := 1; n <= numDays; n++ {
		lanternFish = newDay(lanternFish)
	}
	fmt.Println(len(lanternFish))
}
