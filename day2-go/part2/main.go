package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")
	winningTotal := 0
	for _, round := range inputs {
		p1, p2, _ := strings.Cut(round, " ")
		winningTotal += getScore(p1, p2)
	}

	fmt.Println(winningTotal)
}

func getScore(p1, p2 string) int {
	player1, player2, score := 0, 0, 0
	switch p1 {
	case "A": // rock
		player1 = 0
	case "B": // paper
		player1 = 1
	case "C": // scissors
		player1 = 2
	}

	switch p2 {
	case "X":
		player2 = mod(player1-1, 3)
		score = 0
	case "Y":
		player2 = player1
		score = 3
	case "Z":
		player2 = mod(player1+1, 3)
		score = 6
	}

	return player2 + 1 + score
}
