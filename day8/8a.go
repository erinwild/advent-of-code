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

	var countUnique int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		output := strings.Split(line[1], " ")

		for _, o := range output {
			lenOutput := len(o)
			if lenOutput == 2 || lenOutput == 3 || lenOutput == 4 || lenOutput == 7 {
				countUnique++
			}
		}
	}

	fmt.Println(countUnique)
}
