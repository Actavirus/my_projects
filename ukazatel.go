// работа с указателями в функциях
package main

import (
	"fmt"
)

func negate(myBoolean *bool) { // здесь *bool - объявляем, что функция возвращает указатель на bool
	*myBoolean = !*myBoolean // меняем на противоположное значение
}
func main() {
	truth := true
	negate(&truth)     // передаём в функцию negate указатель типа bool
	fmt.Println(truth) // выводим изменённое значение
	lies := false
	negate(&lies)     // передаём в функцию negate указатель типа bool
	fmt.Println(lies) // выводим изменённое значение
}
