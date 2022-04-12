// Попробуем создать код для расчёта стоимости акции (любой акции). Данные можно брать как из файла, так и из сети Интернет.
package main

import (
	"encoding/json"
	"fmt"
)

// для начала определим новый тип, который будет описывать какую-либо акцию (Share).
// определим её тип как struct, т.к. она будет иметь несколько параметров
type Share struct {
	Name               string  // наименование акции
	Ticker             string  // тикер акции - краткое название в биржевой информации котируемых инструментов (акций, облигаций, индексов).
	StockMarket        string  // наименование фондового рынка, на котором котируется акция (здесь, возможно, нужно будет подумать)
	Value              float64 // значение стоимости акции на конкретную дату (т.е., этот параметр будет меняться в зависимости от даты) - здесь необходимо настроить подключение к файлу или сети Интернет для получения данных
	Currency           rune    // валюта в которой торгуется акция на конкретной фондовой бирже (т.е., будет зависимость от параметра StockMarket)
	Number             int     // количество акций на текущий момент
	SectorOfTheEconomy string  // параметр описывает сектор экономики, к которой относится конкретная акция (на начальном этапе указан тип string, но нужно подумать: этот параметр меняется изсходя из конкретного списка, как всплывающее окно в Excel)
}

// также, возможно, стоит определить типы для каждого инструменты отдельно: акции, облигации, ПИФ (фонды), и пр. (фьючерсы)
// тип Bond описывает облигации
type Bond struct {
	Name   string // наименование облигации
	Number int    // количество облигаций на текущий момент
}

// определим функцию (а затем превратим её в метод) для расчёта всей суммы определённой акции
func Summ(v, n float64) float64 {
	return v * n
}

// превратим функцию Summ в метод
func (s Share) Summ() float64 {
	return s.Value * float64(s.Number)
}

// определим метод для показа информации через метод json
func (s Share) String() string {
	bytes, _ := json.MarshalIndent(s, " ", " ")
	return fmt.Sprintln(string(bytes))
}

// определим метод для покупки акции
func (s Share) Buy() float64 {
	return float64(s.Number) + 100
}

// определим метод для покупки облигации
func (b Bond) Buy() float64 {
	return float64(b.Number) + 50
}

// теперь для обоих методов определим интерфейс
type Purchase interface {
	Buy()
}

func main() {
	a1 := Share{
		Name:               "Apple Inc",
		Ticker:             "AAPL",
		StockMarket:        "NASDAG",
		Value:              172.39,
		Currency:           '$',
		Number:             100,
		SectorOfTheEconomy: "IT",
	}
	// выведем в консоль информацию по конкретной акции
	fmt.Println(a1)
	// найдём сумму определённой акции с помощью функции Summ
	fmt.Printf("Summ = %v %q\n", Summ(a1.Value, float64(a1.Number)), a1.Currency) // выводит: Summ = 17239 '$'
	// найдём сумму определённой акции с помощью метода Summ
	fmt.Printf("Summ = %v %q\n", a1.Summ(), a1.Currency) // выводит: Summ = 17239 '$'
	// попробуем получить данные по конкретной акции из файла (а потом из сети Интернет)

}
