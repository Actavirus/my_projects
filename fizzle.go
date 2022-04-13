// Учебник: создайте карточную игру с Golang (Go)
// https://bestprogrammer.ru/programmirovanie-i-razrabotka/uchebnik-sozdajte-kartochnuyu-igru-s-golang-go
// Найдите максимальную сумму с любого конца массива
package main
import (
	"fmt"
	"math"
)
func main()  {
	deck := []int{5, 3, 4, 4, 2, 3, 2, 6, 3}
	k := 4
	fmt.Println(maxPoints(deck, k))
}

// Шаг 1 : Настройте функцию.
func maxPoints(deck []int, k int) int {
	left := 0
	right := len(deck) - k
	var total, best int
	total = 0

	// Шаг 2 : Предположим, что k карточек справа дают нам максимальное количество очков.
	for i := right; i < len(deck); i++ {
		total += deck[i]
	}
	best = total

	// Шаг 3. Используйте цикл, который выполняется k раз, и проверьте все комбинации.
	for i := 0; i < k; i++{
		// Шаг 4 : Удалите точки карты с правой стороны и добавьте точки с левой стороны.
		total += deck[left] - deck[right]

		// Шаг 5 : Сравните totalточки с текущими bestточками и сохраните максимум.
		best = int(math.Max(float64(best), float64(total)))
		left++
		right++
	}
	return best
}