package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Для реализации программы, в которой главный поток записывает данные в канал,
а несколько воркеров читают данные из канала и выводят их в stdout,
можно использовать горутины и каналы.
Для возможности завершения работы всех воркеров по нажатию Ctrl+C, можно использовать sync.WaitGroup и os.Signal.
*/

func main() {
	// Создаем канал для обмена данными между главным потоком и воркерами
	dataCh := make(chan string)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запускаем N воркеров
	var numWorkers int
	fmt.Print("Введите число воркеров: ")
	fmt.Scan(&numWorkers) // Задайте желаемое количество воркеров

	if numWorkers <= 0 {
		fmt.Errorf("numworkers не больше или не равен 1. Введите корректное значение")
		return
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataCh, &wg)
	}

	// Создаем канал для обработки сигнала Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Главный поток записывает данные в канал
	go func() {
		defer close(dataCh)
		for i := 0; i < 10; i++ {
			dataCh <- fmt.Sprintf("Данные %d", i)
		}
	}()

	// Ожидаем сигнала Ctrl+C для завершения работы
	<-sigCh

	// Завершаем воркеров и ожидаем их завершения
	wg.Wait()
}

func worker(id int, dataCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataCh {
		fmt.Printf("Воркер %d получил: %s\n", id, data)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}
