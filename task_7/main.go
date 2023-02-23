package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Первый вариант через обычный map и мьютексы.
Блочим перед записью -> записываем -> разлочиваем.
*/
func option1() {
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	data := make(map[string]int)

	// Запускаем несколько горутин для записи данных
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			mutex.RLock() // Блокируем доступ к записи сразу нескольким горутинам
			data[strconv.Itoa(i)] = i
			mutex.RUnlock() // Разблокируем
		}(i)
	}

	// Запускаем несколько горутин для чтения данных
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(key int) {
			defer wg.Done()

			// Мьютексы не используем, потому что в этот момент не будет производиться никаких записей
			value := data[strconv.Itoa(key)]

			fmt.Printf("%d: %d\n", key, value)
		}(i)
	}

	wg.Wait()
}

/*
Второй вариант через sync.Map.
Встроенные функции sync.Map используют изнутри мьютексы, поэтому нам не надо вручную заботиться об этом.
При чтении sync.Map могут даже выигрывать немного в скорости, но в подавляющем большиснтве случаев
использование обычных map и RWMutex-ов быстрее, но sync.Map позволяет избавиться от возможных косяков.
*/
func option2() {
	var data sync.Map // Инициализируем map из пакета sync
	var wg sync.WaitGroup

	// Запускаем несколько горутин для записи данных
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			data.Store(strconv.Itoa(i), i) // Используем функцию Store для записи в map
		}(i)
	}
	wg.Wait()

	// Итерация по всем элементам map с помощью функции Range
	data.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true // Продолжаем итерирование(если бы передали false, то итерирование бы прекратилась)
	})
}

func main() {
	option1()
	option2()
}
