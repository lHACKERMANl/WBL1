package main

import (
	"context"
	"fmt"
	"time"
)

// 1. спользование канала для завершения горутины
func worker1(done chan bool) {

	// Выполнение какой-либо работы
	// Когда нужно завершить горутину:
	fmt.Println("Горутина в worker 1 завершена")
	done <- true
}

// 2. Использование контекста для завершения горутины
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина в worker 2 завершена")
			return
		default:
			// Выполнение какой-либо работы
		}
	}
}

// 3. Использование return для выхода из функции горутины
func worker() {
	// Выполнение какой-либо работы
	fmt.Println("Горутина в worker 3 завершена")
	return // Горутина завершится после выполнения этой строки
}

func main() {
	done := make(chan bool)
	go worker1(done)
	// Дожидаемся завершения горутины
	<-done

	/////////////////////////////////////////////////////////

	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	// Через некоторое время отменяем контекст
	time.Sleep(time.Second)
	cancel()
	// Дожидаемся завершения горутины
	time.Sleep(time.Second)

	/////////////////////////////////////////////////////////

	go worker()
	// Дожидаемся завершения горутины (не требуется в данном случае)
	// В большинстве случаев, эта строчка можно опустить
	time.Sleep(time.Second)

}
