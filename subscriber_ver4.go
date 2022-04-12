// передача больших структур с помощью указателей
// ниже приведена программа с функциями defaultSubscriber и функция print Info. Обе функции используют только указатели на структуру, это позволяет использовать и хранить только одну копию структуры и не загружать память.
package main

import "fmt"

type subsrciber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subsrciber) { // здесь *subscriber - изменяется для передачи указателя
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) *subsrciber { // здесь *subscriber - изменяется для возвращения указателя
	var s subsrciber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // возвращает указатель на структуру вместо самой структуры
}
func applyDiscount(s *subsrciber) {
	s.rate = 4.99
}
func main() {
	subsrciber1 := defaultSubscriber("Aman Singh") // здесь subscriber1 - это уже не структура, а только указатель на структуру
	applyDiscount(subsrciber1)                     // внимание! так как это структура, оператор & не нужен
	printInfo(subsrciber1)
	subsrciber2 := defaultSubscriber("Beth Ryan")
	printInfo(subsrciber2)
}
