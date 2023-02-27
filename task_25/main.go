package main

import "time"

func customSleep(t int64) {
	// Возвращает канал, который ждёт данные(данные приходят по истечению определённого времени)
	<-time.After(time.Second * time.Duration(t))
}

func main() {
	customSleep(3) // Засыпаем на 3 секунды
}
