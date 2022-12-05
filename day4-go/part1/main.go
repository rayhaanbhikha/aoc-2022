package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	lower, upper int
}

func NewRange(rawRange string) *Range {
	l, u, ok := strings.Cut(rawRange, "-")
	if !ok {
		log.Fatalf("can't parse pair %s", rawRange)
	}
	ln, _ := strconv.Atoi(l)
	un, _ := strconv.Atoi(u)
	return &Range{ln, un}
}

func (r *Range) IsSubset(otherRange *Range) bool {
	return otherRange.lower <= r.lower && otherRange.upper >= r.upper
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	overlaps := 0
	for _, input := range inputs {
		raw1, raw2, _ := strings.Cut(input, ",")
		r1, r2 := NewRange(raw1), NewRange(raw2)
		if r1.IsSubset(r2) || r2.IsSubset(r1) {
			overlaps++
		}
	}
	fmt.Println(overlaps)
}
