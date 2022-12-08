package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var unitCoords = []*Coord{
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
	{0, -1}, // up
}

type Coord struct {
	x, y int
}

func (c *Coord) Add(oc *Coord, scale int) *Coord {
	return &Coord{x: c.x + (oc.x * scale), y: c.y + (oc.y * scale)}
}

func (c *Coord) neighbours(scale int) []*Coord {
	newCoords := make([]*Coord, 0)
	for _, unitCoord := range unitCoords {
		newCoords = append(newCoords, c.Add(unitCoord, scale))
	}
	return newCoords
}

type Tree struct {
	coord     *Coord
	size      int
	onBorder  bool
	isVisible bool
}

func NewTree(coord *Coord, size int, onBorder bool) *Tree {
	return &Tree{coord: coord, size: size, onBorder: onBorder}
}

func (t *Tree) Hash() string {
	return fmt.Sprintf("%d:%d", t.coord.x, t.coord.y)
}

func (t *Tree) CheckVisibilityInDirection(unitCoord *Coord, trees map[string]*Tree, reachedBoundary func(c *Coord) bool) bool {
	treeCoordsToCheck := make([]*Coord, 0)
	for i := 1; ; i++ {
		coord := t.coord.Add(unitCoord, i)
		treeCoordsToCheck = append(treeCoordsToCheck, coord)
		if reachedBoundary(coord) {
			break
		}
	}

	for _, treeCoord := range treeCoordsToCheck {
		hash := fmt.Sprintf("%d:%d", treeCoord.x, treeCoord.y)
		neighbouringTree, ok := trees[hash]
		if !ok {
			continue
		}

		if neighbouringTree.size >= t.size {
			return false
		}
	}

	return true
}

func (t *Tree) IsVisible(trees map[string]*Tree, maxX, maxY int) bool {

	// check up.
	visibleFromUp := t.CheckVisibilityInDirection(&Coord{0, -1}, trees, func(c *Coord) bool {
		return c.y == 0
	})

	if visibleFromUp {
		return true
	}

	// check down
	visibleFromDown := t.CheckVisibilityInDirection(&Coord{0, 1}, trees, func(c *Coord) bool {
		return c.y == maxY
	})

	if visibleFromDown {
		return true
	}

	// check right
	visibleFromRight := t.CheckVisibilityInDirection(&Coord{1, 0}, trees, func(c *Coord) bool {
		return c.x == maxX
	})

	if visibleFromRight {
		return true
	}
	// check right
	visibleFromLeft := t.CheckVisibilityInDirection(&Coord{-1, 0}, trees, func(c *Coord) bool {
		return c.x == 0
	})

	if visibleFromLeft {
		return true
	}

	return false
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	trees := make(map[string]*Tree)
	maxX := len(inputs[0]) - 1
	maxY := len(inputs) - 1

	for rowIndex, row := range inputs {
		for colIndex, col := range row {
			size, _ := strconv.Atoi(string(col))
			coord := &Coord{colIndex, rowIndex}
			onBorder := false
			if coord.x == 0 || coord.x == maxX || coord.y == 0 || coord.y == maxY {
				onBorder = true
			}

			tr := NewTree(coord, size, onBorder)
			trees[tr.Hash()] = tr
		}
	}

	visible := 0
	for _, tree := range trees {
		if tree.onBorder || tree.IsVisible(trees, maxX, maxY) {
			visible++
		}
	}

	fmt.Println(visible)
}
