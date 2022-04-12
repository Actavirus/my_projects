package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) { // здесь "s subscriber" - объявляется один параметр с типом subscriber
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) subscriber { // возвращает значение subscriber
	var s subscriber // создаётся новое значение subscriber
	// заполняются поля структуры
	s.name = name
	s.rate = 5.99
	s.active = true
	return s // возвращает subscriber
}
func main() {
	subscriber1 := defaultSubscriber("Aman Singh") // создаёт subscriber с заданным значением name
	subscriber1.rate = 4.99                        // использует заданное значение rate
	printInfo(subscriber(subscriber1))             // вывод значений полей
	subscriber2 := defaultSubscriber("Beth Ryan")  // создаёт subscriber с заданным значением name
	printInfo(subscriber(subscriber2))             // вывод значений полей
}
