package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newDay(fishes [9]int) [9]int {
	var newFishes [9]int
	for f := 0; f < 8; f++ {
		newFishes[f] = fishes[f+1]
	}
	newFishes[6] += fishes[0]
	newFishes[8] += fishes[0]
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
	var lanternFish [9]int
	for _, fish := range inputFish {
		fish, _ := strconv.Atoi(fish)
		lanternFish[fish]++
	}

	fmt.Println(lanternFish)
	numDays := 256
	for n := 1; n <= numDays; n++ {
		lanternFish = newDay(lanternFish)
	}
	var sum int
	for _, n := range lanternFish {
		sum += n
	}
	fmt.Println(sum)
}
