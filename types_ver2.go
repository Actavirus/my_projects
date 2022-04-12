package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons { // здесь Liters - метод для Liters; ToGallons - имена могут быть одинаковыми, елси они определяются для разных типов.
	return Gallons(l * 0.264) // блок метода не отличается от блока функции.
}
func (m Milliliters) ToGallons() Gallons { // здесь Milliliters - метод для Milliliters; ToGallons - имена могут быть одинаковыми, если они определяются для разных типов.
	return Gallons(m * 0.000264) // блок метода не оличается от блока функции.
}
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}
func main() {
	soda := Liters(2)                                                                // создание значения Litters.
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())        // здесь soda.ToGallons() - преобразование Liters  в Gallons.
	water := Milliliters(500)                                                        // создание значения Milliliters
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons()) // здесь water.ToGallons() - преобразование Milliliters в Gallons.
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}
