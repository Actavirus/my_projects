// программа получает от пользователя два зачения, проверяет на условие "деление на ноль", и делит одно на другое.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("Нельзя делить на 0")
	}
	return dividend / divisor, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin) //создаём bufio.Reader для чтения ввода с клавиатуры
	//получаем первое значение от пользователя
	fmt.Print("Укажите первое значение:")  //запрашиваем первое значение
	input1, err := reader.ReadString('\n') //прочитать данные, введенные пользователем до нажатия Enter
	if err != nil {                        //если произошла ошибка, программа выводит сообщение и завершается, например ввели тип string
		log.Fatal(err)
	}
	input1 = strings.TrimSpace(input1)            //удаление символа новой строки \n
	guess1, err := strconv.ParseFloat(input1, 64) //входная строка преобразуется в число типа float64 (для типа int используется функция Atoi)
	//получаем второе значение от пользователя
	fmt.Print("Укажите второе значение:")  //запрашиваем второе значение
	input2, err := reader.ReadString('\n') //прочитать данные, введенные пользователем до нажатия Enter
	if err != nil {                        //если произошла ошибка, программа выводит сообщение и завершается, например ввели тип string
		log.Fatal(err)
	}
	input2 = strings.TrimSpace(input2)            //удаление символа новой строки \n
	guess2, err := strconv.ParseFloat(input2, 64) //входная строка преобразуется в число типа float64 (для типа int используется функция Atoi)
	//полученные значения отправляем в функцию divide, проверяем и получаем ответные значения
	quotient, err := divide(guess1, guess2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.2f\n", quotient)
	}
}
