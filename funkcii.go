// тема - Функции в Golang на примерах
// программа перевода градусов по шкале Цельсия, Кельвина, Фаренгейта.
// использование собственных типов
package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func kelvinToFahrenheit(k kelvin) fahrenheit {
	return fahrenheit(celsiusToFahrenheit(kelvinToCelsius(k)))
}

func main() {
	fmt.Printf("233° K is %.2f° C\n", kelvinToCelsius(233))
	fmt.Printf("0° K is %.2f° F\n", kelvinToFahrenheit(0))
}
