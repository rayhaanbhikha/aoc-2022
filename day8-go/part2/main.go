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
	coord       *Coord
	size        int
	onBorder    bool
	scenicScore int
}

func NewTree(coord *Coord, size int, onBorder bool) *Tree {
	return &Tree{coord: coord, size: size, onBorder: onBorder}
}

func (t *Tree) Hash() string {
	return fmt.Sprintf("%d:%d", t.coord.x, t.coord.y)
}

type Predicate func(c *Coord) bool

func (t *Tree) CheckVisibilityInDirection(unitCoord *Coord, trees map[string]*Tree, reachedBoundary Predicate) int {
	treeCoordsToCheck := make([]*Coord, 0)
	for i := 1; ; i++ {
		coord := t.coord.Add(unitCoord, i)
		treeCoordsToCheck = append(treeCoordsToCheck, coord)
		if reachedBoundary(coord) {
			break
		}
	}

	treesSeen := 0

	for _, treeCoord := range treeCoordsToCheck {
		hash := fmt.Sprintf("%d:%d", treeCoord.x, treeCoord.y)
		neighbouringTree, ok := trees[hash]
		if !ok {
			continue
		}

		treesSeen++

		if neighbouringTree.size >= t.size {
			return treesSeen
		}
	}

	return treesSeen
}

func (t *Tree) ComputeVisibleTrees(trees map[string]*Tree, maxX, maxY int) int {

	directions := []struct {
		unitCoord *Coord
		predicate Predicate
	}{
		{
			unitCoord: &Coord{0, -1},
			predicate: func(c *Coord) bool {
				return c.y == 0
			},
		},
		{
			unitCoord: &Coord{0, 1},
			predicate: func(c *Coord) bool {
				return c.y == maxY
			},
		},
		{
			unitCoord: &Coord{1, 0},
			predicate: func(c *Coord) bool {
				return c.x == maxX
			},
		},
		{
			unitCoord: &Coord{-1, 0},
			predicate: func(c *Coord) bool {
				return c.x == 0
			},
		},
	}

	score := 1

	for _, direction := range directions {
		// TODO: cheat made the assumption that a tree on the border will not have the highest scenic score.
		// To properly implement this properly we need to filter out the zero size.
		if t.onBorder {
			continue
		}
		score *= t.CheckVisibilityInDirection(direction.unitCoord, trees, direction.predicate)
	}

	t.scenicScore = score

	return score
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

	highestScenicTreeScore := 0
	for _, tree := range trees {
		visible := tree.ComputeVisibleTrees(trees, maxX, maxY)
		if highestScenicTreeScore < visible {
			highestScenicTreeScore = visible
		}
	}
	fmt.Println(highestScenicTreeScore)

	//t, _ := trees["2:3"]
	//fmt.Println(t.ComputeVisibleTrees(trees, maxX, maxY))
}
