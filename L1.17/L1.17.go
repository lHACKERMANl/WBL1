package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Если средний элемент равен целевому значению, возвращаем его индекс
		if arr[mid] == target {
			return mid
		}

		// Если целевое значение меньше среднего элемента, ищем в левой половине
		if arr[mid] > target {
			right = mid - 1
		} else { // Иначе ищем в правой половине
			left = mid + 1
		}
	}

	// Если целевое значение не найдено в массиве, возвращаем -1
	return -1
}

func main() {
	// Создаем отсортированный массив
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Задаем целевое значение для поиска
	target := 11

	// Вызываем функцию бинарного поиска
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Целевое значение %d найдено по индексу %d\n", target, index)
	} else {
		fmt.Printf("Целевое значение %d не найдено в массиве\n", target)
	}
}
