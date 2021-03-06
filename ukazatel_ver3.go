// пример работы указателя с структурами
package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct // создание значения структуры
	value.myField = 3
	var pointer *myStruct = &value // получает указатель на значение структуры
	//	fmt.Println(*pointer.myField)	// пытаемся получить значение структуры, на которое ссылается указатель - возникает ошибка!, поэтому...
	fmt.Println((*pointer).myField) // ...ставим (), т.е. сначала получаем значение структуры muStruct, после чего обращаемся к полю структуры.
	fmt.Println(pointer.myField)    // круглые скобки и оператор * не обязательны, их можно не указывать. Результат тот же.
	// этот способ (обращение к полям структуры с помощью оператора .) также работает для присваивания значений полям структур по указателю:
	pointer.myField = 9 // присваивание полю структуры по указателю
	fmt.Println(pointer.myField)
}
