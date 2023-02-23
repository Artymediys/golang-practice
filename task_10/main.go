package main

import (
	"fmt"
)

func main() {
	// Задаём последовательность температур
	temperatures := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map, в котором будут храниться группы температурных значений
	groups := make(map[int][]float64)

	// Проходимся по всем значениям в последовательности
	for _, temp := range temperatures {

		/*
		 Округляем значение температуры до ближайшего десятка.
		 Пример работы:
		   1) int(-25.4) -> -25
		   2) (-25 % 10) -> -5  // Оказывается, %(remainder) и mod по-разному берут знак
		   3) (-25) - (-5) -> -20
		*/
		roundedTemp := int(temp) - (int(temp) % 10)

		// Добавляем значение температуры в соответствующую группу
		groups[roundedTemp] = append(groups[roundedTemp], temp)
	}

	// Выводим группы температурных значений
	for key, value := range groups {
		fmt.Printf("%d: %v\n", key, value)
	}
}
