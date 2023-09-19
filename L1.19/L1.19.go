package main

import (
	"fmt"
)

// ReverseString переворачивает строку, содержащую руны
func ReverseString(input string) string {
	// Преобразуем строку в руны
	runes := []rune(input)
	length := len(runes)

	// Переворачиваем руны в массиве
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	// Преобразуем руны обратно в строку
	return string(runes)
}

func main() {
	input := "главрыба — абырвалг"
	fmt.Println("Исходная строка:", input)

	// Переворачиваем строку
	reversed := ReverseString(input)
	fmt.Println("Перевернутая строка:", reversed)
}
