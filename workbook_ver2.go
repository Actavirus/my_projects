package main

import "fmt"

func square(x1 *float64) {
	*x1 = *x1 * *x1
}
func swap(x2, y2 *int) {
	c2 := 0
	*&c2 = *x2
	*x2 = *y2
	*y2 = *&c2
}
func main() {
	// упражнение из учебика
	x1 := 1.5
	fmt.Println(x1)
	square(&x1)
	fmt.Println(x1)

	// упражнение из учебника
	x2 := 1
	y2 := 2
	fmt.Println(x2, y2)
	swap(&x2, &y2)
	fmt.Println(x2, y2)
}
