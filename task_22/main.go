package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализируем переменные типа *big.Int
	x := new(big.Int)
	y := new(big.Int)

	// Задаём им huge значение
	x.SetString("9876543210987654321098765432109876543210", 10)
	y.SetString("1234567890123456789012345678901234567890", 10)

	// Перемножаем
	m := new(big.Int)
	m.Mul(x, y)
	fmt.Println("Multiplication: ", m)

	// Делим
	d := new(big.Int)
	d.Div(x, y)
	fmt.Println("Division: ", d)

	// Складываем
	a := new(big.Int)
	a.Add(x, y)
	fmt.Println("Addition: ", a)

	// Вычитаем
	s := new(big.Int)
	s.Sub(x, y)
	fmt.Println("Subtraction: ", s)
}
