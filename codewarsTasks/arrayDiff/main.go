package main

import "fmt"

func arrDiff(a, b []int) []int {
	for _, v := range b {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == v {
				a = append(a[:i], a[i+1:]...)
			}
		}
	}
	return a
}

func main() {
	fmt.Println(arrDiff([]int{1, 2, 3}, []int{1, 2}))
}
