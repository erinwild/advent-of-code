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
	var data []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		lenData := len(data)
		data = append(data, measurement)
		if lenData > 1 {
			if data[lenData-1] > data[lenData-2] {
				count++
			}
		}
	}

	fmt.Println(count)
}
