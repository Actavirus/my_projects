// Пример 1: Расчет факториала с использованием рекурсии в Go
// https://coderlessons.com/tutorials/kompiuternoe-programmirovanie/uchis-go-programmirovanie/go-rekursiia
// В следующем примере вычисляется факториал данного числа с использованием рекурсивной функции 
package zadacha1
import (
	"fmt"
)
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i - 1)
}

func main() {
	var i int
	fmt.Print("Enter the number: ")
	fmt.Scan(&i)
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}
