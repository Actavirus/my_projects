// программа вычисления факториала числа (n!)
package main

import "fmt"

func factorial(x uint) uint { // factorial вызывает саму себя, что делает эту функцию рекурсивной.
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // здесь factorial(x-1) - функция вызывает саму себя
}
func main() {
	fmt.Println(factorial(10))
}
