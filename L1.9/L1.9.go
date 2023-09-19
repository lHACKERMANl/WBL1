package main

import (
	"fmt"
	"sync"
)

/*
В этой программе мы создаем два канала:
	inputCh для передачи чисел
	outputCh для передачи результатов операции x*2.

    Горутина №1 заполняет inputCh числами из массива.
    Горутина №2 выполняет операцию x*2 и записывает результат в outputCh.
    Горутина №3 выводит результаты из outputCh в stdout.

sync.WaitGroup используется для ожидания завершения вывода результатов.
Как только все числа обработаны и результаты выведены, программа завершается.
*/

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	// Горутина для записи чисел в первый канал
	go func() {
		defer close(inputCh)
		numbers := []int{2, 5, 7, 32, 4, 23}
		for _, num := range numbers {
			inputCh <- num
		}
	}()

	// Горутина для выполнения операции x*2 и записи результата во второй канал
	go func() {
		defer close(outputCh)
		for num := range inputCh {
			result := num * 2
			outputCh <- result
		}
	}()

	// Горутина для вывода результатов в stdout
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range outputCh {
			fmt.Printf("printCh: %v\n", result)
		}
	}()

	// Ожидаем завершения вывода
	wg.Wait()
}
