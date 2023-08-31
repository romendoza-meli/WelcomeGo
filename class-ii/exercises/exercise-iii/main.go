package main

import "fmt"

const (
	CategoryC = "C"
	CategoryB = "B"
	CategoryA = "A"
)

func calculateSalary(hoursWorked int, category string) float64 {
	var hourlyRate, bonusPercentage float64

	switch category {
	case CategoryC:
		hourlyRate = 1000
	case CategoryB:
		hourlyRate = 1500
		bonusPercentage = 0.2
	case CategoryA:
		hourlyRate = 3000
		bonusPercentage = 0.5
	default:
		return 0
	}

	salary := float64(hoursWorked) * hourlyRate
	if bonusPercentage > 0 {
		bonus := salary * bonusPercentage
		salary += bonus
	}

	return salary
}

func main() {
	hoursWorked := 160 // Assuming 160 hours worked per month
	category := CategoryB

	salary := calculateSalary(hoursWorked, category)
	fmt.Printf("The salary for %d hours worked in category %s is: $%.2f\n", hoursWorked, category, salary)
}
