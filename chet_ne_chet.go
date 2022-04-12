// программа получает от пользователя целое число (int) и выдаёт сообщение о том, чётное было число или не чётное
package main

import (
	"fmt"
	"reflect"
)

func proverka(number int) {
	fmt.Println(number)
	if number%2 == 0 { // проверка условия деления на 2 без остатка
		fmt.Println("Число ", number, " чётное")
	} else {
		fmt.Println("Число ", number, " не чётное")
	}
}
func main() {
	fmt.Println("Введите целое число")
	var input int
	fmt.Scanf("%d", &input)            // Scanf заполняет переменную input числом, введённым пользователем
	fmt.Println(reflect.TypeOf(input)) // выводим тип принятого значения
	if input%2 == 0 {
		fmt.Println("Введённое число чётное")
	} else {
		fmt.Println("Введённое число не чётное")
	}
}
