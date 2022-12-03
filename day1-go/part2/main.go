package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type RichList[T any] []T

//func (r *RichList[T]) Fold(cb func(a, b T) any) any {
//	for i, v := range *r {
//		if i == len(*r)-1 {
//			break
//		}
//		(*r)[i+1] = cb(v, (*r)[i+1])
//	}
//	return (*r)[len(*r)-1]
//}

func (r *RichList[T]) For(cb func(index int, b T)) {
	for i, v := range *r {
		cb(i, v)
	}
}

func (r *RichList[T]) Map(cb func(index int, b T) (any, error)) (RichList[any], error) {
	newList := make(RichList[any], len(*r), len(*r))
	for i, v := range *r {
		newVal, err := cb(i, v)
		if err != nil {
			return nil, err
		}
		newList[i] = newVal
	}
	return newList, nil
}

//func (r *RichList[T]) Sort(func())

func main() {
	data, _ := ioutil.ReadFile("../input")
	elves := RichList[string](strings.Split(strings.TrimSpace(string(data)), "\n\n"))

	elvenCalories := make([]int, 0)

	elves.For(func(_ int, elfCalories string) {
		calories := RichList[string](strings.Split(elfCalories, "\n"))
		calorySum := 0
		for _, calory := range calories {
			ca, _ := strconv.Atoi(calory)
			calorySum += ca
		}
		elvenCalories = append(elvenCalories, calorySum)
	})

	sort.SliceStable(elvenCalories, func(i, j int) bool {
		return elvenCalories[i] > elvenCalories[j]
	})

	total := 0
	for _, val := range elvenCalories[:3] {
		total += val
	}

	fmt.Println(total)
}
