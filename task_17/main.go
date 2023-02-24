package main

import (
	"fmt"
)

func binarySort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		left, right := 0, i-1

		for left <= right {
			mid := (left + right) / 2
			if arr[mid] > key {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		for j := i - 1; j >= left; j-- {
			arr[j+1] = arr[j]
		}
		arr[left] = key
	}

	return arr
}

func main() {
	arr := []int{5, 2, 8, 1, 6}
	fmt.Println(binarySort(arr))
}
