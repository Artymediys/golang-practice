package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция для бесконечной отправки данных в канал с опцией завершения
func infiniteSending(wg *sync.WaitGroup, ch *chan int) {
	defer wg.Done()

	/*
	 Создаём буферизированный канал для "стоп-сигналов" SIGINT и SIGTERM.
	 Буферизированный канал нужен для того, чтобы он не заблочился.
	 SIGINT для CTRL+C, а SIGTERM для обработки завершения процесса
	*/
	log.Println("Sending started! Stop the sending - CTRL+C")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Бесконечно отправляем рандомные чиселки в канал, пока не придёт стоп-сигнал
	for {
		select {
		case <-stop:
			close(*ch) // Закрываем канал
			return
		default:
			*ch <- rand.Intn(42)
		}
	}
}

func main() {
	var workersCount int // Инициализируем переменную для кол-ва воркеров

	fmt.Print("Enter the count of workers: ")

	_, err := fmt.Scanf("%d", &workersCount) // Читаем из stdin кол-во воркеров
	if err != nil {
		log.Fatalln(err)
		// Тут можно сделать вместо выхода из программы присвоение дефолтного кол-во воркеров
	}

	ch := make(chan int, 16) // Создаем буферизированный канал типа int с определённой вместимостью

	var wg sync.WaitGroup // Инициализируем переменную/объект wg структуры WaitGroup
	wg.Add(workersCount)  // Увеличиваем счётчик горутин на кол-во воркеров

	// Запускаем воркеров
	for id := 1; id <= workersCount; id++ {
		go func(workerID int) {
			defer wg.Done() // После завершения работы анонимной функции уменьшаем счётчик горутин

			// Получаем данные из канала пока канал не закроется
			for data := range ch {
				fmt.Printf("Worker #%d received: %d\n", workerID, data)
				time.Sleep(time.Millisecond) // Имитация нагрузки
			}
		}(id)
	}

	// Вызываем функцию по бесконечному спаму данных в канал
	wg.Add(1)
	go infiniteSending(&wg, &ch)

	wg.Wait() // Ожидаем, пока счётчик горутин не будет равен 0
	log.Println("The app finished!")
}
