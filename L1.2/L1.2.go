package main

import (
	"fmt"
	"sync"
)

/*
Можно использовать конкурентные горутины в Go для параллельного
вычисления квадратов чисел из массива и вывода результатов в stdout
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int) // Канал для хранения результатов

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для вычисления квадратов
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x
			results <- square
		}(num)
	}

	// Горутина для закрытия канала после завершения всех вычислений
	go func() {
		wg.Wait()
		close(results)
	}()

	// Читаем результаты из канала и выводим их в stdout
	for result := range results {
		fmt.Println(result)
	}
}
