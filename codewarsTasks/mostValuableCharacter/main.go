package main

// import (
// 	"fmt"
// )

func solve(s string) rune {
	var max int
	var st rune
	result := map[rune][]int{}
	for i, v := range s {
		item, ok := result[v]
		if !ok {
			result[v] = append(result[v], i)
			result[v] = append(result[v], i)
			if i == 0 {
				max = i - i
				st = v
			}
			continue
		}
		item[1] = i
		if max < item[1]-item[0] {
			max = item[1] - item[0]
			st = v
		}
	}

	for key, value := range result {
		num := value[1] - value[0]
		if max < num {
			max = num
			st = key
		}
		if max == num && !(key > st) {
			max = num
			st = key
		}
	}

	return st
}

func main() {
	solve("aabccc")
}
