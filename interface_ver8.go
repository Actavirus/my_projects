// тема: Реализация интерфейсов в Golang (https://nuancesprog.ru/p/11974/)
package main

import (
	"fmt"
	"time"
)

type Printer interface { // интерфейс определяет метод Print(). Данный метод представляет действие или поведение, которые могут реализовывать другие объекты.
	Print()
}
type User struct { // объект, реализующий интерфейс Printer
	name     string
	age      int
	lastName string
}
type Document struct { // объект, реализующий интерфейс Printer
	name         string
	documentType string
	date         time.Time
}

func (d Document) Print() { // метод Print() для структуры Document
	fmt.Printf("Document name: %s, type: %s, date: %s \n", d.name, d.documentType, d.date)
}
func (u User) Print() { // метод Print() для структуры User
	fmt.Printf("Hi I am %s %s and I am %d years old \n", u.name, u.lastName, u.age)
}

// Предположим, нам нужно написать новый метод, выводящий подробности этих двух структур. Для этого можно использовать имеющийся интерфейс:
func Process(obj Printer) { // Эта функция получает в качестве аргумента любые объекты, реализующие указанный интерфейс.
	obj.Print()
}

func main() {
	// 1. Определение простого интерфейса
	u := User{name: "John", age: 24, lastName: "Smith"}
	doc := Document{name: "doc.csv", documentType: "csv", date: time.Now()}
	Process(u)   // выводит: Hi I am John Smith and I am 24 years old
	Process(doc) // выводит: Document name: doc.csv, type: csv, date: 2022-02-09 20:00:56.959993 +0300 MSK m=+0.006000301

}
