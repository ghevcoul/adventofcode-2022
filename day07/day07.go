package day07

import (
	"fmt"
	"ghevcoul/aoc22/utils"
	"strconv"
	"strings"
)

/*
 * Build a basic Tree data structure to store the file system as we inspect it
 */
type Tree struct {
	root  *Node
	dirs  []*Node
	files []*Node
}

type Node struct {
	directory bool
	name      string
	size      int
	parent    *Node
	children  []*Node
}

func (node *Node) FindDir(name string) *Node {
	for _, child := range node.children {
		if child.name == name && child.directory {
			return child
		}
	}
	return nil
}

func Day7() {
	fmt.Println("* * * * * * * * * * Day 07 * * * * * * * * * *")
	dataPath := "inputs/day07"

	reader, err := utils.Readlines(dataPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tree := new(Tree)
	cwd := new(Node)
	cwd.directory = true
	cwd.name = "/"
	cwd.size = 0
	tree.root = cwd
	tree.dirs = append(tree.dirs, cwd)

	for line := range reader {
		splitLine := strings.Split(line, " ")
		// Handle commands, which start with $
		if splitLine[0] == "$" {
			switch splitLine[1] {
			case "cd":
				if splitLine[2] == "/" {
					cwd = tree.root
				} else if splitLine[2] == ".." {
					cwd = cwd.parent
				} else {
					cwd = cwd.FindDir(splitLine[2])
				}
			case "ls":
				// do nothing
			}
		} else if splitLine[0] == "dir" {
			// Create a new directory node
			dir := new(Node)
			dir.directory = true
			dir.name = splitLine[1]
			dir.size = 0
			dir.parent = cwd
			cwd.children = append(cwd.children, dir)
			tree.dirs = append(tree.dirs, dir)
		} else if splitLine[0] != "dir" {
			size, _ := strconv.Atoi(splitLine[0])
			// Create a new leaf node for a file
			file := new(Node)
			file.directory = false
			file.name = splitLine[1]
			file.size = size
			file.parent = cwd
			cwd.children = append(cwd.children, file)
			tree.files = append(tree.files, file)
		}
	}

	// Walk the files, adding up the sizes of each directory
	for _, file := range tree.files {
		dir := file.parent
		for dir != nil {
			dir.size += file.size
			dir = dir.parent
		}
	}
	runningTotalTree := 0
	for _, dir := range tree.dirs {
		if dir.size < 100000 {
			runningTotalTree += dir.size
		}
	}
	fmt.Println("The sum of all directories under 100,000 is", runningTotalTree)

	// Find the smallest directory that is larger than 30000000
	spaceRequired := 30000000
	diskSize := 70000000
	diskFree := diskSize - tree.root.size
	smallest := spaceRequired
	targetSize := smallest - diskFree
	for _, dir := range tree.dirs {
		if dir.size > targetSize {
			if dir.size < smallest {
				smallest = dir.size
			}
		}
	}
	fmt.Println("The size of the directory you should delete is", smallest)
	fmt.Println()
}
