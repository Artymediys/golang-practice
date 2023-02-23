package main

import (
	"fmt"
	"log"
)

func main() {
	var number int64 // Инициализация переменной типа int64

	fmt.Print("Enter the number: ")

	_, err := fmt.Scanf("%d", &number) // Читаем из stdin число
	if err != nil {
		log.Fatalln(err)
		// Тут можно сделать вместо выхода из программы присвоение дефолтного числа
	}

	// Выводим введённое число
	fmt.Printf("Number before changing: (%d / %b)\n\n", number, number)

	// Берём i-й бит от 1 до 5 (для примера)
	for i := 1; i <= 5; i++ {

		// Создаём маску для операции XOR(исключающего ИЛИ) и сразу же XOR-им
		var mask int64 = 1 << (i - 1) // Сдвиг влево эквивалентен операции -> 1 * 2^(i - 1)
		number = number ^ mask

		fmt.Printf("Number after changing: (%d / %b), Mask: (%d / %b)\n", number, number, mask, mask)
	}

}
