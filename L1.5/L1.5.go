package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	dataCh := make(chan int)
	doneCh := make(chan struct{})

	// Запускаем горутину для чтения из канала
	go func() {
		for {
			select {
			case value := <-dataCh:
				fmt.Printf("Прочитано: %d\n", value)
			case <-doneCh:
				return // Завершаем горутину по сигналу о завершении
			}
		}
	}()

	var N int
	fmt.Print("Введите число горутину: ")
	fmt.Scan(&N) // Задайте желаемое количество горутину

	// Запускаем горутину для отправки данных в канал каждую секунду
	go func() {
		for i := 1; i <= N; i++ {
			dataCh <- i
			time.Sleep(time.Second)
		}

		// Отправляем сигнал о завершении
		close(doneCh)
	}()

	// Ожидаем N секунд
	time.Sleep(time.Duration(N) * time.Second)

	// Завершаем программу
	fmt.Println("Программа завершена.")
	elapsed := time.Since(start)
	fmt.Printf("Время выполнения: %s\n", elapsed)
}
