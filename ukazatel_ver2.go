// пример программы с указателями
package main

import "fmt"

func main() {
	var value int = 2         // создаёт значение типа int
	var pointer *int = &value // получает указатель на это значение
	fmt.Println(pointer)      // Сюрприз! Выводится указатель, а не значение!
	fmt.Println(*pointer)     // с помощью оператора * выводит значение, на которое ссылается указатель
}
