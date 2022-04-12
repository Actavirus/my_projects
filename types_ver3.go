// тема - Целые цисла в Go (integer). Правильный выбор типа.
package main

import "fmt"

func main() {
	year := "kripto"
	fmt.Printf("Type %t for %v\n", year, year)
	fmt.Printf("Type %T for %v\n", year, year) // Внимание! обратите внимание на вывод сообщения в терминале.
	fmt.Printf("Type %T for %[1]v\n", year)    // Внимание! вместо повторения переменной можно указать [1], чтобы первый аргумент использовался повторно.
	// шестнадцатеричные значения в Go (цвета в CSS)
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5           // эквивалентно: var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("%x %x %X\n", red, green, blue)              // Внимание! обратите внимание на верхний регистр специального символа.
	fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue) // %02x - нулевой отступ
	// целочисленное переполнение в Go
	var yellow uint8 = 255 // у uint8 диапазон от 0 до 255
	yellow++               // прибавляем ещё одну единицу
	fmt.Println(yellow)    // значения выше 255 для uint8 возвращаются к 0.

	// расчёт количества дней полёта до звезды Альфа Центавра (пример использования больших чисел)
	const lightSpeed = 299792 // км/сек
	const secondsPerDay = 86400
	var distance int64 = 41.3e12 // 41300000000000
	fmt.Println("Расстояние до Альфа Центавра составляет", distance, "км")
	days := distance / lightSpeed / secondsPerDay
	fmt.Println("Это", days, "дня полёта на скорости света.")
}
