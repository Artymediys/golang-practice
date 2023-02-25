package main

/*
Реализация паттерна "Адаптер" на примере рубля и доллара.
*/
func main() {
	// Создаём объект рубля
	rub := &RubleValue{
		rubles:  13,
		kopecks: 37,
	}

	// Создаём объект доллара
	usd := &DollarValue{
		dollars: 32,
		cents:   28,
	}

	// Создаём объект адаптера для доллара
	dollarAdapter := &DollarAdapter{
		&DollarWallet{
			balance: *usd,
		},
	}

	// Вызываем метод адаптера
	dollarAdapter.ReceiveRUB(*rub)
}
