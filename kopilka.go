// программа-копилка, версия 1
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	moneta := [3]float64{0.05, 0.10, 0.25} // три монеты номиналом 5, 10 и 25
	sum := 20.0                            // сумма, которую необходимо достичь
	fmt.Printf("%10s | %10s | %10s\n", "Now", "Vklad", "V kopilke")
	fmt.Println("=============================================")
	for now := 0.0; now < sum; {
		vklad := rand.Intn(3)
		kopilka := now + moneta[vklad]
		fmt.Printf("%10.2f | %10.2f | $%10.2f\n", now, moneta[vklad], kopilka)
		now = +kopilka
	}
}
