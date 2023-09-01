package main

import "testing"

func TestCalculateTax(t *testing.T) {
	// Test when the salary is less than $50,000
	salary := 40000.0
	tax := calculateTax(salary)
	if tax != 0 {
		t.Errorf("Expected 0 tax for a salary of $%.2f, but got: $%.2f", salary, tax)
	}

	// Test when the salary is more than $50,000
	salary = 60000.0
	tax = calculateTax(salary)
	if tax != salary*0.17 {
		t.Errorf("Expected $%.2f tax for a salary of $%.2f, but got: $%.2f", salary*0.17, salary, tax)
	}

	// Test when the salary is more than $150,000
	salary = 200000.0
	tax = calculateTax(salary)
	if tax != salary*0.27 {
		t.Errorf("Expected $%.2f tax for a salary of $%.2f, but got: $%.2f", salary*0.27, salary, tax)
	}
}
