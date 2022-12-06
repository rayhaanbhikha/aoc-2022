package main

import (
	"fmt"
	"io/ioutil"
)

func uniqueChars(input string) bool {
	charsSeen := make(map[int32]struct{})
	for _, char := range input {
		if _, ok := charsSeen[char]; ok {
			return false
		}
		charsSeen[char] = struct{}{}
	}
	return true
}

func parse(input string) int {
	charSize := 14
	for i := charSize; i <= len(input); i++ {
		chars := input[i-charSize : i]
		if uniqueChars(chars) {
			return i
		}
	}
	return 0
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	fmt.Println(parse(string(data)))
}
