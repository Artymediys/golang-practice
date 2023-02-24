package main

import "fmt"

/*
 1) Выбираем pivot-элемент.
 2) Сортируем левую и правую половины массива, рекурсивно вызывая функцию.
 3) Объединяем их с помощью pivot-элемента, используя функцию append.
*/

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int
	for _, val := range arr[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{5, 2, 8, 1, 6}
	fmt.Println(quicksort(arr))
}
