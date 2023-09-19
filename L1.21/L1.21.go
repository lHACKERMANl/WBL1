package main

import "fmt"

/*
Паттерн "Адаптер"  используется для преобразования интерфейса одного объекта в интерфейс, ожидаемый клиентом,
чтобы объекты с несовместимыми интерфейсами могли работать вместе. Вот пример адаптера на языке Go.

Допустим, у нас есть структура OldSystem с методом LegacyOperation,
и нам нужно использовать этот метод в новой системе, которая ожидает интерфейс NewInterface.
Мы можем создать адаптер для OldSystem, который реализует NewInterface и вызывает LegacyOperation изнутри.
*/

// OldSystem - структура с устаревшим интерфейсом
type OldSystem struct{}

// LegacyOperation - метод устаревшей системы
func (o *OldSystem) LegacyOperation() string {
	return "Old System Legacy Operation"
}

// NewInterface - новый интерфейс, ожидаемый клиентом
type NewInterface interface {
	NewOperation() string
}

// Adapter - адаптер для структуры OldSystem
type Adapter struct {
	Old *OldSystem
}

// NewOperation - метод адаптера, который вызывает LegacyOperation
func (a *Adapter) NewOperation() string {
	return a.Old.LegacyOperation()
}

func main() {
	oldSystem := &OldSystem{}
	adapter := &Adapter{Old: oldSystem}

	// Используем адаптер для вызова NewOperation
	result := adapter.NewOperation()

	fmt.Println(result)
}
