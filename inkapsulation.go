// работа по теме Инкапсуляция: механизм защиты полей структурного типа от недействительных данных (или практика сокрытия данных в одной части программы от кода в другой части).
// программа-календарь: внесение событий на конкретную дату, проверка правильности внесения данны (даты); использование Set-методов.
package main

import (
	"errors" // позволяет создавать значения ошибок.
	"fmt"
	"log"  // для вывода сообщения об ошибке и завершения программы.
	"time" // необходимо импортировать пакет "time", чтобы использовать тип time.Time.
)

type Date struct {
	Year, Month, Day int
	Time
}
type Time struct {
	Hour, Minute int
}

func (d *Date) SetYear(year int) error { // здесь *Date - здесь нужен указатель на получателя, чтобы обновлялась не копия, а исходное значение; year - получает значение, которое болжно быть присвоено полю.
	if year < 1 {
		return errors.New("invalid year") // если год недействителен, возвращается признак ошибки...
	}
	d.Year = year // ...в противном случае присваивается значение поля... Здесь d - автоматически получает значение по указателю; year - теперь обновляется исходное значение, а не копия.
	return nil    //...и возвращается ошибка "nil"
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}                     // создание Date.
	var timeNow time.Time = time.Now() // метод time.Now возвращает значение time.Time, представляющее текущую дату и время [год, месяц, день, час, минута, секунда и т.д.]
	//yearNow := timeNow.Year()          // у значений time.Time имеется метод Year, который возвращает текущий год.
	err := date.SetYear(timeNow.Year()) // здесь err - сохраняются любый ошибки. Значение (0) не действительно!
	if err != nil {
		log.Fatal(err) // если значение недопустимо, программа выводит ообщение об ошибке и завершается.
	}
	//date.SetYear(1990) // задаёт значение поля Year при помощи метода; автоматически преобразуется в указатель.
	err = date.SetMonth(int(timeNow.Month())) // подставляется значение текущего месяца (с помощью метода time)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(timeNow.Day())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year) // выводит поле Year.
	fmt.Println(date.Month)
	fmt.Println(date.Day)
	fmt.Println(date)
}
