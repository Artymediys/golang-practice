package main

import (
	"fmt"
	"strings"
)

func main() {
	inputStr := "snow dog sun cat"
	words := strings.Split(inputStr, " ") // Разбиваем исходную строку на слова

	reversedWords := make([]string, len(words)) // Инициализируем массив для ревёрса с длиной исходного массива слов

	/*
	 В цикле берём текущий индекс исходного массива
	 и в конец ревёрснутого массива добавляем элемент из начала исходного массива
	*/
	for i := range words {
		reversedWords[len(words)-1-i] = words[i]
	}

	// Склеиваем строку из ревёрснутого массива
	outputStr := strings.Join(reversedWords, " ")
	fmt.Println(outputStr)
}
