// Слайсы и словари в Go (видеокурс OTUS)
package main

import (
	"fmt"
	"sort"
)

func main() {
	// сортировка
	i := []int{3, 2, 1}
	sort.Ints(i)
	fmt.Println(i)
	s := []string{"hello", "cruel", "world"}
	sort.Strings(s)
	fmt.Println(s)
	type User struct {
		Name string
		Age  int
	}
	u := []User{{"vasya", 19}, {"petya", 18}, {"arthur", 32}, {"timur", 15}}
	sort.Slice(u, func(i, j int) bool {
		return u[i].Age < u[j].Age
	})
	fmt.Println(u)

	fmt.Println("----------------------------------------------")
	oldSlice := [][]int{{1, 2, 3}, {4, 5}, {6, 7}}
	fmt.Println(oldSlice)
}
func Concat(slices [][]int) []int {
	var newSlice []int
	for _, s := range slices {
		newSlice = append(newSlice, s...)
	}
	return newSlice
}
