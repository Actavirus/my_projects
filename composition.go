// тема: Композиция и встраивание методов в Golang (В данном уроке мы рассмотрим композицию и встраивание на примере выдуманного прогноза погоды от REMS.)
package main

import "fmt"

type report struct {
	sol         int         // указания текущего марсианского дня (его называют сол)
	temperature temperature // самых низких и высоких температур. Поле temperature является структурой типа temperature
	location    location3
}
type temperature struct {
	high, low celsius2
}
type location3 struct {
	lat, long float64
}
type celsius2 float64

func (t temperature) average() celsius2 {
	return (t.high + t.low) / 2
}
func (r report) average() celsius2 { // Если вы хотите показать среднюю температуру напрямую через тип report, то напишите метод, что встраивает реальную имплементацию.
	return r.temperature.average()
}

type report2 struct {
	sol         int
	temperature // Тип temperature встроен в report
	location3
}
type sol int
type report3 struct {
	sol
	temperature
	location3
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}
func main() {
	// Композиция структур в Golang
	broadbury := location3{-4.5894, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: broadbury}
	fmt.Printf("%+v\n", report)                                 // выводит: {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5894 long:137.4417}}. Внимание! знак "+" нужен для отображения наименования полей структуры.
	fmt.Printf("a balmy %v° C\n", report.temperature.high)      // выводит: a balmy -1° C
	fmt.Printf("average %v° C\n", t.average())                  // выводит: average -39.5° C
	fmt.Printf("average %v° C\n", report.temperature.average()) // При создании отчета о погоде метод average доступен через объединение с полем temperature
	fmt.Printf("average %v° C\n", report.average())             // выводит: average -39.5° C

	// Встраивание методов в Golang
	report2 := report2{ // моё мнение: очень не разумно давать схожее с типом наименование переменной.
		sol:         15,
		location3:   location3{lat: -4.5895, long: 137.4417},
		temperature: temperature{high: -1.0, low: -78.0}, // Все методы типа temperature автоматически делаются доступными через тип report
	}
	fmt.Printf("average %v° C\n", report2.average()) // выводит: average -39.5° C
	fmt.Printf("%v° C\n", report2.high)              // выводит: -1° C
	report2.high = 32.0
	fmt.Printf("%v° C\n", report2.temperature.high) // выводит: 32° C
	report3 := report3{sol: 15}
	fmt.Println(report3.sol.days(1446)) // выводит: 1431. К любым методам, объявленным в типе sol, можно получить доступ через поле sol или через тип report.
	fmt.Println(report3.days(1446))     // выводит: 1431
	fmt.Println(report2.lat)

	// Столкновение, или коллизия названий в Go - непонятная для меня тема, впрочем, далее я разберусь с ней.

	// Есть ли наследование в Golang?

}
