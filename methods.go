// определение методов - т.е. методы, которые мы сами создаём. Создание и работа с методами.
package main

import "fmt"

type MyType string // определяется новый тип, основанный на типе string
type Number int
type Liters float64
type Milliliters float64
type Gallons float64

func (m MyType) sayHi() { // здесь m - определяется параметр получателя; MyType - для MyType определяется метод. Внимание! метод не экспортируется, т.к. имя метода начинается с буквы нижнего регистра.
	fmt.Println("Hi from", m) // здесь m - вывод значения параметра получателя.
}
func (m2 MyType) MethodWithParameters(number int, flag bool) { // здесь m2 - параметр получателя; number - параметр; flag - параметр. Внимание! метод экспортируется, т.к. имя метода начинается с буквы верхнего регистра.
	fmt.Println(m2)
	fmt.Println(number)
	fmt.Println(flag)
}
func (m3 MyType) WithReturn() int { // здесь m3 - возвращаемое значение. Внимание! метод экспортируется, т.к. имя метода начинается с буквы верхнего регистра.
	return len(m3) // здесь len(m3) - возвращает длину основополагающего строкового значения получателя.
}
func (n Number) Add(otherNumber int) {
	fmt.Println(n, " plus", otherNumber, "is", int(n)+otherNumber)
}
func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number) Double() { // здесь *Number - параметр получателя преобразуется в тип указателя.
	*n *= 2 //  здесь *n - обновляется значение по указателю.
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}
func (m Milliliters) ToLiters() Liters {
	return Liters(m / 1000)
}
func main() {
	value := MyType("a MyType value")       // создаётся значение MyType
	value.sayHi()                           // вызывается sayHi для этого значения.
	anotherVulue := MyType("another Value") // создаётся другое значение MyType
	anotherVulue.sayHi()                    // вызывается sayHi для нового значения.

	// определение дополнительных параметров в круглых скобках после имени метода
	value = MyType("MyType value")
	value.MethodWithParameters(4, true) // здесь value - получатель; 4 - аргумент; true - аргумент

	// можно объявить одно или несколько возвращаемых значений, которые будут возвращаться при вызове метода
	value = MyType("MyType value")
	fmt.Println(value.WithReturn()) // выводит возвращаемое значение метода.

	// упражнение из учебника
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	// указатели и параметры получателей
	number := Number(4) // создаём значение Number
	fmt.Println("Original value of number:", number)
	number.Double() // удваиваем Number. Внимание!, изменять вызов метода НЕ НУЖНО!
	fmt.Println("number after calling Double:", number)

	// упражнение из учебника
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	ml := Milliliters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", ml, ml.ToLiters())
}
