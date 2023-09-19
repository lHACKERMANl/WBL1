package main

import (
	"fmt"
	"math"
)

func main() {
	// Создаем карту для хранения температурных групп
	temperatureGroups := make(map[float64][]float64)

	// Последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Шаг для группировки температурных значений
	step := 10.0

	// Группируем температурные значения
	for _, temp := range temperatures {
		groupKey := math.Floor(temp/step) * step // Определяем ключ группы
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	// Выводим группы
	for group, values := range temperatureGroups {
		fmt.Printf("Группа %.0f градусов: %v\n", group, values)
	}
}
