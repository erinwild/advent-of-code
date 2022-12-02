package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var totalHorizontal int
	var totalDepth int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		amount, _ := strconv.Atoi(command[1])
		switch direction {
		case "up":
			totalDepth -= amount
		case "down":
			totalDepth += amount
		case "forward":
			totalHorizontal += amount
		case "back":
			totalHorizontal -= amount
		}
	}
	fmt.Println("total depth:", totalDepth)
	fmt.Println("total horizontal position:", totalHorizontal)
	fmt.Println(totalDepth * totalHorizontal)
}
