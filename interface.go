// работа с типом Интерфейс
package main

import (
	"fmt"
	mypkg "github/headfirstgo/interface"
)

func main() {
	var value mypkg.MyInterface                // объявление переменной с использованием типа Интерфейс
	value = mypkg.MyType(123)                  // значения myType поддерживают myInterface, поэтому это значение может быть присвоено переменной с типом myInterface.
	value.MethodWithoutParameters()            // мы можем вызвать любой метод, входящий в myInterface
	value.MethodWithParameter(127.30)          // мы можем вызвать любой метод, входящий в myInterface
	fmt.Println(value.MethodWithReturnValue()) // мы можем вызвать любой метод, входящий в myInterface
}
