// тема - утверждения типа. Интерфейс.
package main

import "fmt"

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
func (r Robot) MoveRight() {
	fmt.Println("Turn right")
}
func (r Robot) MoveLeft() {
	fmt.Println("Turn left ")
}

type NoiseMaker interface {
	MakeSound()
	MoveLeft()
}

func main() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler") // здесь NoiseMaker - определяем переменную с типом интерфейса...; Robot - ...и присваиваем значение типа, поддерживающего интерфейс.
	noiseMaker.MakeSound()                            // вызов метода, который является частью интерфейса.
	var robot Robot = noiseMaker.(Robot)              // обратное преобразование к конкретному типу с использованием утверждения типа.
	robot.Walk()                                      // вызов метода, определенного для конкретного типа(не для интерфейса).
	robot.MoveRight()                                 // вызов метода, определенного для конкретного типа(не для интерфейса).
	robot.MoveLeft()                                  // работает так...
	noiseMaker.MoveLeft()                             // ...и вот так.
}
