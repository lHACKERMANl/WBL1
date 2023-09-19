package main

import (
	"fmt"
)

func main() {
	// Создаем исходный массив, который нужно отсортировать
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Исходный массив:", arr)

	// Вызываем функцию быстрой сортировки
	quicksort(arr, 0, len(arr)-1)

	// Выводим отсортированный массив
	fmt.Println("Отсортированный массив:", arr)
}

// Функция быстрой сортировки
func quicksort(arr []int, low, high int) {
	// Проверяем, нужно ли сортировать этот подмассив
	if low < high {
		// Выбираем опорный элемент и выполняем разделение
		pivot := partition(arr, low, high)

		// Рекурсивно сортируем левую и правую части массива
		quicksort(arr, low, pivot-1)
		quicksort(arr, pivot+1, high)
	}
}

// Функция для разделения массива
func partition(arr []int, low, high int) int {
	// Задаем опорный элемент (в данном случае, последний элемент)
	pivot := arr[high]
	i := low - 1

	// Проходим по массиву и переставляем элементы, чтобы разделить их на две части
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Переставляем опорный элемент в правильное положение
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Возвращаем индекс опорного элемента
	return i + 1
}
