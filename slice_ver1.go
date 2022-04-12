// Срез массива в Golang
package main

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice []string
type Planets4 []string

func hyperspace(worlds []string) { // это функция. Данный аргумент worlds является срезом
	for i := range worlds { // пройти до конца среза worlds (т.е., по всем его элементам)
		worlds[i] = strings.TrimSpace(worlds[i]) // TrimSpace - возвращает срез строки с удалением всех начальных и конечных пробелов.
	}
}

// func (p StringSlice) Sort() {

// }

func main() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун", // запятая в самом конце является обязательной
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants) // выводит: [Меркурий Венера Земля Марс] [Юпитер Сатурн] [Уран Нептун]
	fmt.Println(gasGiants[0])                      // выводит: Юпитер

	// можно разрезать массив, а потом полученный срез разрезать ещё раз
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice) // выводит:

	// присваивание нового значения к элементу среза меняет базовый массив и другие срезы
	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Посейдон"
	fmt.Println(planets)                         // выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Посейдон]
	fmt.Println(iceGiants, iceGiantsMarkII, ice) // выводит: [Уран Посейдон] [Уран Посейдон] [Уран Посейдон]

	// индексы для среза по умолчанию
	terrestrial = planets[:4] // аналог terrestrial := planets[0:4] - это первые 4 элемента массива (0, 1, 2 и 3; 4 элемент не учитывается!!!)
	gasGiants = planets[4:6]
	iceGiants = planets[6:]  // аналог iceGiants := planets[6:8]
	allPlanets := planets[:] // срез содержит все восемь планет
	fmt.Println(allPlanets)  // выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Посейдон]

	// срез строк в Golang
	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println(tune)          // выводит: tune
	neptune = "Poseidon"       // Внимание! присваивание нового значени массиву не изменит значение среза, и наоборот
	fmt.Println(tune)          // выводит: tune
	question := "¿Cómo estás?" // Внимание! индексы учитывают количество байтов, а не рун.
	fmt.Println(question[:6])  // выводит: ¿Cóm

	// Композитные литералы для срезов
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"} // инициализация среза
	fmt.Printf("slice %T\n", dwarfs)                                      // выводит: slice []string

	// Преимущества среза массива в Golang
	// задача из учебника: удаление пробелов в начале и в конце каждого элемента среза
	planets2 := []string{"   Венера     ", "Земля   ", "   Марс"}
	hyperspace(planets2)
	fmt.Println(planets2)                    // выводит: [Венера Земля Марс]
	fmt.Println(strings.Join(planets2, "-")) // Join - объединяет элементы своего первого аргумента для создания единой строки. Выводит: Венера-Земля-Марс
	a := strings.Join(planets2, " ")         // т.е., теперь "Венера Земля Марс" - это одна целая строка.
	fmt.Println(a)                           // выводит: Венера Земля Марс

	// Срезы с методами в Golang
	sort.StringSlice(dwarfs).Sort() // сортирует planets в алфавитном порядке
	fmt.Println(dwarfs)             // выводит: [Макемаке Плутон Хаумеа Церера Эрида]
	// fmt.Println(sort.Strings(allPlanets)) // выводит: sort.Strings(allPlanets) used as value. - не могу понять почему это ошибка???

	// задача из учебника: Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты. Используйте программу для изменения названий планет Марс, Уран и Нептун. - моё решение задачи
	planets3 := []string{"Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}
	terraform(planets3)            // использование функции
	fmt.Println()                  // просто пустая строка
	Planets4(planets3).terraform() // преобразование типа и использование метода

	// задача из учебника: Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты. Используйте программу для изменения названий планет Марс, Уран и Нептун. - решение из учебника
	Planets4(planets3[3:4]).terraform1()
	Planets4(planets3[6:]).terraform1()
	fmt.Println(planets3)
}
func terraform(planets3 []string) { // это функция с использованием If else
	for _, planet := range planets3 {
		var new string = "Новый "
		if planet == "Марс" {
			planet = new + planet
		} else if planet == "Уран" {
			planet = new + planet
		} else if planet == "Нептун" {
			planet = new + planet
		}
		fmt.Println(planet)
	}
}
func (p Planets4) terraform() { // это метод с использованием Switch
	for planet := range p {
		var new string = "Новый "
		switch p[planet] {
		case "Марс":
			p[planet] = new + p[planet]
		case "Уран":
			p[planet] = new + p[planet]
		case "Нептун":
			p[planet] = new + p[planet]
		}
		fmt.Println(p[planet])
	}
}
func (planets Planets4) terraform1() {
	for i := range planets {
		planets[i] = "Новый " + planets[i]
	}
}
