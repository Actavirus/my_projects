// тема: Структуры и методы — объектно-ориентированный подход в Golang
package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64 // обозначение координат d/m/s
	h       rune    // направление: W, S, E, N - север, запад, юг, восток.
}
type location struct {
	lat, long float64
}
type world struct {
	radius float64
}

func (c coordinate) decimal() float64 { // метод decimal конвертирует координаты d/m/s в десятичные градусы.
	sign := 1.0
	switch c.h { // для обозначения W и S - т.е. отрицательные данные
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func newLocation(lat, long coordinate) location { // функция конструктора под названием newLocation. В Go нет языковых инструментов для конструкторов (в отличие от Python, Ruby или PHP)
	return location{lat.decimal(), long.decimal()}
}
func (w world) distance(p1, p2 location) float64 { // вычисление расстояния через Сферическую теорему косинусов.
	// TODO: Добавить математические выражения, используя w.radius
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Использует поле радиуса world
}
func rad(deg float64) float64 { // rad конвертирует градусы в радианы.
	return deg * math.Pi / 180
}
func main() {
	// Прикрепление методов к структурам в Golang
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal()) // выводит: -4.5895 137.4417

	// Функции конструктора в Golang
	curiosity := location{lat.decimal(), long.decimal()}
	fmt.Println(curiosity)               // выводит: {-4.5895 137.4417}
	curiosity2 := newLocation(lat, long) // т.е., мы применяем функцию newLocation
	fmt.Println(curiosity2)              // выводит: {-4.5895 137.4417}

	// Классы в Golang
	// В Go нет понятия class, как в других популярных языках программирования вроде Python, Ruby или Java. Однако структуры в сочетании с несколькими методами можно использовать для тех же целей.
	var mars = world{radius: 3389.5} // Вместо объявления 3389,5 как константы, используйте тип world, чтобы объявить Марс одним из многих возможных миров.
	spirit := location{-25.5132, 94.2358}
	opportunity := location{20.1548, 264.1987}
	dist := mars.distance(spirit, opportunity) // Использует метод distance для mars
	fmt.Printf("%.2f km\n", dist)              // выводит: 10016.36 km
	var earth = world{radius: 6371.0}
	dist2 := earth.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist2) // выводит: 18827.04 km

	// задача из учебника: Используя Листинги 1, 2 и 3, напишите программу, что объявляет location для каждой локации из Таблицы 1. Код должен выводить каждую локацию в десятичных градусах. - моё решение задачи.
	coordSpirit := local{
		location2{14, 34, 6.2, 'S'},
		location2{175, 28, 21.5, 'E'},
	}
	// coordOpportunity := local{
	// 	location2{1, 56, 46.3, 'S'},
	// 	location2{354, 28, 24.2, 'E'},
	// }
	// coordCuriosity := local{
	// 	location2{4, 35, 22.2, 'S'},
	// 	location2{137, 26, 30.1, 'E'},
	// }
	coordInSight := local{
		location2{4, 30, 0.0, 'N'},
		location2{135, 54, 0, 'E'},
	}
	g := perevod(coordInSight.Lat)
	h := perevod(coordInSight.Long)
	spirit2 := table{"Spirit", "Columbia Memorial Station", coordSpirit, coordSpirit}
	opportunity2 := table2{"Opportunity", "Challenger Memorial Station", local{location2{1, 56, 46.3, 'S'}, location2{354, 28, 24.2, 'E'}}}
	curiosity3 := table3{"Curiosity", "Bradbary Landing", location2{4, 35, 22.2, 'S'}, location2{137, 26, 30.1, 'E'}}
	insight := table4{"InSight", "Elysium Planitia", g, h}
	fmt.Printf("%s %s %.2f %.2f\n", spirit2.Module, spirit2.Landing, perevod(coordSpirit.Lat), perevod(coordSpirit.Long)) // Spirit Columbia Memorial Station 14.57 175.47
	fmt.Printf("%s %s %.2f %.2f\n", opportunity2.Module, opportunity2.Landing, perevod(opportunity2.LocalLatLong.Lat), perevod(opportunity2.LocalLatLong.Long))
	fmt.Printf("%s %s %.2f %.2f\n", curiosity3.Module, curiosity3.Landing, perevod(curiosity3.Lat), perevod(curiosity3.Long))
	bytes, _ := json.MarshalIndent(insight, "---", "|")
	fmt.Println(string(bytes))

	// второе решение задачи через использование типа struct в бОльшем формате
	table := []table5{
		{"Spirit", "Columbia Memorial Station", location2{14, 34, 6.2, 'S'}, location2{175, 28, 21.5, 'E'}},
		{"Opportunity", "Challenger Memorial Station", location2{1, 56, 46.3, 'S'}, location2{354, 28, 24.2, 'E'}},
		{"Curiosity", "Bradbary Landing", location2{4, 35, 22.2, 'S'}, location2{137, 26, 30.1, 'E'}},
		{"InSight", "Elysium Planitia", location2{4, 30, 0.0, 'N'}, location2{135, 54, 0, 'E'}},
	}
	fmt.Println(table)

}

type local struct {
	Lat  location2
	Long location2
}
type location2 struct {
	D, M, S float64 // d - градус, m - минута, s - секунда
	H       rune
}
type table struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     local
	Long    local
}
type table2 struct {
	Module       string `json:"The module:"`
	Landing      string `json:"Landing site"`
	LocalLatLong local
}
type table3 struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     location2
	Long    location2
}
type table4 struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     float64
	Long    float64
}
type table5 struct {
	Module  string
	Landing string
	Lat     location2
	Long    location2
}

// В одной минуте 60 секунд («), в одном градусе 60 минут (‘)
func perevod(a location2) float64 { // это функция.
	b := 1.0
	if a.H == 'S' {
		b = -1.0
	}
	return b * (a.D + a.M/60 + a.S/3600)
}
