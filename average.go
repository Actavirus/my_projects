// average вычисляет среднее значение чисел, взятые их другого файла (data.txt)
package main

import (
	"fmt"
	"log"

	"github/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt") // загружает файл data.txt, разбирает содержащиеся в нем числа и сохраняет массив
	if err != nil {                                // если произошла ошибка, программа сообщает о ней и завершает работу
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers)) // здесь получаем длину массива (len) и переводим в тип float64
	fmt.Printf("Среднее арифметическое массива:%.2f\n", sum/sampleCount)
}
