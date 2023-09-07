package main

import (
	"testing"
)

func TestDog(t *testing.T) {
	animalDog, _ := Animal(dog)
	dogFood := animalDog(10)
	if dogFood != 100 {
		t.Errorf("Expected 100, but got %.2f", dogFood)
	}
}

func TestCat(t *testing.T) {
	animalCat, _ := Animal(cat)
	catFood := animalCat(10)
	if catFood != 50 {
		t.Errorf("Expected 50, but got %.2f", catFood)
	}
}

func TestHamster(t *testing.T) {
	animalHamster, _ := Animal(hamster)
	hamsterFood := animalHamster(10)
	if hamsterFood != 2.5 {
		t.Errorf("Expected 2.5, but got %.2f", hamsterFood)
	}
}

func TestTarantula(t *testing.T) {
	animalTarantula, _ := Animal(tarantula)
	tarantulaFood := animalTarantula(10)
	if tarantulaFood != 1.5 {
		t.Errorf("Expected 1.5, but got %.2f", tarantulaFood)
	}
}
