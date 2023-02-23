package main

import "fmt"

func main() {
	// Слайс для пересекающихся элементов(ещё можно просто в stdout выводить пересекающиеся элементы)
	var intersection []int

	/*
	 Два слайса с тестовыми значениями.
	 Ещё можно было создать maps со значениями в виде ключей для быстрого сравнения,
	 но самурай же не ищет лёгких путей.
	*/
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Создаём map с ключами, равными элементам первого слайса, для облегчённого поиска+сравнения со вторым слайсом
	mapSet1 := make(map[int]struct{})
	for _, value := range set1 {
		mapSet1[value] = struct{}{}
	}

	// Ищем наличие элементов из слайса №2 в map с элементами слайса №1
	for _, value := range set2 {
		if _, ok := mapSet1[value]; ok {
			intersection = append(intersection, value)
		}
	}

	fmt.Println(intersection)
}
