// фукнция sum принимает срез (сегмент) чисел и складывает их вместе
// программа работает совместно с функцией sum, которая суммирует числа из полученного сегмента (среза)
// ввод данных идёт совместно с командой запуска программы, например: >go run variant_2.go 56 58 54.2 59.02
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("сумма введённых чисел равна %.2f\n", sum(numbers...))
}
