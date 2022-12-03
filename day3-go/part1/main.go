package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func MapToNum(item int32) int {
	// a through z ascii
	if item <= 122 && item >= 97 {
		return int(item) - 96
	}
	return int(item) - 38
}

type Compartment struct {
	itemsSeen int
	items     map[int32]struct{}
}

func NewCompartment() *Compartment {
	return &Compartment{
		items: map[int32]struct{}{},
	}
}

func (c *Compartment) Add(item int32) {
	c.itemsSeen++
	c.items[item] = struct{}{}
}

func (c *Compartment) Has(item int32) bool {
	_, ok := c.items[item]
	return ok
}

func (c *Compartment) ItemsAdded() int {
	return c.itemsSeen
}

type Rucksack struct {
	compartmentSize int
	compartments    [2]*Compartment
}

func NewRuckSack(items string) *Rucksack {
	r := &Rucksack{
		compartmentSize: len(items) / 2,
		compartments: [2]*Compartment{
			NewCompartment(),
			NewCompartment(),
		},
	}
	for _, item := range items {
		r.Add(item)
	}
	return r
}

func (r *Rucksack) Has(item int32) bool {
	return r.compartments[0].Has(item) || r.compartments[1].Has(item)
}

func (r *Rucksack) Add(item int32) {
	index := 1
	if r.compartments[0].ItemsAdded() < r.compartmentSize {
		index = 0
	}

	r.compartments[index].Add(item)
}

func (r *Rucksack) HasDuplicate() (int, bool) {
	m := r.compartments[0]
	for item := range m.items {
		if r.compartments[1].Has(item) {
			return MapToNum(item), true
		}
	}

	return 0, false
}

func main() {
	data, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	matchedItems := 0

	for _, input := range inputs {
		r := NewRuckSack(input)
		matchedItem, hasMatch := r.HasDuplicate()
		if !hasMatch {
			continue
		}
		matchedItems += matchedItem
	}

	fmt.Println(matchedItems)
}
