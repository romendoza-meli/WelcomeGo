package main

import (
	"fmt"
)

func calculateAverage(grades ...int) float64 {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade < 0 {
			continue
		}
		sum += grade
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func main() {
	grades := []int{90, -10, 85, 92, 77}
	average := calculateAverage(grades...)
	fmt.Printf("The average grade is: %.2f\n", average)
}
