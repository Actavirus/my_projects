// average_ver4 так же считает среднее арифметическое, но использует для этого функцию average (как в ver3) и сегмент значений (как в ver2)
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 { //получаем любое количество аргументов float64
	var sum float64 = 0              // создание переменной для хранения суммы аргументов
	for _, number := range numbers { // обработать каждый аргумент из переменной части
		sum += number // значение аргумента прибавляется к сумме
	}
	return sum / float64(len(numbers)) // среднее значение вычисляется делением суммы на количество аргументов
}
func main() {
	arguments := os.Args[1:] // получение сегмента строк со всеми элементами os.Args, кроме первого
	var numbers []float64    // в этом сегменте будут храниться числа, для которых вычисляется среднее
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number) // преобразованное число присоединяется к сегменту
	}
	fmt.Printf("Среднее арифметическое чисел: %.2f\n", average(numbers...)) // здесь ... - это необходимо, т.к. функция average рассчитывает получить один или несколько аргументов float64, а не сегмент значений float64
}
