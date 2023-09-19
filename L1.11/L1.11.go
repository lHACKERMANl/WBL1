package main

import (
	"fmt"
)

func main() {
	// Создаем два неупорядоченных множества (карты)
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true, 7: true}

	// Создаем пустую карту для хранения пересечения
	intersection := make(map[int]bool)

	// Находим пересечение множеств
	for element := range set1 {
		if set2[element] {
			intersection[element] = true
		}
	}

	// Выводим результат пересечения
	fmt.Println("Пересечение множеств:")
	for element := range intersection {
		fmt.Println(element)
	}
}
