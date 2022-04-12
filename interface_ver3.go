// тема: Интерфейсы в Golang
package main

import (
	"fmt"
	"strings"
	"time"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk()) // здесь ToUpper - меняет строчные буквы на заглавные
	fmt.Println(louder)
}

type crater struct {
	c int
}

func (c crater) talk() string {
	return strings.Repeat("one", c.c) // здесь Repeat - встроенный метод, который повторяет (string)  заданное количество раз (int)
}

type starship struct {
	laser
}
type rover1 struct {
	r int
}

func (r rover1) talk() string {
	return strings.Repeat("whir", r.r)
}
func stardate(t time.Time) float64 { // stardate возвращает выдуманное измерение времени для указанной даты. Функция stardate из Листинга 8 ограничивается земными датами.
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type stardater interface {
	YearDay() int
	Hour() int
}

func stardate2(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668) // Марсианской год состоит из 668 дней
}
func (s sol) Hour() int {
	return 0 // Неизвестный час
}

type location4 struct { // location с широтой и долготой в десятичных градусах
	lat, long float64
}

func (c coordinate5) String() string { // String форматирует координаты DMS
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}
func (l location5) String() string { // String форматирует location с широтой и долготой
	return fmt.Sprintf("%v, %v\n", l.lat, l.long)
}

type location5 struct { // location с широтой и долготой в десятичных градусах
	lat, long coordinate5
}

type coordinate5 struct {
	d, m, s float64
	h       rune
}

func main() {
	// Тип interface в Golang
	var t interface { // объявление переменной с типом interface
		talk() string
	}
	t = martian{}
	fmt.Println(t.talk()) // выводит: nack nack
	t = laser(3)
	fmt.Println(t.talk()) // выводит: pewpewpew
	shout(martian{})      // выводит: NACK NACK
	shout(laser(5))       // выводит: PEWPEWPEWPEWPEW
	t = crater{}
	shout(crater{7}) // выводит: ONEONEONEONEONEONEONE
	s := starship{laser: 3}
	fmt.Println(s.talk()) // выводит: pewpewpew
	shout(s)              // выводит: PEWPEWPEW
	shout(rover1{8})      // выводит: WHIRWHIRWHIRWHIRWHIRWHIRWHIRWHIR

	// Примеры использования интерфейсов в коде на Golang. В следующем листинге выводится выдуманная дата, что состоит из дня года и часа дня.
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day)) // выводит: 1219.2 Curiosity has landed
	s1 := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate2(s1)) // выводит: 1086.0 Happy birthday

	// Удовлетворение требований интерфейса в Go
	curiosity := location4{-4.5895, 137.4417}
	fmt.Println(curiosity) // выводит: -4.6, 137.4417
	// Напишите метод String для типа coordinate и location и используйте его для отображения координат в более читабельном формате - мой вариант решения.
	Elysium2 := location5{
		lat:  coordinate5{4, 30, 0.0, 'N'},
		long: coordinate5{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at ", Elysium2) // выводит: Elysium Planitia is at  4°30'0.0" N, 135°54'0.0" E
}
