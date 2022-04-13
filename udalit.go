package main
import (
"fmt"
"sort"
)
func main() {
	hand := []int{5, 8, 6, 4, 3, 11, 82, 6, 5, 7}
	sort.Ints(hand)
	fmt.Println(len(hand))
	fmt.Println(hand)
	count := make(map[int]int)
	fmt.Println(count)
	for _, i := range hand {
		count[i] = count[i] + 1
	}
	fmt.Println(count)

}
