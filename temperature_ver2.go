// программа для перевода температуры из Цельсия в Фаренгейты и обратно - моя версия решения задачи.
package main

import "fmt"

type celsius float64
type fahrenheit float64

func fakeSensor(value string) {
	if value == "c" {
		for i := 40; i <= 60; i += 5 {
			fmt.Printf("%12.2d °C|%12.2f °F\n", i, celsius(i).fahrenheit())
		}
	} else if value == "f" {
		for i := 40; i <= 60; i += 5 {
			fmt.Printf("%12.2d °C|%12.2f °F\n", i, fahrenheit(i).celsius())
		}
	}
}
func (f fahrenheit) celsius() celsius {
	c := (f - 32.0) * 5.0 / 9.0
	return celsius(c)
}

func (c celsius) fahrenheit() fahrenheit {
	f := (c * 9.0 / 5.0) + 32.0
	return fahrenheit(f)
}

func main() {
	fmt.Println("===============|===============")
	fmt.Printf("%15v|%15v\n", "Temperature C", "Temperature F")
	fmt.Println("===============|===============")
	fakeSensor("c")
	fmt.Println("===============|===============")
	fmt.Printf("%15v|%15v\n", "Temperature F", "Temperature C")
	fmt.Println("===============|===============")
	fakeSensor("f")
	fmt.Println("===============|===============")
}
