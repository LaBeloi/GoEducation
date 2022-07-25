package sliceFuncs

import "fmt"

func TotalOfSlice(slice []int, i int) int {
	var result int
	for _, v := range slice {
		result += v
	}
	fmt.Printf("slice %v: %v\n", i, result)
	return result
}
