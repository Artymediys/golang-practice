package main

import (
	"fmt"
	"strings"
)

/*
Исходный код:

	var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}

Проблема этого кода в том, что здесь идёт работа с огромными строками,
на которые не хочется по сто раз выделять память.
Плюс тут используется глобальная переменная, которая выделяется в хипе(куче), а он работает медленнее.
*/

func createHugeString(size int) string {
	var builder strings.Builder

	// С помощью билдера выделяем определённое кол-во памяти и записываем нужную строку
	builder.Grow(size)
	builder.WriteString("jdkasfhldshafdbskjfnkdsnfkdsnfbkdbvdlskfdslfnksabfnksbfdklsafbldksabfldsbfadksbnfkdbsfkabdsfdsbfdsbf.jbdslafbdslbvdbskvbdfsbvldsbhblsflfbnadsbfdksbfvdshvkasvkjdbfejkgv.dksbnfajkbnvk.dfbaskgvbdfsgbkbfgasbfkdg")

	return builder.String()
}

func someFunc() {
	/*
	 Создаём большую строку и сразу кконвертируем в байты.
	 Все последующие изменения/копирования лучше делать в рамках слайса байтов,
	 и вот только когда понадобится финальная строка, тогда уже конвертировать в string
	*/
	b := []byte(createHugeString(1 << 10))

	// Допустим мы тут несколько раз изменяем/копируем строку
	b = b[:100]

	fmt.Println(string(b)) // Вывод финального результата
}

func main() {
	someFunc()
}
