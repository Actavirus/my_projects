//программа по вычислени суммы чисел
package main

import "fmt"

func main() {
	myValue := []int{1, 5, 8, 3, 7}
	sum := 0
	for _, value := range myValue {
		sum += value
	}
	fmt.Println("Сумма чисел равна:", sum)
}
