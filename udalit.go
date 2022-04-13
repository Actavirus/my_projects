package main
import (
"fmt"
"sort"
)
func main() {
	hand := []int{5, 8, 6, 4, 3, 11, 82, 6, 5, 7}
	sort.Ints(hand)
	fmt.Println(len(hand))	// 10
	fmt.Println(hand)	// [3 4 5 5 6 6 7 8 11 82]
	count := make(map[int]int)
	fmt.Println(count)	// map[]
	for _, i := range hand {
		count[i] = count[i] + 1
	}
	fmt.Println(count)	// map[3:1 4:1 5:2 6:2 7:1 8:1 11:1 82:1]

	// k := 3
	// for j := 0; j < k; j++ {
	// 	if _, ok := count[current + j]; !ok || count[current + j] == 0{
	// 		return false
	// 	}
	// 	count[current + j] - 1
	// }
	current := hand[5]
	fmt.Println(current)	// 6
	a := count[current + 1]
	fmt.Println(a)	// 1
	fmt.Println(count[current])	// 2
	count[current] --
	fmt.Println(count[current])	// 1

}
