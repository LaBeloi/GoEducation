package main

import "fmt"

func recursion(quantity, num, step int, slice []int) []int {
	if len(slice) == quantity {
		return slice
	}
	newSlice := append(slice, num)
	return recursion(quantity, num+step, step, newSlice)
}

func test(n int) [][]int {
	result := [][]int{}
	for i := 1; i <= n; i++ {
		startSlice := []int{}
		result = append(result, recursion(n, i, i, startSlice))
	}
	return result
}

func main() {
	fmt.Println(test(3))
}
