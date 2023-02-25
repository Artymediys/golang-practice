package main

import "log"

// Объявляем структуру доллара. Для разнообразия примера, здесь необычные типы данных, отличные от рубля
type DollarValue struct {
	dollars rune
	cents   rune
}

// Объявляем структуру долларового кошелька
type DollarWallet struct {
	balance DollarValue
}

// Метод для получения долларов(её мы и адаптируем)
func (dw *DollarWallet) ReceiveUSD(usd DollarValue) {
	dw.balance.dollars += usd.dollars
	dw.balance.cents += usd.cents

	log.Printf("Received: (%d.%d)USD", usd.dollars, usd.cents)
}
