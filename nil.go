// Значение nil в Golang (https://golangify.com/nil)
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Вызывает ли nil сбои в Golang?
	// Если указатель никуда не указывает, попытка разыменования указателя не сработает, что показано в Листинге 1.
	// Разыменование указателя nil приведет к сбою программы.
	var nowhere *int
	fmt.Println(nowhere) // выводит: <nil>
	// fmt.Println(*nowhere) // выводит: "nil pointer dereference"
	// Избежать сбоя несложно. Это вопрос защиты от разыменования указателя nil с оператором if, что показано в следующем листинге:
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	// Защита методов в Golang
	// Go вызывает методы даже в том случае, если у приемника значение nil.
	// Приемник nil ведет себя так же, как и параметр nil. Это значит, что методы могут защищать от значений nil, как показано в следующем примере.
	// Вместо проверки на наличие nil перед вызовом метода "a" следующий листинг защищает от приемников nil внутри метода.
	type person struct {
		age int
	}
	a := func(p *person) {
		if p == nil {
			return
		}
		p.age++
	}
	var nobody *person
	fmt.Println(nobody) // выводит: <nil>
	a(nobody)

	// Значения функций nil в Golang
	// Когда переменная объявляется как тип функции, ее значение по умолчанию будет nil.
	// В следующем листинге у fn тип функции, но нет присваивания к какой-то конкретной функции.
	var fn func(a, b int) int
	fmt.Println(fn == nil) // выводит: true
	// Если бы предыдущий код вызывал fn(1, 2), программа привела бы к сбою с разыменованием указателя nil, потому что нет функции, присвоенной fn.
	sortStrings := func(s []string, less func(i, j int) bool) {
		if less == nil {
			less = func(i, j int) bool {
				return s[i] < s[j]
			}
		}
		sort.Slice(s, less) // встроенная функция сортировки
	}
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food) // выводит: [carrot celery onion]
	// Напишите одну строку кода для сортировки food от самой короткой до самой длинной строки.
	// sortStringsTwo := func(food, func (i, j int) bool {return len(food[i]) < len(food[j]) }) - почему-то возникает ошибка

	// Срезы nil в Golang
	// Срез, объявленный без композитного литерала или встроенной функции make, будет со значением nil.
	var soup []string         // это срез из стрингов
	fmt.Println(soup == nil)  // выводит: true
	soup2 := []string{}       // это срез из стрингов
	fmt.Println(soup2 == nil) // выводит: false
	// К счастью, ключевое слово range, а также встроенные функции len и append будут работать со срезами nil, как показано в следующем примере:
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	fmt.Println(len(soup)) // выводит: 0
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // выводит: [onion carrot celery]
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	fmt.Println(len(soup)) // выводит: 3
	// Пустой срез и срез nil не эквивалентны, но зачастую они заменяют друг друга.
	// Следующий пример передает nil для функции, что принимает срез, пропуская шаг создания пустого среза:
	soup = mirepoix(nil)
	fmt.Println(soup) // выводит: [tomato paper crypto]

	// Карты nil в Golang
	// Как и срезы, карты объявленные без композитного литерала или встроенной функции make, обладают значением nil.
	var soup3 map[string]int
	fmt.Println(soup3 == nil) // выводит: true
	soup4 := make(map[string]int)
	fmt.Println(soup4 == nil) // выводит: false
	// Карты можно прочитать даже когда у них значение nil, что показано в следующем листинге, хотя запись в карту nil вызовет сбой (для записи используем слово make):
	measurement, ok := soup3["onion"] // measurement - "измерение"
	if ok {
		fmt.Println(measurement)
	}
	for ingredient, measurement := range soup3 {
		fmt.Println(ingredient, measurement)
	}

	// Интерфейсы nil в Go
	// Когда объявляется переменная типа интерфейса без присваивания, нулевым значением является nil.
	// Следующий листинг показывает, что тип интерфейса и значение являются nil, и переменная считается равной nil:
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // выводит: <nil> <nil> true
	// Оба, тип интерфейса и значение должны быть типа nil для того чтобы переменная была равна nil, что показано в следующем примере:
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil) // выводит: *int <nil> false
	// Специальный символ %#v  нужен для того, чтобы увидеть тип и значение, а также отображения того, что переменная содержит (*int)(nil), а не просто <nil>, как показано выше:
	fmt.Printf("%#v\n", v) // выводит: (*int)(nil)

	// Альтернатива nil в Golang
	// Вместо использования указателя можно объявить небольшую структуру с несколькими методами.
	// Это требует большего количества кода, но не запрашивает указатель или nil, как показано в следующем листинге.
	n := newNumber(42)
	fmt.Println(n) // выводит: 42
	e := number{}
	fmt.Println(e) // выводит: not set
}

type number struct {
	value int  // "стоимость"
	valid bool // "действительный"
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "tomato", "paper", "crypto")
}
func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set" // "не установлен"
	}
	return fmt.Sprintf("%d", n.value)
}
