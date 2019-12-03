package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s []int
	s = []int{1, 2, 3}

	var input string
	fmt.Print("Please enter a sequence of integer: ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Printf("Invalid input %s\n", input)
		return
	}
	strs := strings.Fields(input)
	fmt.Println(strs)

	bubbleSort(s)

	fmt.Println(s)
}

func bubbleSort(unorderedSlice []int) {
	length := len(unorderedSlice)

	for i := 0; i < length - 1; i++ {
		for j := 1; j < length; j++ {
			if unorderedSlice[j-1] > unorderedSlice[j] {
				swap(unorderedSlice, j-1)
			}
		}
	}
}

func swap(s []int, index int) {
	//tmp := s[index]
	//s[index] = s[index+1]
	//s[index+1] = tmp
	s[index], s[index+1] = s[index+1], s[index]
}
