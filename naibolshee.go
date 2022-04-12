// программа находит наибольшее число в списке
package main

import "fmt"

func search(a ...int) int {
	y := 0
	for _, x := range a {
		if x > y {
			y = x
		}
	}
	return y
}
func main() {
	myValue := []int{25, 56, 87, 92, 14, 55, 0, 101, 102}
	fmt.Println("Наибольшее число из списка:", search(myValue...))
}
