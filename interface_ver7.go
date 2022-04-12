// Использование интерфейсов в Go (https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go-ru)
package main

import (
	"fmt"
	. "math"
)

type Stringer interface {
	string() string
}
type Article struct {
	Title  string
	Author string
}

func (a Article) string() string { // Метод — это специальная функция, относящаяся к определенному типу в Go. В отличие от функции, метод можно вызвать только из экземпляра типа, для которого он определен.
	return fmt.Sprintf("The %q article was written by %s.\n", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) string() string {
	return fmt.Sprintf("The %q Book was written by %s It has %v pages.\n", b.Title, b.Author, b.Pages)
}

func Print(s Stringer) {
	fmt.Println(s.string())
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return Pi * Pow(c.Radius, 2)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Height * s.Width
}

type Sizer interface {
	Area() float64
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}
func main() {
	// Определение поведения
	// В отличие от структур, для интерфейсов мы определяем только поведение, то есть «что может делать этот тип».
	// Один из самых часто используемых интерфейсов стандартной библиотеки Go — интерфейс fmt.Stringer
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	fmt.Println(a.string()) // выводит: The "Understanding Interfaces in Go" article was written by Sammy Shark.

	// Определение интерфейса
	Print(a) // выводит: The "Understanding Interfaces in Go" article was written by Sammy Shark.
	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b) // выводит: The "All About Go" Book was written by Jenny Dolphin It has 25 pages.

	// Различные варианты поведения в интерфейсе.
	c := Circle{Radius: 10}
	s := Square{Height: 10, Width: 10}
	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l) // выводит: {Width:10 Height:10} is the smallest
}
