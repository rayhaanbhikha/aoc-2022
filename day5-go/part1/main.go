package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{items: []string{}}
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}

func (s *Stack) Add(item ...string) {
	s.items = append(s.items, item...)
}

func (s *Stack) Push(item string) {
	s.items = append([]string{item}, s.items...)
}

func (s *Stack) PopN(n int) []string {
	if n < 0 {
		return []string{}
	}

	indexToDeleteFrom := len(s.items) - n
	if indexToDeleteFrom < 0 {
		indexToDeleteFrom = 0
	}

	itemsToReturn := s.items[indexToDeleteFrom:]
	s.items = s.items[:indexToDeleteFrom]

	return itemsToReturn
}

func parseStacks(rawStacks []string, expectedNumOfStacks int) []*Stack {
	stacks := make([]*Stack, 0)
	for i := 0; i < expectedNumOfStacks; i++ {
		stacks = append(stacks, NewStack())
	}
	for _, s := range rawStacks {
		fs := s[0:]
		charSize := 3
		stackIndex := 0
		for i := 0; i < len(fs); {
			char := fs[i : i+charSize]
			if char != "   " {
				stacks[stackIndex].Push(string(char[1]))
			}
			stackIndex++
			i++
			i += charSize
		}
	}
	return stacks
}

var re = regexp.MustCompile(`move (?P<NumItems>\d+) from (?P<FromStack>\d+) to (?P<ToStack>\d+)`)

type Instruction struct {
	From, To, N int
}

func (i *Instruction) String() string {
	return fmt.Sprintf("move %d from %d from %d", i.N, i.From, i.To)
}

func parseInstructions(instructions []string) []*Instruction {
	parsedInstructions := make([]*Instruction, 0)
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		res := re.FindStringSubmatch(instruction)
		n, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])

		inst := &Instruction{
			From: from - 1,
			To:   to - 1,
			N:    n,
		}
		parsedInstructions = append(parsedInstructions, inst)
	}
	return parsedInstructions
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(string(data), "\n")

	rawStacks := inputs[0:8]
	numOfStacks := len(strings.Split(strings.TrimSpace(inputs[8]), "  "))
	moveInstructions := inputs[10:]

	stacks := parseStacks(rawStacks, numOfStacks)
	instructions := parseInstructions(moveInstructions)

	for _, inst := range instructions {
		removed := stacks[inst.From].PopN(inst.N)
		for i, j := 0, len(removed)-1; i < j; i, j = i+1, j-1 {
			removed[i], removed[j] = removed[j], removed[i]
		}
		stacks[inst.To].Add(removed...)
	}

	result := ""
	for _, s := range stacks {
		result += s.Peek()
	}
	fmt.Println(result)
}
