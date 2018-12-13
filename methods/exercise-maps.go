package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)
	for _, value := range strings.Fields(s) {
		_, ok := countMap[value]
		if ok {
			countMap[value] += 1
		} else {
			countMap[value] = 1
		}
		fmt.Println(countMap)
		return countMap
	}
}

func main() {
	wc.Test(WordCount)
}
