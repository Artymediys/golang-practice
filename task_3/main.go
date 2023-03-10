package main

import (
	"fmt"
	"sync"
)

func option1() {
	numbers := [5]uint8{2, 4, 6, 8, 10}      // Инициализиурем массив nums и сразу же присваиваем значения элементов
	var sum uint8                            // Инициализируем переменную для последующего накопления суммы квадратов
	result := make(chan uint8, len(numbers)) // Инициализируем буферизированный канал типа uint8

	// Проходимся по массиву nums, выбирая только значения(индексы опускаем)
	for _, v := range numbers {
		/*
		 Выделяем новую горутину на выполнение анонимной функции
		 и передаём в функцию определённое значение из массива
		*/
		go func(x uint8) {
			result <- x * x // Передаём в канал квадрат числа
		}(v)
	}

	// В цикле читаем из канала полученные квадраты и суммируем их
	for i := 0; i < len(numbers); i++ {
		sum += <-result
	}

	fmt.Printf("The sum of squared = %d", sum) // Выводим в поток вывода расчёты
}

func option2() {
	numbers := [5]uint8{2, 4, 6, 8, 10} // Инициализиурем массив nums и сразу же присваиваем значения элементов
	var sum uint8                       // Инициализируем переменную для последующего накопления суммы квадратов

	var wg sync.WaitGroup  // Инициализируем переменную/объект wg структуры WaitGroup
	var mutex sync.RWMutex // Инициализируем переменную/объект mutex структуры RWMutex

	// Проходимся по массиву nums, выбирая только значения(индексы опускаем)
	for _, v := range numbers {
		wg.Add(1) // Увеличиваем счётчик горутин на единичку

		/*
		 Выделяем новую горутину на выполнение анонимной функции
		 и передаём в функцию определённое значение из массива
		*/
		go func(x uint8) {
			defer wg.Done() // После завершения работы анонимной функции уменьшаем счётчик горутин

			mutex.RLock() // Блокируем следующим потокам доступ к записи
			sum += x * x
			mutex.RUnlock() // Разблокируем
		}(v)
	}

	wg.Wait() // Ожидаем, пока счётчик горутин не будет равен 0

	fmt.Printf("The sum of squared = %d", sum) // Выводим в поток вывода расчёты
}

func main() {
	// Вызов первого или второго варианта реализации
	option1()
	option2()
}
