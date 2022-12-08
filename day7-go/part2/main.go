package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type FileType int

const (
	FILE FileType = iota
	DIR
)

type File struct {
	name     string
	size     int
	fileType FileType
}

func NewFile(name string, size int, fileType FileType) *File {
	return &File{name: name, size: size, fileType: fileType}
}

type Node struct {
	val      *File
	prev     *Node
	children []*Node
}

func (n *Node) Add(item *File) {
	node := &Node{val: item, prev: n}
	n.children = append(n.children, node)
}

func (n *Node) Next(name string) *Node {
	if name == ".." {
		if n.prev == nil {
			return n
		}
		return n.prev
	}
	for _, child := range n.children {
		if child.val.name == name {
			return child
		}
	}
	return n
}

func (n *Node) HasChildren() bool {
	return len(n.children) != 0
}

func (n *Node) Walk(cb func(n *Node)) {
	if !n.HasChildren() {
		return
	}
	for _, child := range n.children {
		child.Walk(cb)
	}
	cb(n)
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	rootNode := &Node{val: &File{name: "/", fileType: DIR}}
	currentNode := rootNode

	for _, input := range inputs {
		if input[0] == '$' {
			res := strings.Split(input[2:], " ")
			command := res[0]
			if command == "cd" {
				currentNode = currentNode.Next(res[1])
			}
			continue
		}
		p1, p2, ok := strings.Cut(input, " ")
		if !ok {
			continue
		}
		var file *File
		if p1 == "dir" {
			file = NewFile(p2, 0, DIR)
		} else {
			size, err := strconv.Atoi(p1)
			if err != nil {
				panic(err)
			}
			file = NewFile(p2, size, FILE)
		}
		currentNode.Add(file)
	}

	rootNode.Walk(func(n *Node) {
		if n.val.fileType == FILE {
			return
		}
		n.val.size = 0
		for _, child := range n.children {
			n.val.size += child.val.size
		}
	})

	totalDiskSpace := 70000000
	minimumUnusedSpace := 30000000

	spaceLeftOnDevice := totalDiskSpace - rootNode.val.size
	spaceThatNeedsDeleting := minimumUnusedSpace - spaceLeftOnDevice

	candidates := make([]*Node, 0)
	rootNode.Walk(func(n *Node) {
		if n.val.fileType == DIR && n.val.size >= spaceThatNeedsDeleting {
			candidates = append(candidates, n)
		}
	})

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].val.size < candidates[j].val.size
	})

	fmt.Println(candidates[0].val.size)
}
