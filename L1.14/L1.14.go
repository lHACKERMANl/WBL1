package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Пример переменных разных типов
	var v1 interface{} = 42
	var v2 interface{} = "Hello, World!"
	var v3 interface{} = true
	var v4 interface{} = make(chan int)

	// Функция для определения типа  и вывода его
	checkType := func(v interface{}) {
		switch v.(type) {
		case int:
			fmt.Printf("Тип для %v: int\n", v)
		case string:
			fmt.Printf("Тип для %v: string\n", v)
		case bool:
			fmt.Printf("Тип для %v: bool\n", v)
		case chan int:
			fmt.Printf("Тип для %v: chan int\n", v)
		default:
			fmt.Printf("Тип для %v: неизвестен\n", v)
		}
	}

	// Проверяем типы переменных
	fmt.Println("Использование checkType функции:")
	checkType(v1)
	checkType(v2)
	checkType(v3)
	checkType(v4)

	fmt.Println("\nИспользование reflect.TypeOf функции:")
	fmt.Println("Тип v1:", reflect.TypeOf(v1))
	fmt.Println("Тип v2:", reflect.TypeOf(v2))
	fmt.Println("Тип v3:", reflect.TypeOf(v3))
	fmt.Println("Тип v4:", reflect.TypeOf(v4))
}
