package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Для конкурентной записи данных в map в Go, необходимо учитывать,
что map не является потокобезопасной структурой данных.
Однако можно использовать мьютексы для синхронизации доступа к map.
*/

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.data[key]
}

func main() {
	safeMap := NewSafeMap()

	// Запускаем несколько горутин для конкурентной записи данных
	for i := 0; i < 10; i++ {
		go func(index int) {
			key := fmt.Sprintf("key%d", index)
			value := index * 10
			safeMap.Set(key, value)
			fmt.Printf("Установлено: %s -> %d\n", key, value)
		}(i)
	}

	// Ждем завершения всех горутин
	// Так как горутины не блокируются на чтении данных, программа может завершиться, прежде чем все горутины завершат запись
	// Для упрощения примера, используем задержку
	// В реальном приложении следует использовать sync.WaitGroup для ожидания завершения горутин
	fmt.Println("Ждем завершения записи данных...")
	// В реальных приложениях следует использовать более надежные методы ожидания завершения горутин
	// Например, sync.WaitGroup или другие способы синхронизации
	// Здесь используется простая задержка для демонстрационных целей
	// Не следует использовать такой код в производстве
	<-time.After(time.Second * 2)

	// Чтение данных из map
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := safeMap.Get(key)
		fmt.Printf("Прочитано: %s -> %d\n", key, value)
	}
}
