// average вычисляет среднее значение чисел, которые пользователь сам вводит следующим образом: >go run average_ver2 52.4 45 65.02 58 69.0
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]             // получение сегмента строк со всеми элементами os.Args, кроме первого
	var sum float64 = 0                  // объявление переменной для хранения суммы чисел
	for _, argument := range arguments { // обрабатываем каждый аргумент командной строки
		number, err := strconv.ParseFloat(argument, 64) // строка преобразуется в float64
		if err != nil {                                 // если при преобразовании произошла ошибка, программа сообщает о ней и завершает работу
			log.Fatal(err)
		}
		sum += number // число прибавляется к сумме
	}
	sampleCount := float64(len(arguments))                                       // длина сегмента arguments используется как количество значений данных
	fmt.Printf("Среднее арифметическое введённых чисел:%.2f\n", sum/sampleCount) // вычисление среднего значения и его вывод
}
