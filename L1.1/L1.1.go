package main

import "fmt"

/*
На Go нет классической концепции наследования, как в объектно-ориентированных языках программирования.
Вместо этого Go использует композицию для достижения подобного поведения.
Необходимо встроить методы структуры Human в структуру Action с помощью композиции
*/

// Определяем структуру Human с произвольными полями и методами
type Human struct {
	Name string
	Age  int
}

// Метод SayHello для структуры Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Определяем структуру Action, встраивая в нее структуру Human
type Action struct {
	Human
	ActionName string
}

// Метод DoAction для структуры Action
func (a *Action) DoAction() {
	fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionName)
}

func main() {
	// Создаем объект структуры Action
	action := Action{
		Human:      Human{Name: "Влад", Age: 23},
		ActionName: "делает L1.1",
	}

	// Вызываем методы из структуры Human и структуры Action
	action.SayHello() // Вызываем метод SayHello структуры Human
	action.DoAction() // Вызываем метод DoAction структуры Action
}
