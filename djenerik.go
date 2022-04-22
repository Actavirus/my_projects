// https://uproger.com/dzheneriki-v-go-s-primerami-koda-generics-in-golang/
// Дженерики в Go с примерами кода (Generics in Golang)
package main
import (
	"fmt"
)
func main(){
	value := []string{"Arthur", "Renat", "Venera", "Timur"}
	PrintgString(value)
	value = []int{5, 8, 9, 1, 4}
	PrintInt(value)

}
func PrintgString(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func PrintInt(i []int) {
	for _, v := range i {
		fmt.Println(v)
	}
}

// с помощью дженериков мы можем объявлять наши функции следующим образом:
func PrintStringInt[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}