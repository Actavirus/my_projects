// задача из учебника: Используйте метод distance из Листинга 4 и напишите программу, что определяет расстояние между каждой парой мест посадки из Таблицы 1. Какие места посадки находятся ближе всего друг к другу? Какие места дальше всего друг от друга? - мой варинат решения
package main

import (
	"fmt"
	. "math" // знак "." - для того, чтобы каждый не писать префикс "math" перед Cos/Sin
)

// const (
// 	earth = 6371.0
// )
type dataTable struct {
	module  string
	landing string
	lat     coord
	long    coord
}
type coord struct {
	d, m, s float64
	h       rune
}

type planet struct {
	radius float64
}

func (p planet) distance(d1, d2 dataTable) float64 {
	s1, c1 := Sincos(toRad(toGrad(d1.lat)))
	s2, c2 := Sincos(toRad(toGrad(d2.lat)))
	clong := Cos(toRad(toGrad(d1.long) - toGrad(d2.long)))
	return p.radius * Acos(s1*s2+c1*c2*clong)
}

func toRad(deg float64) float64 { // функция конвертирует градусы в радианы.
	return deg * Pi / 180
}
func toGrad(c coord) float64 { // функция переводит "ГрадусМинутаСекунда" в "Градусы"
	sign := 1.0
	switch c.h {
	case 'S', 'W':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func outputResult(value1, value2 dataTable, p planet) {
	fmt.Printf("%15s%8v|%15s%8v|%15.2f\n", value1.module, "", value2.module, "", p.distance(value1, value2))
}

func main() {
	spirit := dataTable{"Spirit", "Columbia Memorial Station", coord{14, 34, 6.2, 'S'}, coord{175, 28, 21.5, 'E'}}
	opportunity := dataTable{"Opportunity", "Challenger Memorial Station", coord{1, 56, 46.3, 'S'}, coord{354, 28, 24.2, 'E'}}
	curiosity := dataTable{"Curiosity", "Bradbury Landing", coord{4, 35, 22.2, 'S'}, coord{137, 26, 30.1, 'E'}}
	insight := dataTable{"InSight", "Elysium Planitia", coord{4, 30, 0.0, 'N'}, coord{135, 54, 0, 'E'}}
	// var value = []dataTable{spirit, opportunity, curiosity, insight}
	// fmt.Println(value)
	// var mercury, venus, earth, mars, jupiter, saturn, uranus, neptune = planet{2439.7}, planet{6051.8}, planet{6371.0}, planet{3389.5}, planet{699911.0}, planet{58232}, planet{25362}, planet{24622}
	earth, mars := planet{radius: 6371.0}, planet{3389.5}
	fmt.Printf("%15s%8s|%15s%8s|%15s\n", "Точка А", "", "Точка В", "", "Расстояние, км")
	fmt.Println("-----------------------|-----------------------|---------------")
	outputResult(spirit, opportunity, mars)
	outputResult(spirit, curiosity, earth)
	outputResult(spirit, insight, earth)
	outputResult(opportunity, curiosity, earth)
	outputResult(opportunity, insight, earth)
	outputResult(curiosity, insight, earth)

}
