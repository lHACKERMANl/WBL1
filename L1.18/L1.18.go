package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++ // Увеличиваем значение счетчика на 1
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value // Возвращаем текущее значение счетчика
}

func main() {
	var wg sync.WaitGroup // WaitGroup используется для ожидания завершения всех горутин
	counter := Counter{}

	// Запускаем 1000 горутин для инкрементации счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup на 1 перед запуском горутины
		go func() {
			counter.Increment() // Инкрементируем счетчик в каждой горутине
			wg.Done()           // Уменьшаем счетчик WaitGroup после завершения горутины
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}
