package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаем значения a и b, которые больше 2^20
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	// Умножение a на b
	multiplication := new(big.Int)
	multiplication.Mul(a, b)
	fmt.Printf("Умножение: %s\n", multiplication.String())

	// Деление a на b
	division := new(big.Rat)
	division.SetFrac(a, b)
	fmt.Printf("Деление: %s\n", division.FloatString(10))

	// Сложение a и b
	addition := new(big.Int)
	addition.Add(a, b)
	fmt.Printf("Сложение: %s\n", addition.String())

	// Вычитание b из a
	subtraction := new(big.Int)
	subtraction.Sub(a, b)
	fmt.Printf("Вычитание: %s\n", subtraction.String())
}
