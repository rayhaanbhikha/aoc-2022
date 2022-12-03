package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func (r *Rucksack) UniqueItems() int {
	return len(r.compartments[0].items) + len(r.compartments[1].items)
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

func (r *Rucksack) Iterate() <-chan int32 {
	result := make(chan int32)
	go func() {
		defer close(result)
		for _, compartment := range r.compartments {
			for item := range compartment.items {
				result <- item
			}
		}
	}()
	return result
}

func findBadge(rucksacks []*Rucksack) (int32, bool) {
	sort.SliceStable(rucksacks, func(i, j int) bool {
		return rucksacks[i].UniqueItems() < rucksacks[j].UniqueItems()
	})

	for rucksackItem := range rucksacks[0].Iterate() {
		if rucksacks[1].Has(rucksackItem) && rucksacks[2].Has(rucksackItem) {
			return rucksackItem, true
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

	priorities := 0

	groupedInputs := make([][]*Rucksack, 0)

	for i := 0; i < len(inputs); i += 3 {
		n := i + 3
		group := make([]*Rucksack, 0)
		for _, rucksackInput := range inputs[i:n] {
			group = append(group, NewRuckSack(rucksackInput))
		}
		groupedInputs = append(groupedInputs, group)
	}

	for _, groupedInput := range groupedInputs {
		res, ok := findBadge(groupedInput)
		if !ok {
			panic("Badge not found")
		}
		priorities += MapToNum(res)
	}

	fmt.Println(priorities)
}
