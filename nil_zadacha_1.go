// задача из учебника: Рыцарь встал на пути Артура. Герой безоружен, он представлен значением nil для leftHand *item.
// Имплементируйте структуру character с методами вроде pickup(i *item) и give(to *character).
// Потом используйте выученное в данном уроке для написания скрипта, в котором Артур достает объект и передает его рыцарю.
// Каждое действие должно отображаться с соответствующим описанием.
package main

import (
	"fmt"
)

type item struct { // "пункт"
	name string
}
type character struct { // "персонаж"
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v поднимает %v \n", c.name, i.name)
	c.leftHand = i
}
func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v ничего не может дать\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v с занятыми руками\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v дает %v %v\n", c.name, to.name, to.leftHand.name)
}
func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v ничего не несет", c.name)
	}
	return fmt.Sprintf("%v несет %v", c.name, c.leftHand.name)
}
func main() {
	arthur := &character{name: "Артур"}
	shrubbery := &item{name: "кустарник"}  // "кустарник"
	arthur.pickup(shrubbery)               // выводит: Артур поднимает кустарник
	knight := &character{name: "Рыцарь/ю"} // "рыцарь"
	arthur.give(knight)                    // выводит: Артур дает Рыцарь/ю кустарник
	fmt.Println(arthur)                    // выводит: Артур ничего не несет
	fmt.Println(knight)                    // выводит: Рыцарь/ю несет кустарник
	fmt.Println("-----------------------------")

	// поэкспериментируем с задачкой:
	arthur2 := &character{name: "A"}
	var shrubbery2 *item // shrubbery2 == nil
	arthur2.pickup(shrubbery2)
	knight2 := &character{name: "B"}
	arthur2.give(knight2) // выводит: A ничего не может дать
	fmt.Println(arthur2)  // выводит: A ничего не несет
	fmt.Println(knight2)  // выводит: B ничего не несет
}
