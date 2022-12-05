package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	order     int // The index of the Elf as read from the input (static)
	calories  int // The number of calories the Elf is carrying
	heapIndex int // The index within the heap (changes)
}

type Heap []*Elf

func (elves Heap) Len() int {
	return len(elves)
}

func (elves Heap) Less(i, j int) bool {
	return elves[i].calories > elves[j].calories
}

func (elves Heap) Swap(i, j int) {
	elves[i], elves[j] = elves[j], elves[i]
	elves[i].heapIndex = i
	elves[j].heapIndex = j
}

func (elves *Heap) Push(e any) {
	n := len(*elves)
	elf := e.(*Elf)
	elf.heapIndex = n
	*elves = append(*elves, elf)
}

func (elves *Heap) Pop() any {
	old := *elves
	n := len(old)
	elf := old[n-1]
	old[n-1] = nil     // avoid memory leak
	elf.heapIndex = -1 // for safety
	*elves = old[0 : n-1]
	return elf
}

func (elves *Heap) addCalories(elf *Elf, calories int) {
	elf.calories += calories
	heap.Fix(elves, elf.heapIndex)
}

func createNewElf(order int) *Elf {
	return &Elf{
		order:    order,
		calories: 0,
	}
}

func main() {
	f, _ := os.Open("input1.txt")
	defer f.Close()

	order := 0
	elves := make(Heap, 0)
	currentElf := createNewElf(order)
	heap.Push(&elves, currentElf) // Initialize first elf

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		newElf := len(line) == 0
		if newElf {
			order++
			currentElf = createNewElf(order)
			heap.Push(&elves, currentElf)
		} else {
			calories, _ := strconv.Atoi(line)
			elves.addCalories(currentElf, calories)
		}
	}

	top := 3
	index := 0
	totalCalories := 0
	for index < top {
		elf := heap.Pop(&elves).(*Elf)
		totalCalories += elf.calories
		fmt.Printf("Elf %d is carrying %d calories\n", elf.order, elf.calories)
		index++
	}
	fmt.Printf("The top %d elves are carrying a total of %d calories\n", top, totalCalories)
}
