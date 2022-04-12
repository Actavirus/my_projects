// структура для хранения данных подписчиков журнала
package main

import "fmt"

type subscriber struct { // объявление типа с именем subscriber
	name   string  // поле name для хранения строки
	rate   float64 // поле rate для хранения float64
	active bool    // поле active для хранения bool
}

func main() {
	var subscriber1 subscriber // объявляем переменную subscriber1 типа subscriber
	subscriber1.name = "Aman Singh"
	subscriber1.rate = 4.99
	subscriber1.active = true
	fmt.Println(subscriber1.name)
	fmt.Println(subscriber1.rate)
	fmt.Println(subscriber1.active)
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	subscriber2.rate = 4.55
	subscriber2.active = true
	fmt.Println(subscriber2.name)
	fmt.Println(subscriber2.rate)
	fmt.Println(subscriber2.active)
}
