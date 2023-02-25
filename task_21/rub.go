package main

import "log"

// Объявляем интерфейс с методом получения рубля
type IRubReceiver interface {
	ReceiveRUB(rub RubleValue)
}

// Объявляем структуру рубля
type RubleValue struct {
	rubles  uint16
	kopecks uint8
}

// Объявляем структуру рублёвого кошелька
type RubleWallet struct {
	balance RubleValue
}

// Реализация интерфейса IRubReceiver
func (rw *RubleWallet) ReceiveRUB(rub RubleValue) {
	rw.balance.rubles += rub.rubles
	rw.balance.kopecks += rub.kopecks

	log.Printf("Received: (%d.%d)RUB", rub.rubles, rub.kopecks)
}
