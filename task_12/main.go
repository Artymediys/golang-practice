package main

import "fmt"

func main() {
	// Массив с исходными строками
	source := [5]string{"cat", "cat", "dog", "cat", "three"}

	// Создаём множество с использованием map
	var set = make(map[string]int) // вариант без счётчика map[string]struct{}

	/*
	 Можно было бы хранить только уникальные значения, но раз уж это map,
	 то решил включить ещё и счётчик повторяющихся элементов в исходном массиве
	*/
	for _, value := range source {
		set[value]++ // вариант без счётика set[value] = struct{}{}
	}

	fmt.Println(set)
}
