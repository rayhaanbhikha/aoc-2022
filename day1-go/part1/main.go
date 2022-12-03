package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type RichList[T any] []T

func (r *RichList[T]) Fold(cb func(a, b T) T) T {
	for i, v := range *r {
		if i == len(*r)-1 {
			break
		}
		(*r)[i+1] = cb(v, (*r)[i+1])
	}
	return (*r)[len(*r)-1]
}

func (r *RichList[T]) For(cb func(index int, b T)) {
	for i, v := range *r {
		cb(i, v)
	}
}

func (r *RichList[T]) Map(cb func(index int, b T) (any, error)) (RichList[any], error) {
	newList := make(RichList[any], 0, len(*r))
	for i, v := range *r {
		newVal, err := cb(i, v)
		if err != nil {
			return nil, err
		}
		newList[i] = newVal
	}
	return newList, nil
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	elves := RichList[string](strings.Split(strings.TrimSpace(string(data)), "\n\n"))
	maxCalories := 0

	elves.For(func(_ int, elfCalories string) {
		calories := strings.Split(elfCalories, "\n")

		calorySum := 0
		for _, calory := range calories {
			ca, _ := strconv.Atoi(calory)
			calorySum += ca
		}
		if maxCalories < calorySum {
			maxCalories = calorySum
		}
	})

	fmt.Println(maxCalories)
}
