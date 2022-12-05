package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateScorePartOne(strategy string) int {
	switch strategy {
	case "A X": // Rock vs Rock
		return (1 + 3)
	case "A Y": // Rock vs Paper
		return (2 + 6)
	case "A Z": // Rock vs Scissors
		return (3 + 0)
	case "B X": // Paper vs Rock
		return (1 + 0)
	case "B Y": // Paper vs Paper
		return (2 + 3)
	case "B Z": // Paper vs Scissors
		return (3 + 6)
	case "C X": // Scissors vs Rock
		return (1 + 6)
	case "C Y": // Scissors vs Paper
		return (2 + 0)
	case "C Z": // Scissors vs Scissors
		return (3 + 3)
	default:
		fmt.Println("Unknown strategy")
		return 0
	}
}

func calculateScorePartTwo(strategy string) int {
	switch strategy {
	case "A X": // Rock vs ? = Lose
		return (3 + 0)
	case "A Y": // Rock vs ? = Draw
		return (1 + 3)
	case "A Z": // Rock vs ? = Win
		return (2 + 6)
	case "B X": // Paper vs ? = Lose
		return (1 + 0)
	case "B Y": // Paper vs ? = Draw
		return (2 + 3)
	case "B Z": // Paper vs ? = Win
		return (3 + 6)
	case "C X": // Scissors vs ? = Lose
		return (2 + 0)
	case "C Y": // Scissors vs ? = Draw
		return (3 + 3)
	case "C Z": // Scissors vs ? = Win
		return (1 + 6)
	default:
		fmt.Println("Unknown strategy")
		return 0
	}
}

func main() {
	f, _ := os.Open("input2.txt")
	defer f.Close()

	var partOneScore, partTwoScore int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		partOneScore += calculateScorePartOne(line)
		partTwoScore += calculateScorePartTwo(line)
	}

	fmt.Printf("You have %d points using strategy 1\n", partOneScore)
	fmt.Printf("You have %d points using strategy 2\n", partTwoScore)
}
