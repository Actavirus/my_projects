// программа запрашивает у пользователя температуру в градусах по Фаренгейту и преобразует введенное значение в градусы по Цельсию
package main

import (
	"fmt"
	"keyone/keytwo/keyboard"
	"log"
)

func main() {
	fmt.Print("Введите значение градуса по Фаренгейту: ")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (farenheit - 32) * 5 / 9
	fmt.Print(farenheit)
	fmt.Printf(" градусов по Фаренгейту это %0.2f градуса по Цельсию \n", celsius)
}
