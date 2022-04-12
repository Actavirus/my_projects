// тема - Конвертирование типов данных в Golang
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// объединение строк
	countdown := "Launch in T minus " + "10 seconds"
	fmt.Println(countdown)

	// конвертирование данных в Golang string в int
	value1 := "15"                    // string
	value2, _ := strconv.Atoi(value1) // string -> int. Обязательно присутствует в синтаксисе 2 значения, которые передаёт функция strconv.Atoi
	value3 := strconv.Itoa(value2)    // int -> string
	fmt.Printf("%f %f %f\n", value1, value2, value3)
	value4 := fmt.Sprintf("", value2) // ещё один способ конвертации int -> string
	fmt.Println(value4)

	// конвертация булевых значений в строку
	launch := true                          // задаётся булево значение
	launchText := fmt.Sprintf("%v", launch) // используется функция Sprintf для конвертации булевой переменной launch в текст.
	fmt.Println(launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
	// обратная конвертация
	launch2 := (yesNo == "yes")
	fmt.Printf("%f\n", launch2)

	// задача из учебника: Как можно конвертировать булев тип в целое число, чтобы 1 соответствовала значению true, а 0 — false
	var launchInt int
	if launch {
		launchInt = 1
	} else {
		launchInt = 0
	}
	fmt.Println(launchInt)

	// задача из учебника: Напишите программу, что конвертирует строки в булевы значения: «true», «yes» или «1» соответствуют значению true; «false», «no» или «0» соответствуют значению false; Для других значений выводит сообщение об ошибке.
	launchString := fmt.Sprintf("%v", launch)
	var launchBool bool
	switch launchString {
	case "true", "yes", "1":
		launchBool = true
	case "false", "no", "0":
		launchBool = false
	default:
		fmt.Println("Ошибка")
	}
	fmt.Println(launchBool)
}
