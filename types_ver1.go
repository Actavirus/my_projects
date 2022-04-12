// определяемые - т.е. типы, которые мы сами создаём
package main

import "fmt"

type Population float64
type Liters float64
type Gallons float64
type Milliliters float64

func ToGallons(l Liters) Gallons { // значение типа Gallons чуть больше 1/4 от значения типа Liters
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters { // значение типа Liters почти вчетверо больше значения типа Gallons
	return Liters(g * 3.785)
}

func main() {
	// упражнение из учебника
	var population Population
	population = Population(572.0)
	fmt.Println("Sleepy Greek County population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Greek County population:", population)

	// упражнение из учебника: перевод из галлонов в литры и наоборот перед сложением
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0)) // Liters преобразуется в Gallons перед сложением
	busFuel += ToLiters(Gallons(30.0)) // Gallons преобразуется в Liters перед сложением
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
	fmt.Printf("%.2f литров в 1 галлоне и %.2f галлонов в 1 литре\n", ToLiters(1), ToGallons(1))
}
