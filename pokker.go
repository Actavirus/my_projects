// Учебник: создайте карточную игру с Golang (Go)
// https://bestprogrammer.ru/programmirovanie-i-razrabotka/uchebnik-sozdajte-kartochnuyu-igru-s-golang-go
package main
import (
	"fmt"
	"sort"
)
func main(){
	hand := []int{5, 2, 4, 4, 1, 3, 5, 6, 3}
	k := 3
	fmt.Println(isHandOfStraights(hand, k))

	hand2 := []int{1, 9, 3, 5, 7, 4, 2, 9, 11}
	k = 2
	fmt.Println(isHandOfStraights(hand, k))
}

// Шаг 1 : Настройте функцию.
func isHandOfStraights(hand []int, k int) bool {
	// Шаг 2 : Проверьте, делится ли количество карт в руке на k. В противном случае мы не можем создавать группы, поэтому возвращайтесь false.
	if len(hand) % k != 0 {
		return false
	}

	// Шаг 3 : Подсчитайте количество появлений каждой карты в данной руке.
	count := make(map[int]int)
	for _, i := range hand {
		count[i] = count[i] + 1
	}

	// Шаг 4 : отсортируйте список и начните перемещаться по нему с карты с самым низким рейтингом. Мы можем использовать хэш-карту, сохраняя номера карт как ключи и вхождения как значения.
	sort.Ints(hand)
	i := 0
	n := len(hand)

	// Шаг 5 : Используйте вложенный цикл, который выполняется k раз.
	for i < n {
		current := hand[i]
		for j := 0; j < k; j++ {
			// Шаг 5.1 : Проверьте, есть ли на карте текущая карта и следующие k-1 карты (в возрастающем рейтинге) count. Если какой-либо из них не существует, вернитесь false.
			if _, ok := count[current + j]; !ok || count[current + j] == 0{
				return false
			}
			// Шаг 5.2 : Когда каждая из требуемых карточек найдена, уменьшите количество ее появлений в count.
			count[current + j]
		}
		// Шаг 5.3 : После того, как будет найдена полная группа, используйте цикл while, чтобы найти самую маленькую карту следующей группы и определить, у какой из следующих карт countосталось больше нуля вхождений.
		for i < n && count[hand[i]] == 0{
			i++
		}
	}

	//Шаг 6 : Вернитесь, trueесли все карточки отсортированы по группам.
	return true
}
