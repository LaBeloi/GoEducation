package main

import (
	"fmt"
	"math/rand"
)

type Animal struct {
	weight    float64
	kind      string
	feedPerKg float64
}

func (an Animal) calcFoodQuantity() float64 {
	return float64(an.weight) * float64(an.feedPerKg)
}

type foodQuantity interface {
	calcFoodQuantity() float64
}

func genFarm(quantityOfAnimals int) []Animal {
	// took from the internet
	var animalStats = map[string]map[string]float64{
		"dog": {
			"foodPerKg": 10,
			"minWeight": 5,
			"maxWeight": 70,
		},
		"cat": {
			"foodPerKg": 7,
			"minWeight": 4,
			"maxWeight": 8,
		},
		"cow": {
			"foodPerKg": 25,
			"minWeight": 720,
			"maxWeight": 1100,
		},
	}

	var result []Animal
	for i := 0; i < quantityOfAnimals; i++ {
		random := rand.Intn(3) // 0 - cow, 1 - cat, 2 - dog
		switch random {
		case 0:
			stats := animalStats["cow"]
			n := int(stats["maxWeight"]) - int(stats["minWeight"])
			weight := float64(rand.Intn(n)) + stats["minWeight"]
			result = append(result, Animal{feedPerKg: stats["foodPerKg"], kind: "Cow", weight: weight})
		case 1:
			stats := animalStats["cat"]
			n := int(stats["maxWeight"]) - int(stats["minWeight"])
			weight := float64(rand.Intn(n)) + stats["minWeight"]
			result = append(result, Animal{feedPerKg: stats["foodPerKg"], kind: "Cat", weight: weight})
		case 2:
			stats := animalStats["dog"]
			n := int(stats["maxWeight"]) - int(stats["minWeight"])
			weight := float64(rand.Intn(n)) + stats["minWeight"]
			result = append(result, Animal{feedPerKg: stats["foodPerKg"], kind: "Dog", weight: weight})
		default:
			stats := animalStats["cow"] // this is farm, so i fell to decision that cows quantity should be the biggest
			n := int(stats["maxWeight"]) - int(stats["minWeight"])
			weight := float64(rand.Intn(n)) + stats["minWeight"]
			result = append(result, Animal{feedPerKg: stats["foodPerKg"], kind: "Cow", weight: weight})
		}
	}
	return result
}

func main() {
	var totatFood float64
	farmAnimalsQuantity := 27
	listOfAnimals := genFarm(farmAnimalsQuantity)

	var inter foodQuantity

	for i, v := range listOfAnimals {
		inter = v
		foodForAnimal := inter.calcFoodQuantity()
		totatFood += foodForAnimal
		fmt.Printf("%v: For %v, with %v kg of weight, needs %v kg of food.\n", i+1, v.kind, v.weight, foodForAnimal)
	}
	fmt.Printf("Farm needs %v kg of food for animals.", totatFood)
}
