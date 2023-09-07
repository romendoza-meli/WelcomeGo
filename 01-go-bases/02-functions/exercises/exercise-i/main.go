package main

import "fmt"

func calculateTax(salary float64) float64 {
	if salary > 150000 {
		return salary * 0.27
	} else if salary > 50000 {
		return salary * 0.17
	} else {
		return 0
	}
}

func main() {
	salary := 60000.0
	tax := calculateTax(salary)
	fmt.Printf("The tax for a salary of $%.2f is: $%.2f\n", salary, tax)
}
