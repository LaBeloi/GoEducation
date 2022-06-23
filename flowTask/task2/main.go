package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "-5 -9 -6 1000 -5 100"
	var result string

	arr := strings.Split(input, " ")
	mapResult := map[string]int32{"max": 0, "min": 0}

	for index, value := range arr {
		intValue, err := strconv.Atoi(value)
		transformedInt := int32(intValue)

		if err != nil {
			log.Fatal(err)
		}

		if index == 0 {
			mapResult["max"] = transformedInt
			mapResult["min"] = transformedInt
			continue
		}

		if transformedInt > mapResult["max"] {
			mapResult["max"] = transformedInt
			continue
		}

		if transformedInt < mapResult["min"] {
			mapResult["min"] = transformedInt
			continue
		}
	}

	result = fmt.Sprintf("%v %v", mapResult["max"], mapResult["min"])
	println(result)
}
