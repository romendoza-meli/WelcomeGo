package main

import (
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

type feedFunc func(int) float64

func Animal(animalType string) (feedFunc, error) {
	switch animalType {
	case dog:
		return func(quantity int) float64 {
			return 10 * float64(quantity)
		}, nil

	case cat:
		return func(quantity int) float64 {
			return 5 * float64(quantity)
		}, nil

	case hamster:
		return func(quantity int) float64 {
			return 0.250 * float64(quantity)
		}, nil

	case tarantula:
		return func(quantity int) float64 {
			return 0.150 * float64(quantity)
		}, nil

	default:
		return nil, fmt.Errorf("unknown animal type")
	}
}

func main() {
	var amount float64

	animalDog, err := Animal(dog)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount += animalDog(10)

	animalCat, err := Animal(cat)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount += animalCat(10)

	fmt.Printf("Total amount of food: %.2f Kg\n", amount)
}
