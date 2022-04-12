// программа сообщает, сдал ли пользователь экзамен
// инкапсуляция
// использование set и get методов
package main

import (
	"bufio"
	"fmt"
	"github/headfirstgo/calendar"
	"github/headfirstgo/geo"
	"github/headfirstgo/magazine"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil { // если при чтении ввода произошла ошибка, она будет возвращена функцией
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil { // также возвращаются ошибки, которые могут возникнуть при преобразовании строки в float64
		return 0, err
	}
	return number, nil
}
func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func main() {
	/*	fmt.Print("Введите процент правильных ответов:")
		grade, err := getFloat() // вызываем функцию getFloat для получения grade
		if err != nil {          // если фукнция вернула ошибку, программа сообщает об этом и завершается
			log.Fatal(err)
		}
		var status string
		if grade >= 60 {
			status = "Ты успешно сдал экзамен"
		} else {
			status = "Экзамен не пройден"
		}
		fmt.Println(status)*/
	//fmt.Println(os.Args[1:])
	fmt.Println(sum(4, 1, 2))
	fmt.Println(sum(7, 9))
	var s magazine.Subscriber // имя типа теперь должно иметь префикс с именем пакета
	s.Rate = 4.99
	fmt.Printf("%f\n", s)

	// упражнение из учебника
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)

	// пробуем использовать дополнительный пакет calendar
	date := calendar.Date{} // здесь calendar - необходимо указать пакет, и которого импортируется тип; Создаётся новое значение Date
	// поля Date задаются напрямую
	//	date.year = 2019 // Внимание! имена полей начинаются с букв нижнего регистра. Это значит, что к полям не удастся обратиться напрямую или через литералы.
	//	date.month = 14  // Внимание! имена полей начинаются с букв нижнего регистра. Это значит, что к полям не удастся обратиться напрямую или через литералы.
	//	date.day = 50    // Внимание! имена полей начинаются с букв нижнего регистра. Это значит, что к полям не удастся обратиться напрямую или через литералы.
	// использование set и get методов
	err := date.SetYear(2019) // используется set-метод
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5) // используется set-метод
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27) // используется set-метод
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year()) // значения, возвращаемые get-методами.
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	//	date = calendar.Date{year: 0, month: 0, day: -2}	// поля Date задаются с помощью литералов

}
