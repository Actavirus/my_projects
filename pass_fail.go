// программа сообщает, сдал ли пользователь экзамен
package main

import (
	"fmt"
	"keyone/keytwo/keyboard"
	"log"
)

func main() {
	fmt.Print("Введите процент правильных ответов:")
	grade, err := keyboard.GetFloat() // вызываем функцию getFloat для получения grade
	if err != nil {                   // если фукнция вернула ошибку, программа сообщает об этом и завершается
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "Ты успешно сдал экзамен"
	} else {
		status = "Экзамен не пройден"
	}
	fmt.Println(status)

}
