package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	ftype    string
	name     string
	size     int
	children map[string]*Node
	parent   *Node
}

func newNode(t string, n string, s int, p *Node) *Node {
	return &Node{
		ftype:    t,
		name:     n,
		size:     s,
		children: make(map[string]*Node),
		parent:   p,
	}
}

// builds the filesystem tree and returns the root node, and a slice of all directory nodes.
func buildDirTree(input string) (*Node, []*Node) {
	var dirs []*Node

	fileExp := regexp.MustCompile(`(\d+) (.*)`)
	cmdExp := regexp.MustCompile(`\$ (\w+)(?: (.*)|$)`)

	root := newNode("dir", "/", 0, nil)
	var currentNode *Node = root
	for _, line := range strings.Split(string(input), "\n")[1:] {

		if string(line[0]) == "$" {
			if line[2:4] == "cd" {
				cmdGroups := cmdExp.FindStringSubmatch(line)
				_, arg := cmdGroups[1], cmdGroups[2]

				if arg == ".." {
					currentNode = currentNode.parent
				} else {
					node := newNode("dir", arg, 0, currentNode)
					currentNode.children[arg] = node
					currentNode = node

					dirs = append(dirs, node)
				}
			}

		} else if line[:3] != "dir" {
			fileGroups := fileExp.FindStringSubmatch(line)
			size, _ := strconv.Atoi(fileGroups[1])
			name := fileGroups[2]

			node := newNode("file", name, size, currentNode)
			currentNode.children[name] = node
		}
	}

	return root, dirs
}

func calcDirSize(node *Node) int {
	for _, child := range node.children {
		node.size += calcDirSize(child)
	}

	return node.size
}

func printTree(node *Node, depth int) {
	fmt.Printf(
		"%s - %s (%s, size=%d)\n",
		strings.Repeat(" ", depth*4),
		node.name,
		node.ftype,
		node.size,
	)
	for _, v := range node.children {
		printTree(v, depth+1)
	}
}

func main() {
	inputFile := os.Args[1]
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to error: %s \n", err)
	}

	root, dirs := buildDirTree(string(file))
	calcDirSize(root)
	printTree(root, 0)

	// part 1
	dirSize := 0
	for _, node := range dirs {
		if node.size <= 100000 {
			dirSize += node.size
		}
	}
	fmt.Printf("\ncombined size of all dirs <= 100000: %d\n", dirSize)

}
