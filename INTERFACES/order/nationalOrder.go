/* Этот файл представляет вторую фактическую реализацию интерфейса Operations.
Здесь содержится тип структуры NationalOrder, который тоже использует тип структуры Order.
Далее идет функция инициализации, устанавливающая товары, информацию о клиенте и адрес доставки конкретного заказа внутри страны.
Затем, как и в предыдущем файле, следуют определения всех методов, необходимых для реализации интерфейса.
Теперь структура NationalOrder тоже реализует интерфейс Operations, так как отвечает на все его методы.
*/
package order

import (
	"fmt"
)

// экземпляр national
var national = &NationalOrder{}

// структура NationalOrder
type NationalOrder struct {
	Order
}

// функция NewNatioanlOrder
func NewNationalOrder() *NationalOrder {
	national.products = append(national.products, GetProductDetail("Sugar", 12, 3, 36))
	national.products = append(national.products, GetProductDetail("Cereal", 16, 2, 36))
	national.Client = SetClient("Phill", "Heat", "phill@gmail.com", "8415748569")
	national.ShippingAddress = SetShippingAddress("North Ave", "San Antonio", "USA", "854789")
	return national
}

// функция FillOrderSummary
func (nato *NationalOrder) FillOrderSummary() {
	var taxes float32 = 0.20     // налог
	var shippingCost float32 = 5 // стоимость доставки
	subtotal = CalculateSubTotal(nato.products)

	totalBeforeTax = (subtotal + shippingCost)
	totalTaxes = (taxes * subtotal)
	total = (subtotal + totalTaxes)

	nato.Summary = Summary{
		total,
		subtotal,
		totalBeforeTax,
	}
}

// функция Notify
func (nato *NationalOrder) Notify() {
	email := nato.Client.email
	fmt.Println("---Natioanl Order---")
	fmt.Println("Sending email notification to:", email)
}

// функция PrintOrderDetails
func (nato *NationalOrder) PrintOrderDetails() {
	fmt.Println()
	fmt.Println("Natioanl Summary")
	fmt.Println("Order details: ")
	fmt.Println("Total: ", nato.Summary.total)
	fmt.Printf("Delivery Address to: %s %s %s \n", nato.ShippingAddress.street, nato.ShippingAddress.city, nato.ShippingAddress.country)
}

// добавим дополнительную функцию, которая будет выводить данные о продукте
func (into NationalOrder) MyFunc() {
	fmt.Println()
	fmt.Println("Работа функции MyFunc")
	fmt.Println("Выводим информацию о продукте/-ах:")
	fmt.Println(into.products)
}
