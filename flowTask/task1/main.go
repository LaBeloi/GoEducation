package main

import (
	"fmt"
)

func deleteDublicates(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	mapOfIndexes := make(map[int]bool)
	sliceOfItems := []int{}

	for _, item := range slice {
		_, ok := mapOfIndexes[item]
		if !ok {
			mapOfIndexes[item] = true
			sliceOfItems = append(sliceOfItems, item)
		}
	}

	return sliceOfItems
}

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	result = deleteDublicates(arr)

	fmt.Println(result)
}
