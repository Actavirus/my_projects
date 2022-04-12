// тема - Методы в Go — Создание и использование методов в Golang
package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func (k kelvin) celsius() celsius { // пример метода celsius для типа kelvin
	return celsius(k - 273.15)
}
func kelvinToCelsius(k kelvin) celsius { // пример функции kelvinToCelsius
	return celsius(k - 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k.celsius()).fahrenheit())
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32)*5/9 + 273.15)
}
func main() {
	var k kelvin = 294
	var c celsius = 294
	var f fahrenheit = 294
	fmt.Printf("%.2f° K is %.2f° C\n", k, kelvinToCelsius(k))                    // пример вызова функции
	fmt.Printf("%.2f° K is %.2f° C\n", k, k.celsius())                           // пример вызова метода
	fmt.Printf("%.2f° K is %.2f° F\n", k, k.fahrenheit())                        // пример вызова метода
	fmt.Printf("%.2f° C is %.2f° F\n", c, c.fahrenheit())                        // пример вызова метода
	fmt.Printf("%.2f° F is %.2f° C\n", f, f.celsius())                           // пример вызова метода
	fmt.Printf("%.2f° F is %.2f° K\n", f, f.kelvin())                            // пример вызова метода
	fmt.Printf("%.2f° K is %.2f° K\n", k, ((k.celsius()).fahrenheit()).kelvin()) // K -> C -> F -> K
}
