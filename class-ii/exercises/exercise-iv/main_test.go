package main

import (
	"testing"
)

func TestMinimum(t *testing.T) {
	minFunc, _ := operation(minimum)
	minValue, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	if minValue != 2 {
		t.Errorf("Expected 2, but got %.2f", minValue)
	}
}

func TestAverage(t *testing.T) {
	avgFunc, _ := operation(average)
	averageValue, _ := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if averageValue != 3 {
		t.Errorf("Expected 3, but got %.2f", averageValue)
	}
}

func TestMaximum(t *testing.T) {
	maxFunc, _ := operation(maximum)
	maxValue, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if maxValue != 5 {
		t.Errorf("Expected 5, but got %.2f", maxValue)
	}
}
