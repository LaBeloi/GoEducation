package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, 4, -4, 6, 3, 8, 8}
	var result []int

	if len(arr) == 1 {
		result = arr
	} else {
		mapOfIndexes := make(map[int]bool)
		sliceOfItems := []int{}

		for _, item := range arr {
			_, ok := mapOfIndexes[item]
			if !ok {
				mapOfIndexes[item] = true
				sliceOfItems = append(sliceOfItems, item)
			}
		}

		result = sliceOfItems
	}

	fmt.Println(result)
}
