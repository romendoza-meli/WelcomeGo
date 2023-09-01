package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	grades := []float64{90.0, -10.0, 85.0, 92.0, 77.0}
	expectedAverage := 86.0

	average := calculateAverage(grades...)
	if average != expectedAverage {
		t.Errorf("Expected average: %.2f, but got: %.2f", expectedAverage, average)
	}
}
