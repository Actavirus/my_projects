// тема: Указатели в Golang (https://golangify.com/pointers)
package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}
type stats struct {
	level             int
	endurance, health int
}

func main() {
	// Амперсанд (&) и звездочка астериск (*) в Golang
	answer := "42"
	fmt.Println(&answer) // выводит: "0xc00000a088". Здесь: &answer - это место в памяти, где компьютер хранит 42
	// & - "амперсанд" - оператор адреса, который предоставляет адрес памяти значения. Обратная операция называется "разыменованием".
	address := &answer
	fmt.Println(*address) // выводит: "42"

	// Типы указателей в Golang
	// Указатели хранят адреса памяти.
	fmt.Printf("address это %T\n", address) // выводит: "address это *string". Переменная address является указателем типа *string, об этом сообщает специальный символ %Т
	canada := "Canada"
	var home *string                   // Звездочка перед типом обозначает тип указателя, ...
	fmt.Printf("home is a %T\n", home) // выводит: "home is a *string"
	home = &canada
	fmt.Println(*home) // выводит: "Canada" ... а звездочка перед названием переменной нужна для указания на значение, к которому отсылается указатель.

	// Указатели для указания Golang
	var administrator *string
	scolese := "Christopher J.Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // выводит: Christopher J.Scolese
	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator) // выводит: Charles F. Bolden
	// Изменить значение bolden можно в одном месте, потому что переменная administrator указывает на bolden вместо хранения копии:
	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator)         // выводит: Charles Frank Bolden Jr.
	fmt.Println(&bolden, administrator) // выводит: 0xc000084280 0xc000084280, т.е. они оба указывают на один и тот же адрес.
	// Можно разыменовать administrator для непрямого изменения значения bolden:
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden) // выводит: Maj. Gen. Charles Frank Bolden Jr.
	// Результатом присваивания major к administrator является новый указатель, что также указывает на строку bolden:
	major := administrator
	*major = "Major General Charles rank Bolden Jr."
	fmt.Println(bolden) // выводит: Major General Charles rank Bolden Jr.
	// Указатели major и administrator оба содержат один и тот же адрес памяти, следовательно, они равны:
	fmt.Println(administrator == major) // выводит: true
	// На смену Чарльзу Болдену пришел Роберт М. Лайтфут. После данного изменения administrator и major перестали указывать на одинаковый адрес памяти.
	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major) // выводит: false, т.к. теперь administrator указывает на другой адрес.
	// Присваивание разыменованного значения major к другой переменной создает копию строки.
	// После создания клона прямые и непрямые изменения с bolden не будут иметь эффект над значением charles и наоборот:
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles) // выводит: Major General Charles rank Bolden Jr.
	fmt.Println(bolden)  // выводит: Charles Bolden
	// Если две переменные содержат одинаковую строку, они считаются равными, как в случае с charles и bolden в следующем коде.
	// Даже несмотря на то, что их адреса памяти отличаются:
	charles = "Charles Bolden"
	fmt.Println(charles == bolden)   // выводит: true
	fmt.Println(&charles == &bolden) // выводит: false

	// Указатели и структуры Go
	// В отличие от строк и чисел, перед композитными литералами ставится префикс в виде оператора адреса.
	// В следующем примере переменная timmy содержит адрес памяти, указывающий на структуру person.
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	// Кроме того, нет необходимости разыменования структур для получения доступа к их полю.
	// Следующий листинг предпочтительнее, чем эта запись: (*timmy).superpower, но по факту это одно и то же:
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy) // выводит: &{name:Timothy superpower:flying age:10}

	// Указатели и массивы в Go
	// Массивы также предоставляют автоматическое разыменование, как показано в следующем примере, т.е нет необходимости писать более громоздкий (*superpowers)[0].
	superpowers := &[3]string{"flight", "invisibility", "super strenght"}
	fmt.Println(superpowers[0])   // выводит: flight
	fmt.Println(superpowers[1:2]) // выводит: [invisibility]
	// Композитным литералам для срезов и карт также можно добавить префиксы с оператором адреса (&), однако тогда не будет автоматического разыменования:
	mySuperpowers := &[]string{"read", "listen", "eat", "write"}
	fmt.Println((*mySuperpowers)[1]) // выводит: listen

	// Указатели в качестве параметров Golang
	// Указатели используются для возможности использования редактирования через границы методов и функций.
	// Параметры функции и метода передаются через значение. Это значит, что функции всегда оперируют копией переданных аргументов.
	// Когда указатель передается функции, функция получает копию адреса памяти.
	// Через разыменование адреса памяти функция может изменить значение, на которое указывает указатель.
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca) // выводит: {name:Rebecca superpower:imagination age:15}
	// Какой код вернет Timothy 11? Отталкивайтесь от Листинга 6:
	birthday(timmy)        // т.к. переменная timmy уже является указателем
	fmt.Println(timmy.age) // выводит: 11

	// Приемники указателя в Golang
	// Метод birthday в следующем примере использует указатель для приемника, что позволяет методу изменять атрибуты персоны (смотри за границей main).
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry) // выводит: &{name:Terry superpower: age:16}
	// Альтернативно, вызов метода в следующем листинге не использует указатель, однако по-прежнему работает.
	// Go автоматически определяет адрес переменной (&) при вызове метода через пояснение с точкой, поэтому вам не нужно писать (&nathan).birthday().
	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan) // выводит: {name:Nathan superpower: age:18}
	// к вопросу о том, для чего нужны указатели:
	nathan.birthday2()
	fmt.Printf("%+v\n", nathan) // выводит: {name:Nathan superpower: age:18} - т.е. значение age не изменилось, т.к. в функции birthday2 была изменена только копия age.

	// Внутренние указатели Golang
	// В Go есть удобные внутренние указатели, которые нужны для определения адреса памяти поля внутри структуры.
	// Функция levelUp в следующем примере изменяет структуру stats, и поэтому нуждается в указателе (смотри за границей main)
	// Оператор адреса в Go может использоваться для указания на поле внутри структуры, как показано в следующем примере:
	type character struct {
		name  string
		stats stats
	}
	player := character{name: "Matthias"}
	levelUp(&player.stats)            // Код &player.stats предоставляет указатель на внутреннюю часть структуры.
	fmt.Printf("%+v\n", player.stats) // выводит: {level:1 endurance:56 health:280}

	// Изменение, или мутации массивов в Golang
	// Хотя предпочтительнее использовать срезы, а не массивы, в случаях, когда нет нужды менять их длину, можно задействовать и массивы.
	// В качестве примера можно рассмотреть шахматную доску. В следующем примере показано, как указатели позволяют функциям изменять элементы массива.
	// В данном примере функция находится внутри другой функции (main), т.е. это наглядное "замыкание".
	a := func(board *[8][8]rune) {
		board[0][0] = 'r'
		board[7][7] = 'a'
	}
	var board [8][8]rune
	a(&board)
	fmt.Printf("%c\n", board) // выводит: r

	// Карты в роли указателей в Go
	// во время присваивания или передачи в качестве аргументов карты не копируются.
	// Карты являются скрытыми указателями, поэтому указание на карту является лишним. Не делайте этого: func demolish(planets *map[string]string) - лишний указатель

	// Срезы в Go — указатели на массив
	// Явный указатель на срез полезен только в том случае, когда нужно модифицировать сам срез: длину, вместимость или начальный набор.
	// В следующем примере функция b редактирует длину среза planets. Вызов функции (main) не увидит изменения, если b не использует указатель:
	b := func(planets *[]string) {
		*planets = (*planets)[0:8]
	}
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
	b(&planets)
	fmt.Println(planets) // выводит: [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune], т.е. функция b уменьшила срез.
	// Вместо изменения переданного среза, как в примере выше, лучше написать функцию b для возвращения нового среза.

}
func birthday(p *person) {
	p.age++ // этот синтаксис предпочтительнее чем: (*p).age
}
func (p *person) birthday() {
	p.age++
}
func (p person) birthday2() {
	p.age++
}
func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}
