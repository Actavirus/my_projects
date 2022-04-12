// работа с картами
package main

import (
	"fmt"
	"sort"
)

func main() {
	//	var myMap map[string]float64	// объявление переменной для карты
	//	ranks = make(map[string]int)	// непосредственное создание карты
	ranks := make(map[string]int) // короткое объявление переменной (создание карты и объявление переменной для её хранения)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	fmt.Println(ranks)
	// другая карта, в которой и ключами, и значениями являются строки
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])
	fmt.Println(elements)
	// а это карта с целочисленными ключами и логическими значениями
	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime)
	// если набор ключей и значений, которыми должна инициализироваться карта, известен заранее, то как и в случае с массивами и сегментами, можно воспользоваться литералом карты для её создания.
	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	fmt.Println(myMap)
	// пара предыдущих примеров, записанных в форме литералов карт
	ranks = map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	// многострочный литерал карты
	elements = map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}

	// упражнение из учебника
	myProduct := map[string]float64{
		"Earrings": 79.99,
		"Necklace": 89.99,
		"Shirt":    39.99,
		"Pants":    59.99,
	}
	fmt.Println(myProduct)

	// нулевые значения упрощают работу со значениями в картах даже в том случае, если им еще не было явно присвоено значение
	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters)
	fmt.Println(counters["a"], counters["b"], counters["c"])

	// удаление пар "ключ/значение" функцией delete из двух карт ranks и isPrime
	var ok bool
	ranks = make(map[string]int)
	var rank int
	ranks["bronze"] = 3        // присваиваем значение для ключа bronze
	rank, ok = ranks["bronze"] // ok содержит true, потому что значение присутствует
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")    // удаляем ключ bronze с соответсвующим значением
	rank, ok = ranks["bronze"] // ok содержит false, потому что значение было удалено
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	isPrime = make(map[int]bool)
	var prime bool
	isPrime[5] = true      // присваивает значение для ключа 5
	prime, ok = isPrime[5] // ok содержит true, потому что значение присутствует
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)     // удаляет ключ 5 с соответствующим значением
	prime, ok = isPrime[5] // ok содержит false, потому что значение было удалено
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

	// вывод ключей и значений карты в алфавитном порядке с использованием цикла for...range и функции strings пакета sort
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string // построение сегмента со всеми ключами карты
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)          // сегмент сортируется в алфавитном порядке.
	for _, name := range names { // имена обрабатываются в алфавитном порядке
		fmt.Printf("%s набрал(-а) %0.1f%%\n", name, grades[name]) // здесь name - текущее имя используется для выборки оценки из карты
	}
}
