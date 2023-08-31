package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

type operationFunc func(...int) (float64, error)

func operation(operationType string) (operationFunc, error) {
	switch operationType {
	case minimum:
		return func(numbers ...int) (float64, error) {
			if len(numbers) == 0 {
				return 0, fmt.Errorf("no numbers provided")
			}

			min := numbers[0]
			for _, num := range numbers[1:] {
				if num < min {
					min = num
				}
			}

			return float64(min), nil
		}, nil

	case average:
		return func(numbers ...int) (float64, error) {
			if len(numbers) == 0 {
				return 0, fmt.Errorf("no numbers provided")
			}

			sum := 0
			for _, num := range numbers {
				sum += num
			}

			average := float64(sum) / float64(len(numbers))
			return average, nil
		}, nil

	case maximum:
		return func(numbers ...int) (float64, error) {
			if len(numbers) == 0 {
				return 0, fmt.Errorf("no numbers provided")
			}

			max := numbers[0]
			for _, num := range numbers[1:] {
				if num > max {
					max = num
				}
			}

			return float64(max), nil
		}, nil

	default:
		return nil, fmt.Errorf("unknown operation type")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}
	minValue, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("Minimum value: %.2f\n", minValue)

	avgFunc, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}
	averageValue, _ := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Average value: %.2f\n", averageValue)

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}
	maxValue, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Maximum value: %.2f\n", maxValue)
}
