package main

import (
	"fmt"
	"strings"
)

func countSheep(quantity int) string {
	slice := []string{}

	for i := 1; i <= quantity; i++ {
		str := fmt.Sprintf("%v sheep...", i)
		slice = append(slice, str)
	}

	return strings.Join(slice, "")
}

func main() {
	fmt.Println(countSheep(5))
}
