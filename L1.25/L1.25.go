package main

import (
	"fmt"
	"time"
)

// Sleep - собственная функция для задержки выполнения программы
func Sleep(seconds int) {
	select {
	case <-time.After(time.Second * time.Duration(seconds)):
		return
	}
}

func main() {
	fmt.Println("Начало выполнения программы")

	// Используем собственную функцию Sleep для задержки на 3 секунды
	Sleep(3)

	fmt.Println("Завершение выполнения программы после задержки")
}
