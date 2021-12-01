package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var count int
	var slidingWindow []int
	var slidingSum []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		slidingWindow = append(slidingWindow, measurement)
		if len(slidingWindow) > 3 {
			slidingWindow = slidingWindow[1:4]
		}
		if len(slidingWindow) == 3 {
			sum := 0
			for _, v := range slidingWindow {
				sum += v
			}
			slidingSum = append(slidingSum, sum)
		}
		lenData := len(slidingSum)
		if lenData > 1 {
			if slidingSum[lenData-1] > slidingSum[lenData-2] {
				count++
			}
		}
	}

	fmt.Println(count)
}
