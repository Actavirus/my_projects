/* Этот файл является первой фактической реализацией интерфейса Operations.
Сначала мы создали тип структуры InternationalOrder, определив с помощью структуры Order его свойства и объекты.
Далее идет функция инициализации NewInternationalOrder, которая будет устанавливать товары для заказа, информацию о клиенте и адрес доставки.
Для инициализации ProductDetail, Client и ShippingAddress мы используем вспомогательную функцию, которую вскоре тоже реализуем.
В оставшейся части файла мы объявляем фактическую реализацию функций FillOrderSummary, Notify и PrintOrderDetails.
Теперь можно сказать, что тип структуры InternationalOrder реализует интерфейс Operations, потому что содержит определения всех его методов.
Круто!
*/
package order

import "fmt"

// структура international
var international = &InternationalOrder{}

// структура InternationalOrder
type InternationalOrder struct {
	Order
}

// фукнция NewInternationalOrder
func NewInternationalOrder() *InternationalOrder {
	international.products = append(international.products, GetProductDetail("Lap Top", 450, 1, 450.50))
	international.products = append(international.products, GetProductDetail("Video Game", 600, 2, 1200.50))
	international.Client = SetClient("Carl", "Smith", "carlsmith@gmail.com", "9658521365")
	international.ShippingAddress = SetShippingAddress("Colfax Avenue", "Seattle", "USA", "45712")
	return international
}

// функция FillOrderSummary
func (into *InternationalOrder) FillOrderSummary() {
	var extraFee float32 = 0.5                  // коэффициент дополнительной платы
	var taxes float32 = 0.25                    // коэффициент налога
	var shippingCost float32 = 35               // стоимость доставки
	subtotal = CalculateSubTotal(into.products) // передаём в функцию CalculateSubTotal данные по всем продуктам и поучаем общую стоимость (это если товаров заказано более 1).

	totalBeforeTax = (subtotal + shippingCost)      // считаем цену до налогооблажения товара/-ов (промежуточный итог + стоимость доставки)
	totalTaxes = (taxes * subtotal)                 // считаем цену налогооблажения для товара/-ов (коэффициент налога * промежуточный итог)
	totalExtraFee = (totalTaxes * extraFee)         // считаем цену дополнительной платы (цена налогооблажения * коэффициент дополнительной платы)
	total = (subtotal + totalTaxes) + totalExtraFee // считаем всю сумму (общая стоимость + цена налогооблажения + цена дополнительной платы)
	into.Summary = Summary{
		total:          total,
		subtotal:       subtotal,
		totalBeforeTax: totalBeforeTax,
	}
}

// функция Notify
func (into *InternationalOrder) Notify() {
	email := into.Client.email
	name := into.Client.name
	phone := into.Client.phone

	fmt.Println()
	fmt.Println("---Международный заказ---")
	fmt.Println("Уведомление: ", name)
	fmt.Println("Отправить уведомление на электронную почту :", email)
	fmt.Println("Отправить СМС-уведомление на номер ", phone)
	fmt.Println("Отправить уведомление на WhatsApp:", phone)
}

// функция PrintOrderDetails - метод вывода информации о международном заказе
func (into *InternationalOrder) PrintOrderDetails() { // думаю, что это Get-метод (Get-методы - это методы, предназначенные для получения значения поля структуры или переменной)
	fmt.Println()
	fmt.Println("Международный заказ")
	fmt.Println("Информация о заказе: ")
	fmt.Println("-- ИТОГО до налогооблажения: ", into.Summary.totalBeforeTax)
	fmt.Println("-- Промежуточный итог: ", into.Summary.subtotal)
	fmt.Println("-- Всего: ", into.Summary.total)
	fmt.Printf("-- Адрес доставки: %s, %s, %s \n", into.ShippingAddress.street, into.ShippingAddress.city, into.ShippingAddress.country) // можно использовать сокращённоый вариант: into.street, into.city, into.country
	fmt.Printf("-- Клиент: %s %s \n", into.Client.name, into.Client.lastName)                                                            // можно использовать сокращённый вариант: into.name, into.lastName
}

// добавим дополнительную функцию, которая будет выводить данные о продукте
func (into InternationalOrder) MyFunc() {
	fmt.Println()
	fmt.Println("Работа функции MyFunc")
	fmt.Println("Выводим информацию о продукте/-ах:")
	fmt.Println("-- Наименование товара: ", into.Order.products) // пытаемся указать наименование товара (тема для изучения: Указатели в Go). выводит: [0xc0000c04b0 0xc0000c04e0]
}
