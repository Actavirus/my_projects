// программа генерирует последовательность четных чисел
package main

import "fmt"

func makeEvenGenerator() func() uint { // возвращает функцию, которая генерирует чётные числа
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2 // при каждом вызове функции к переменной i прибавляется 2. Но!, в отличие от локальных переменных её значение сохраняется между вызовами.
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	value := 10 // количество проходов для цикла For.
	for i := 0; i < value; i++ {
		fmt.Println(nextEven())
	}
}
