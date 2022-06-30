package main

import (
	"fmt"
	"strings"
)

func byteCalc(s string) rune {
	var result rune
	for _, v := range s {
		result += v - 96
	}
	return result
}

func high(s string) string {
	var result rune
	var returnValue string
	slice := strings.Split(s, " ")
	for i, v := range slice {
		value := byteCalc(v)
		if i == 0 {
			result = value
			returnValue = v
			continue
		}
		if result < value {
			result = value
			returnValue = v
			continue
		}
	}

	fmt.Println(returnValue)
	return returnValue

}

func main() {
	high("take me to semynak")
}
