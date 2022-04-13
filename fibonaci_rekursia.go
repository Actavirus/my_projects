// Пример 2: ряд Фибоначчи с использованием рекурсии в го
// https://coderlessons.com/tutorials/kompiuternoe-programmirovanie/uchis-go-programmirovanie/go-rekursiia
// В следующем примере показано, как сгенерировать ряд Фибоначчи для заданного числа с помощью рекурсивной функции
package main
import "fmt"
func main()  {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d \n", fibonaci(i))
	}
}
func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i - 1) + fibonaci(i - 2)
}