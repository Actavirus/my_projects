// https://golangify.com/pointers
// Напишите программу с черепахой, которая может двигаться вверх, вниз, налево или направо.
// Черепаха должна сохранить координаты (x, y) места, где положительные значения для передвижения вниз и направо.
// Используйте методы для увеличения /уменьшения соответствующей переменной.
// Функция main должна использовать методы, которые вы написали, и выводить окончательное местоположение.
package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}
func (t *turtle) down() {
	t.y++
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}
func main() {
	var t turtle
	t.up()
	t.up()
	t.left()
	t.left()
	fmt.Println(t) // выводит: {-2 -2}
}
