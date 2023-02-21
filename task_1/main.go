package main

import (
	"fmt"
	"log"
)

// Определяем структуру Human
type Human struct {
	name string
	age  uint8
}

/*
Определяем публичный метод Greeting структуры Human с возвращаемым значением типа string.
Функция возвращает строку, содержащую значения полей name и age объекта human после форматирования функцией Sprintf.
*/
func (human *Human) Greeting() string {
	return fmt.Sprintf(`Hi, I'm %s, %d years old.`, human.name, human.age)
}

/*
Определяем структуру Action с полем типа Human.
Встраивание подобным образом наделяет структуру Action методами структуры Human.
*/
type Action struct {
	Human
}

/*
Определяем публичный метод структуры Action с возвращаемым значением типа string.
Функция вызывает метод Greeting поля Human и результат конкатенирует с другой сторокой.
*/
func (action *Action) Run() string {
	return action.Greeting() + " And now i'm running!"
}

/*
В функции main инициализируем переменную action
и сразу же присваиваем ей значение в виде ссылки на экземпляр структуры Action.
Во время инициализации объекта структуры Action мы сразу же инициализируем и объект структуры Human внутри Action.
В конце с помощью функции Println пакета log выводим результат метода Run.
*/
func main() {
	action := &Action{Human{
		name: "Artyom",
		age:  23,
	}}

	log.Println(action.Run())
}
