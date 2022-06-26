package main

import (
	"fmt"
)

func findSumans(n int) []int {
	sum := n * n * n
	start := sum / n
	result := []int{}
	if start%2 == 1 {
		for i := (-2 * (n / 2)); i <= (2 * (n / 2)); i += 2 {
			result = append(result, start+i)
		}
		return result
	}
	start -= 1
	for i := (-2 * (n/2 - 1)); i <= (2 * n / 2); i += 2 {
		result = append(result, start+i)
	}
	return result
}

func main() {
	fmt.Println(findSumans(6))
}
