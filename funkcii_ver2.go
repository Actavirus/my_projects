// тема - Функции первого класса, замыкания и анонимные функции в Golang
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

var f = func(message string) { // присваивает анонимную функцию переменной. Внимание! переменная в области видимости всего пакета (но можно сделать объявление внутри функции).
	fmt.Println(message)
}

func fakeSensor() kelvin { // функция рандомного значения температуры (с использованием встроенной функции time для неповторения)
	value := time.Now().Second()          // объявляем переменную с значением текущей секунды
	return kelvin(rand.Intn(value) + 150) // используя значение текущей секунды определяем рандомное значение
}
func realSensor() kelvin { // функция-имитация подключенного датчика температуры
	return 0
}

func measureTemperature(samples int, sensor func() kelvin) kelvin { // measureTemperature принимает функцию в качестве второго параметра
	//func measureTemperature(samples int, sensor sensor) {	// другой вариант записи объявления функции
	var sum kelvin
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		sum = sum + k
		time.Sleep(300 * time.Millisecond) // где 300 - это количество Millisecond
	}
	return sum / kelvin(samples)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // объявляет и возвращает анонимную функцию
		return s() + offset
	}
}
func main() {
	// функции первого класса
	sensor := fakeSensor //	присваивает функцию переменной. Внимание! переменная присваивает функцию, а не результат вызова функции, поэтому скобок () нет.
	fmt.Println(sensor())

	sensor = realSensor // то же самое, переменная присваивает функцию, а не результат вызова функции.
	fmt.Println(sensor())

	// передача функции другой функции
	number := 5
	result := measureTemperature(number, fakeSensor)
	fmt.Printf("Средняя температура по %d показаниям равна: %.2f K\n", number, float64(result))

	// анонимные функции Golang.
	// анонимные фукнции - это функции без названия.
	f("My name's Arthur King") // вызов функции.
	// можно объявить и задействовать анонимную функцию другим образом:
	func() { // обявляем анонимну функцию
		fmt.Println("Functions anonymous")
	}() // вызов функции

	// замыкания (Closures)
	sensor = calibrate(realSensor, 5)
	fmt.Println(sensor()) // выводит: 5
}
