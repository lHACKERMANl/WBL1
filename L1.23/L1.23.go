package main

import "fmt"

func removeElement(slice []int, index int) []int {
	// Проверяем, что индекс находится в пределах допустимых значений
	if index < 0 || index >= len(slice) {
		return slice // Если индекс недопустим, возвращаем исходный слайс без изменений
	}

	// Удаляем элемент, смещая остальные элементы влево
	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]

	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2 // Удаляем третий элемент (индекс 2)

	slice = removeElement(slice, indexToRemove)

	fmt.Println("Слайс после удаления элемента:", slice)
}
