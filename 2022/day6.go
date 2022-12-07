package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input6.txt")
	defer f.Close()

	marker := ""
	markerIndex := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		char := scanner.Text()
		index := strings.Index(marker, char)
		if index != -1 {
			marker = marker[index+1:]
		}
		marker += char
		markerIndex++

		if len(marker) == 14 {
			fmt.Println(marker)
			break
		}
	}
	fmt.Println(markerIndex)
}
