package main

import (
	"fmt"
)

/*
Для решения задачи можно использовать битовые маски и операторы | и &^ для установки i-го бита в 1 или 0,
в зависимости от значения setBit. Если setBit равен true, то i-й бит устанавливается в 1,
в противном случае он устанавливается в 0.
*/

func main() {
	var num int64 = 42 // Пример исходного числа
	var i int64        // Индекс бита, который нужно установить (считая с нуля)
	fmt.Print("Введите позицию числа: ")
	fmt.Scan(&i)

	if i < 0 || i > 64 {
		fmt.Errorf("Число должно быть в диапазоне от 0 до 64")
	}

	setBit := true // Установить бит в 1 (true) или 0 (false)

	if setBit {
		// Устанавливаем i-й бит в 1
		num |= (1 << i)
	} else {
		// Устанавливаем i-й бит в 0
		num &^= (1 << i)
	}

	fmt.Printf("Число после установки %d-го бита: %d\n", i, num)
}
