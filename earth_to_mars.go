// программа-пример для вывода списка билетов на рейс Земля-Марс
// задача из учебника
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const data string = "13 October 2020"
	const distance int = 62100000 // км
	//	var speed int                 // от 16 до 30 км/ч
	spaceLine := []string{"Virgin Galactic", "SpaceX", "Space Adventures"}
	tripType := []string{"Round-trip", "One-way"}

	fmt.Printf("%20s | %20s | %5s | %12s | %6s\n", "Spaceline", "Data of fly", "Days", "Trip type", "Price")

	fmt.Println("===============================================================================")

	for i := 0; i < 10; i++ {
		j := rand.Intn(3)
		k := rand.Intn(2)
		days := rand.Intn(50) + 1
		price := rand.Intn(100) + 1

		if tripType[k] == "Round-trip" {
			price = price * 2
		}
		fmt.Printf("%20s | %20s | %5d | %12s | $%6d\n", spaceLine[j], data, days, tripType[k], price)
	}

}
