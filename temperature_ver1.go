// программа для перевода температуры из Цельсия в Фаренгейты и обратно
package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64

// func fakeValue() celsius { // это функция - если нужны рандомные числа от 40 до 100
// return celsius(rand.Intn(60) + 40)
// }

func celsiusValue() (int, celsius, fahrenheit) { // это функция - если нужны числа от 40 до 100 с шагом 5
	var meanC celsius
	var meanF fahrenheit
	var number int = 100
	for i := 40; i <= number; i += 5 {
		fmt.Printf("%12.2d °C|%12.2f °K\n", i, celsius(i).fahrenheit())
		meanC += celsius(i)
		meanF += celsius(i).fahrenheit()
	}
	number = (100 - 40) / 5
	return number, meanC, meanF
}

// func celsiusToFahrenheit(c celsius) fahrenheit { // это функция - перевод из Цельсия в Фаренгейт
// 	f := (c * 9 / 5) + 32
// 	return fahrenheit(f)
// }

func (c celsius) fahrenheit() fahrenheit { // это метод - перевод из Цельсия в Фаренгейт
	f := (c * 9 / 5) + 32
	return fahrenheit(f)
}
func main() {
	fmt.Printf("%15v|%15v\n", "Temperature C", "Temperature F")
	fmt.Println("===============|===============")
	// var meanC celsius
	// var meanF fahrenheit
	var number int
	// for i := 0; i < number; i++ {
	// 	var c celsius = fakeValue()
	// 	fmt.Printf("%12.2f °C|%12.2f °K\n", c, c.fahrenheit())
	// 	meanC += meanC
	// 	meanF += meanF
	// }
	number, meanC, meanF := celsiusValue()
	fmt.Println("===============|===============")
	fmt.Printf("%12.2f °C|%12.2f °K\n", meanC/celsius(number), meanF/fahrenheit(number))
}
