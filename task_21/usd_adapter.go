package main

// Объявляем структуру адаптара для доллара с полем - долларовый кошелёк
type DollarAdapter struct {
	dw *DollarWallet
}

/*
Реализуем адаптер, который будет рубли конвертировать в доллары.
Адаптер реализует интерфейс IRubReceiver.
P.S. Пример сугубо ради демонстрации паттерна.
*/
func (adapter *DollarAdapter) ReceiveRUB(rub RubleValue) {
	adapter.dw.ReceiveUSD(DollarValue{
		dollars: rune(rub.rubles * 72),
		cents:   rune(rub.kopecks),
	})
}
