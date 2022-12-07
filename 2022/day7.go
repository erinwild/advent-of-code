package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name           string
	parent         *Directory
	subdirectories map[string]*Directory
	size           int
}

func printDirectory(root *Directory, fromRoot int) {
	// Prints a graphical representation of the tree of directories
	fmt.Printf("%s (size=%d)\n", root.name, root.size)
	for _, dir := range root.subdirectories {
		fmt.Print(strings.Repeat("  ", fromRoot+1))
		printDirectory(dir, fromRoot+1)
	}
}

func sumSignificantDirs(root *Directory, sizeThreshold int) int {
	// Returns the sum of the directory sizes under the provided threshold
	dirSize := root.size
	if dirSize > sizeThreshold {
		dirSize = 0
	}
	var subdirSizes int
	for _, dir := range root.subdirectories {
		subdirSizes += sumSignificantDirs(dir, sizeThreshold)
	}
	return dirSize + subdirSizes
}

func findClosestSubdirectory(root *Directory, desiredSize int) *Directory {
	// Returns the subdirectory closest in size to desiredSize without going under
	// Returns the given root directory if no subdirectory satisfies this
	closestDir := root
	for _, dir := range root.subdirectories {
		if dir.size >= desiredSize {
			closestSubdir := findClosestSubdirectory(dir, desiredSize)
			if closestSubdir.size <= closestDir.size {
				closestDir = closestSubdir
			}
		}
	}
	return closestDir
}

func main() {
	f, _ := os.Open("input7.txt")
	defer f.Close()

	filesystem := Directory{name: "/", subdirectories: make(map[string]*Directory)}
	var currentDir *Directory

	TotalDiskSpace := 70000000
	SpaceRequired := 30000000

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if strings.HasPrefix(line, "$ cd") {
			dir := tokens[2]
			if dir == ".." {
				currentDir = (*currentDir).parent
			} else if dir == "/" {
				currentDir = &filesystem
			} else {
				currentDir = (*&currentDir).subdirectories[dir]
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "dir") {
			subdirectoryName := tokens[1]
			subdirectory := Directory{
				name:           subdirectoryName,
				parent:         currentDir,
				subdirectories: make(map[string]*Directory),
			}
			(*currentDir).subdirectories[subdirectoryName] = &subdirectory
		} else {
			filesize, _ := strconv.Atoi(tokens[0])
			(*currentDir).size += filesize
			parentDir := (*currentDir).parent
			for parentDir != nil {
				(*parentDir).size += filesize
				parentDir = (*parentDir).parent
			}
		}
	}
	// Part 1
	// printDirectory(&filesystem, 0)
	fmt.Printf("Sum of directory sizes over 100000: %d\n", sumSignificantDirs(&filesystem, 100000))

	// Part 2
	totalUsedSpace := filesystem.size
	spaceToFree := SpaceRequired - (TotalDiskSpace - totalUsedSpace)
	fmt.Printf("%d space currently in use; need to delete %d space for the update\n", totalUsedSpace, spaceToFree)

	directoryToDelete := findClosestSubdirectory(&filesystem, spaceToFree)
	fmt.Printf("Delete directory %s to free up %d space\n", directoryToDelete.name, directoryToDelete.size)
}
