package main

import (
	"fmt"
	"strings"
)

/*
Проверяем, что все символы в строке уникальные и возвращаем true, если они уникальные
и false, если они не уникальные.
*/
func IsUnique(str string) bool {
	str = strings.ToLower(str) // Приводим строку к нижнему регистру

	uniqueChars := make(map[rune]bool) // Создаем словарь для отслеживания уникальных символов

	// Проходимся по каждому символу в строке и проверяем, был ли он уже добавлен в словарь
	for _, char := range str {
		if uniqueChars[char] {
			return false // Если символ уже был добавлен, то строка не уникальна
		}
		uniqueChars[char] = true // Если символ еще не был добавлен, то добавляем его в словарь
	}

	return true // Если мы прошли все символы и не обнаружили повторяющихся, то строка уникальна
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(str1, IsUnique(str1))
	fmt.Println(str2, IsUnique(str2))
	fmt.Println(str3, IsUnique(str3))
}
