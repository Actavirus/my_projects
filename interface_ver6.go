// тема: Разбираемся с интерфейсами в Go.
// Что такое пустой интерфейс?
package main

import "fmt"

type forMap interface{}

func main() {
	person := make(map[string]interface{}, 0) // пример объявления карты. Внимание! второй параметр необязателен, т.е. можно записать вот так: person := make(map[string]interface{})
	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64
	fmt.Printf("%v\n", person)              // выводит: map[age:21 height:167.64 name:Alice]
	age := person["age"].(int)              // Внимание! Значение, хранящееся в map, принимает тип interface{} и теряет свой исходный, базовый тип int. И поскольку значение больше не целочисленное, мы не можем прибавить к нему 1. Чтобы увеличить или уменьшить значение "age", необходимо сделать значение снова целочисленным, и ...
	person["age"] = age + 1                 // ... только потом совершать с ним операции.
	fmt.Printf("%v\n", person)              // выводит: map[age:22 height:167.64 name:Alice]
	person["age"] = person["age"].(int) + 1 // здесь просто записали выражение в одну строку не добавляя новую переменную (т.е., для экономии пространства)
	fmt.Printf("%v\n", person)              // выводит: map[age:23 height:167.64 name:Alice]
	// попробуем сделать цикл увеличения значения "age"
	for i := 0; i < 10; i++ {
		person["age"] = person["age"].(int) + 1
		fmt.Printf("%v\n", person) // интересный факт: значение age сохраняется из области вне фикла for, т.е. счёт начинается с "24"
	}
	// поробуем провести эту же процедуру с другими значениями карты
	person["height"] = person["height"].(float64) + 1.00
	fmt.Printf("%v\n", person) // выводит: map[age:33 height:168.64 name:Alice]
	person["name"] = person["name"].(string) + " - is your name."
	fmt.Printf("%v\n", person) // выводит: map[age:33 height:168.64 name:Alice - is your name.]
	// попробуем добавить ещё один параметр внутрь нашей карты
	person["town"] = "New-York City"
	fmt.Printf("%v\n", person) // выводит: map[age:33 height:168.64 name:Alice - is your name. town:New-York City]
	// попробуем определить новую карту с типизирванными значениями - не удалось указать конкретные типы.
	human := make(map[string]forMap)
	human["Name"] = "Alex"
	human["Age"] = 23
	fmt.Println(human) // выводит: map[Age:23 Name:Alex]

}
