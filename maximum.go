// использование функций с переменным количеством аргументов
package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 { // здесь ... - получает любое количество аргументов float64
	max := math.Inf(-1)              // начинает с очень низкого значения
	for _, number := range numbers { // обработка всех аргументов в переменной части
		if number > max {
			max = number // находит наибольшее значение среди аргументов
		}
	}
	return max
}
func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
}
