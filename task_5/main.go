package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
Определяем функцию для отправки рандомных данных в канал. Отправка происходит до того момента,
пока полю(каналу) объекта t не прилетит инфа о текущем времени(по истечению таймера)
*/
func sendValues(wg *sync.WaitGroup, ch *chan int, t *time.Timer) {
	defer wg.Done()

	for {
		select {
		// Когда объект t получает текущее время, то закрываем канал и выходим из функции(+wg.Done)
		case <-t.C:
			close(*ch)
			return
		default:
			*ch <- rand.Intn(42) // Отправляем в канал рандомные инты от 0 до 41

			time.Sleep(time.Millisecond) // Имитация нагрузки
		}
	}
}

/*
Определяем функцию для приёма данных из канала
*/
func readValues(wg *sync.WaitGroup, ch *chan int) {
	defer wg.Done()

	// В цикле проходимся по данным из канала или ждём новые данные(пока не закроется)
	for data := range *ch {
		fmt.Printf("Data: %d\n", data)
	}
}

func main() {
	ch := make(chan int, 16) // Создаем буферизированный канал типа int с определённой вместимостью
	var wg sync.WaitGroup    // Инициализируем переменную/объект wg структуры WaitGroup

	// Объявляем переменную/объект timer структуры Timer с таймером в определённое кол-во секунд
	timer := time.NewTimer(5 * time.Second)

	wg.Add(2)
	go sendValues(&wg, &ch, timer) // Вызываем в отдельной горутине функцию отправки данных в канал
	go readValues(&wg, &ch)        // Вызываем в отдельной горутине функцию приёма данных из канала

	wg.Wait()
	log.Println("The app finished!")
}
