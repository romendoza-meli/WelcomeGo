package main

import (
	"testing"
)

func TestCalculateSalary_CategoryA(t *testing.T) {
	hoursWorked := 160
	category := CategoryA
	expectedSalary := float64(hoursWorked) * 3000 * 1.5

	salary := calculateSalary(hoursWorked, category)
	if salary != expectedSalary {
		t.Errorf("Expected salary: $%.2f, but got: $%.2f", expectedSalary, salary)
	}
}

func TestCalculateSalary_CategoryB(t *testing.T) {
	hoursWorked := 160
	category := CategoryB
	expectedSalary := float64(hoursWorked) * 1500 * 1.2

	salary := calculateSalary(hoursWorked, category)
	if salary != expectedSalary {
		t.Errorf("Expected salary: $%.2f, but got: $%.2f", expectedSalary, salary)
	}
}

func TestCalculateSalary_CategoryC(t *testing.T) {
	hoursWorked := 160
	category := CategoryC
	expectedSalary := float64(hoursWorked) * 1000

	salary := calculateSalary(hoursWorked, category)
	if salary != expectedSalary {
		t.Errorf("Expected salary: $%.2f, but got: $%.2f", expectedSalary, salary)
	}
}
