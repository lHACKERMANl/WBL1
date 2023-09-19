package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова обратно в строку
	result := strings.Join(words, " ")

	return result
}

func main() {
	input := "snow dog sun — sun dog snow"
	fmt.Println("Исходная строка:", input)

	// Переворачиваем слова в строке
	reversed := ReverseWords(input)
	fmt.Println("Перевернутая строка:", reversed)
}
