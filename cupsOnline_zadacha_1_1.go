// Задача 1 со второго тура (https://cups.online/)
// Сложность: 1 из 5
// Условие задачи:
// Никита собирается открывать свой пункт выдачи заказов. 
// Для обеспечения безопасности посетителей Никита решил обеспечить ПВЗ одноразовыми медицинскими масками, 
// которые будут поставляться каждый месяц в количестве n штук. 
// Местные поставщики продают маски по цене 0.55 рублей за одну штуку, но можно сэкономить, 
// купив упаковку из 24 масок за 11 рублей или коробку из 16 упаковок за 160 рублей. 
// Помогите Никите, определив, сколько коробок, упаковок и отдельных масок он должен купить, 
// чтобы заплатить как можно меньше денег. 
// Примечание: Никита готов купить больше масок, чем ему нужно, если это позволит сэкономить.
// Формат входных данных:
// В одной строке вводится одно целое число N(1<= N <= 50000)
// Формат выходных данных:
// Требуется вывести три числа в одну строку через пробел - количество 
// отдельных масок, упаковок и коробок, которые надо купить.
package main
import (
	"fmt"
	"os"
	"math"
	"errors"
)
func main() {
	var value int

	fmt.Print("Введите количество посетителей: ")
	fmt.Fscan(os.Stdin, &value)

	err := proverka(value)
	if err != nil {
		fmt.Println(err)
	} else {

	v := float64(value)
	onePiece := 0.55
	onePackage := 11.00
	oneBox := 160.00

	pricePiece := v * onePiece

	pricePackage := math.Ceil(v / 24.0) * onePackage

	priceBox := math.Ceil(v / 384.0) * oneBox

	pricePiecePackage := math.Floor(v / 24.0) * onePackage + (v - 24.0 * (math.Floor(v / 24.0))) * onePiece

	pricePieceBox := math.Floor(v / 384.0) * oneBox + (math.Ceil((v - 384.0 * (math.Floor(v / 384.0))) / 24.0)) * onePackage

	pricePiecePackageBox := math.Floor(v / 384.0) * oneBox + math.Floor((v - math.Floor(v / 384) * 384) / 24) * 11 + (v - (math.Floor(v / 384.0) * 384 + math.Floor((v - math.Floor(v / 384) * 384) / 24) * 24)) * onePiece

	minim := min(pricePiece, pricePackage, priceBox, pricePiecePackage, pricePieceBox, pricePiecePackageBox)

	switch minim {
	case pricePiece:
		fmt.Printf("%v %v %v\n", v, 0, 0)
	case pricePackage:
		fmt.Printf("%v %v %v\n", 0, math.Ceil(v / 24.0), 0)
	case priceBox:
		fmt.Printf("%v %v %v\n", 0, 0, math.Ceil(v / 384.0))
	case pricePiecePackage:
		fmt.Printf("%v %v %v\n", (v - 24.0 * (math.Floor(v / 24.0))), math.Floor(v / 24.0), 0)
	case pricePieceBox:
		fmt.Printf("%v %v %v\n", 0, math.Ceil((v - 384.0 * (math.Floor(v / 384.0))) / 24.0), math.Floor(v / 384.0))
	case pricePiecePackageBox:
		fmt.Printf("%v %v %v\n", (v - (math.Floor(v / 384.0) * 384 + math.Floor((v - math.Floor(v / 384) * 384) / 24) * 24)), math.Floor((v - 384.0 * (math.Floor(v / 384.0))) / 24.0), math.Floor(v / 384.0))
	}
	}
}
func min(prices ...float64) float64 { 
	min := 1000.0
	for _, price := range prices { 
		if price < min {
			min = price 
		}
	}
	return min
}
func proverka(value int) error {
	
	if value < 1 || value > 50000 {
		return errors.New("Введите целое число в диапазоне от 1 до 50000")
	}
	return nil
}