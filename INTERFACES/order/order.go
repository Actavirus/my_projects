// Для него нужно будет создать разные типы структур и интерфейсы. Вот как он будет выглядеть:
package order

import (
	"fmt"
)

// New ProcessOrder - Цель данной функции в создании нового экземпляра внутреннего заказа (national) и международного (international).
func New() { // экспортируемая функция, т.к. начинается с заглавной буквы
	fmt.Println("Order package!!")       // "Пакет заказа!!"
	natOrd := NewNationalOrder()         // новый внутренний заказ
	intOrd := NewInternationalOrder()    // новый международный заказ
	ords := []Operations{natOrd, intOrd} //
	ProcessOrder(ords)
}

// Product struct - структура "продукт"
type Product struct {
	name  string
	price int
}

// ProductDetail struct - структура "детали продукта"
type ProductDetail struct {
	Product         // структура "продукт"
	amount  int     // количество
	total   float32 // всего (возможно это цена за единицу)
}

// Summary struct - структура "цены"
type Summary struct {
	total          float32 // итого
	subtotal       float32 // промежуточный итог
	totalBeforeTax float32 // цена до налогооблажения
}

// ShippingAddress struct - структура "адрес доставки"
type ShippingAddress struct {
	street  string // улица
	city    string // город
	country string // страна
	cp      string // возможно почтовый индекс (postcode)
}

// Client struct - структура "клиент"
type Client struct {
	name     string // имя
	lastName string // фамилия
	email    string // электронная почта
	phone    string // телефон
}

// Order struct - структура "заказ"
type Order struct {
	products []*ProductDetail // указатель на срез (slice) товаров
	Summary
	ShippingAddress
	Client
}

// Processer interface - интерфейс для функции FillOrderSummar,
type Processer interface {
	FillOrderSummary()
}

// Printer interface
type Printer interface {
	PrintOrderDetails()
}

// Notifier interface
type Notifier interface {
	Notify()
}

// Добавим ещё один интерфейс для разбора работы программы
type MyInt interface {
	MyFunc() // данная функция есть у InternatioanlOrder, и у NationalOrder
}

// Operations interface - компонуем несколько других интерфейсов, что оказывается очень кстати, поскольку позволяет программе объединять объекты и делает код удобным для повторного использования.
type Operations interface {
	Processer
	Printer
	Notifier
	MyInt
}

// ProcessOrder function - получает массив заказов. Здесь у нас интересный момент. Вместо того, чтобы получать массив фактических объектов, эта функция получает их абстракцию.
func ProcessOrder(orders []Operations) {
	for _, order := range orders {
		order.FillOrderSummary()
		order.Notify()
		order.PrintOrderDetails()
		order.MyFunc()
	}
}
