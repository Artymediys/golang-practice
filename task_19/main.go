package main

import "fmt"

func reverseString(str string) string {
	// Преобразуем строку в массив рун(int32) (так сможем изменять элементы строки)
	runes := []rune(str)

	// В цикле инициализируем i и j с помощью которых будем свапать первый и последний элемент с шагом в 1
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "♫главрыба©"
	fmt.Println(reverseString(str))
}
