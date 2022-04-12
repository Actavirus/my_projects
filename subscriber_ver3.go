// прорамма, которая заменяет одно из значений структуры на другое значение с помощью функции
// обратите внимание, что функции applyDiscount передаётся не сама структура, а лишь указатель на неё
package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *subscriber) { // здесь *subacriber - получает указатель на структуру, а не саму структуру
	s.rate = 4.99 // обновляет поле структуры. Обратите внимание, что при заполнении поля rate перед ней не ставится знак *, т.к. точечная запись при обращении к полям работает как с указателями на структуры, так и с самими структурами.
}
func main() {
	var s subscriber
	applyDiscount(&s)   // передаётся указатель, а не структура
	fmt.Println(s.rate) // Обратите внимание, что при обращении к значению поля rate не ставится оператор * - это работает для структур.
}