package main

import (
	"fmt"
	"sync"
)

func main() {
	const elemCount uint8 = 5

	// Задаём элементы массива
	arr := [elemCount]int{2, 4, 6, 8, 10}

	// Создаём буферизированные каналы
	firstChan := make(chan int, elemCount)
	secondChan := make(chan int, elemCount)

	// В цикле отправляем данные в два канала и выводим из второго канала
	var wg sync.WaitGroup
	for _, value := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()

			firstChan <- x      // В первый канал пишем число из массива
			secondChan <- x * 2 // Во второй канал пишем число из массива, умноженное на 2
		}(value)

		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(<-secondChan) // Выводим данные из второго канала в stdout
		}()
	}

	wg.Wait()
}
