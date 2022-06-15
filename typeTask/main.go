package main

import (
	"fmt"
)

var (
	apple  float64 = 5.99
	pear   int8    = 7
	budged int8    = 23
)

func defineAswer(param bool) string {
	if param {
		return "Так, можемо придбати!"
	}
	return "Ні, на жаль, ми бідні..."
}

func main() {
	fmt.Printf("Дано: \n Бюджет: %v \n Ціна за одиницю яблука: %v \n Ціна за одиницю груші: %v \n\n\n", budged, apple, pear)

	//first Task
	firstAnswer := apple*9 + float64(pear)*8
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? \n Відповідь: %v \n", firstAnswer)

	//second Task
	fmt.Printf("2. Скільки груш ми можемо купити? \n Відповідь: %v \n", budged/pear)

	//third Task
	thirdAswer := int(float64(budged) / apple)
	fmt.Printf("3. Скільки яблук ми можемо купити? \n Відповідь: %v \n", thirdAswer)

	// fourth Task
	sum := apple*2 + float64(pear)*2
	availability := float64(budged) >= sum

	fmt.Printf("4. Чи ми можемо купити 2 груші та 2 яблука? \n Відповідь: %v", defineAswer(availability))
}
