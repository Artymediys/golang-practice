package main

import "fmt"

func checkType(variable interface{}) {
	switch t := variable.(type) {
	case int:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	case string:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	case bool:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	case chan int:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	case chan string:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	case chan bool:
		// Какой-то код
		fmt.Printf("Type = %T, Value: %#v\n", t, t)
	default:
		fmt.Println("No way...")
	}
}

func main() {
	var testVar interface{} // Под капотом InterfaceValue = nil

	// Для примера присвоим chan int
	ch := make(chan int, 5)
	testVar = ch // Под капотом InterfaceValue = (chan int)(адрес)
	checkType(testVar)

	// Ещё один пример с int
	num := 213
	testVar = num // Под капотом InterfaceValue = (int)(213)
	checkType(testVar)
}
