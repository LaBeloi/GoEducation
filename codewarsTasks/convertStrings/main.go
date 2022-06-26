package main

import (
	"fmt"
	"regexp"
	"strings"
)

func toCamelCase(s string) string {
	regex := regexp.MustCompile("[-,_]")
	result := regex.Split(s, -1)
	for i, v := range result {
		if i > 0 {
			result[i] = strings.Title(v)
		}
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(toCamelCase("Test-your_luck"))
}
