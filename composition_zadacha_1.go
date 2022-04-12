// задача из учебника: Напишите программу со структурой gps для Системы Глобального Позиционирования (GPS). Данная структура должна состоять из текущей местности, места назначения и мира. - мой вариант решения.
package main

import (
	"fmt"
	. "math"
)

type GPS struct { // структура GPS состоит из текущей местности, места назначения и мира.
	locationA locationAll
	locationB locationAll
	planet    worldAll
}
type locationAll struct {
	localityName string
	lat, long    float64
}

// Тип world должен имплементировать метод для расстояния, используемый в уроке о структурах и методах.
type worldAll struct {
	planet string
	rad    float64
}

// Реализация метода description для типа location возвращает строку, содержащую название, широту и долготу.
func (l locationAll) description() string {
	return fmt.Sprintf("%s %.2f° %.2f°\n", l.localityName, l.lat, l.long)
}

// метод для расстояния, используемый в уроке о структурах и методах.
func (w worldAll) distance(a, b locationAll) float64 {
	s1, c1 := Sincos(rad(a.lat))
	s2, c2 := Sincos(rad(b.lat))
	clong := Cos(rad(a.long - b.long))
	return w.rad * Acos(s1*s2+c1*c2*clong)
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * Pi / 180
}
func (g GPS) distance() float64 {
	return g.planet.distance(g.locationA, g.locationB)
}
func (g GPS) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.locationB.description())
}

var (
	earth = worldAll{planet: "Earth", rad: 6371.0}
	mars  = worldAll{planet: "Mars", rad: 3389.5}
)

type rover struct {
	GPS
}

func main() {
	localisationA := locationAll{localityName: "Position A", lat: -4.5895, long: 137.4417}
	localisationB := locationAll{"Position B", 4.5, 135.9}

	GPS := GPS{
		planet:    mars,
		locationA: localisationA,
		locationB: localisationB,
	}

	curiosity := rover{
		GPS: GPS,
	}

	fmt.Println(curiosity.message())

}
