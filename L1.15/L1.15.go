package main

import (
	"fmt"
	"math/rand"
)

/*
Данный фрагмент кода может привести к проблеме с утечкой памяти. В этом коде мы создаем большую строку v с
использованием createHugeString и затем присваиваем justString первые 100 символов этой строки.
Однако, оригинальная большая строка v не освобождается, и это может привести к утечке памяти.
*/

func someFunc() string {
	// Создаем случайную строку определенной длины с помощью функции createHugeString
	v := createHugeString(1 << 10)

	// Возвращаем первые 100 символов этой строки
	return string([]rune(v)[:100])

}

func createHugeString(length int) string {
	// Определяем набор символов, из которых будет создаваться случайная строка
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	// Создаем слайс для хранения случайных символов
	b := make([]rune, length)

	// Заполняем слайс случайными символами из letterRunes
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	// Преобразуем слайс символов в строку и возвращаем ее
	return string(b)
}

func main() {
	// Вызываем функцию someFunc и выводим результат (первые 100 символов случайной строки)
	result := someFunc()
	fmt.Println(result)
}
