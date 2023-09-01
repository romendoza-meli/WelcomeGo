package main

import (
	"fmt"
	"sort"
)

// calculateAverage calcula el promedio de las calificaciones proporcionadas.
// Ignora las calificaciones negativas y devuelve 0 si no hay calificaciones válidas.
func calculateAverage(grades ...float64) float64 {
	// Ordena las calificaciones de forma ascendente.
	sort.Float64s(grades)

	sum := 0.0
	count := 0

	// Calcula la suma de las calificaciones válidas y cuenta cuántas hay.
	for _, grade := range grades {
		if grade < 0 {
			continue
		}
		sum += grade
		count++
	}

	// Si no hay calificaciones válidas, devuelve 0 para evitar una división por cero.
	if count == 0 {
		return 0
	}

	// Calcula y devuelve el promedio de las calificaciones válidas.
	return float64(sum) / float64(count)
}

func main() {
	grades := []float64{90.0, -10.0, 85.0, 92.0, 77.0}
	average := calculateAverage(grades...)
	fmt.Printf("The average grade is: %.2f\n", average)
}
