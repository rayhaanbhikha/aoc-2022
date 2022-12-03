package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")
	winningTotal := 0
	for _, round := range inputs {
		p1, p2, _ := strings.Cut(round, " ")
		p1Score := getScore(p1)
		p2Score := getScore(p2)

		winningTotal += computeScore(p1Score, p2Score)
	}

	fmt.Println(winningTotal)

}

func computeScore(p1, p2 int) int {
	switch {
	case p2-p1 == 1 || p2 == 1 && p1 == 3:
		return p2 + 6
	case p1 == p2:
		return p2 + 3
	default:
		return p2 + 0
	}
}

func getScore(p string) int {
	switch p {
	case "A", "X": // rock
		return 1
	case "B", "Y": // paper
		return 2
	case "C", "Z": // scissors
		return 3
	default:
		log.Fatalf("invalid value of p: %s", p)
		return 0
	}
}
