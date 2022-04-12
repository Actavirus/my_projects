// тема - Структуры в Golang — Экспорт структур в JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnError(err error) { // exitOnError выводит любые ошибки и выходит из программы.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Объявление структуры в Golang
	var curiosity struct {
		lat  float64
		long float64
	}

	// Присваивание значений полям структуры
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity.lat, curiosity.long) // выводит: -4.5895 137.4417
	fmt.Println(curiosity)                     // выводит: {-4.5895 137.4417}

	// Повторное использование структур с типами в Go
	type location struct {
		lat, long float64
	}
	var spirit location // Повторное использование типа location
	spirit.lat = -14.5684
	spirit.long = 175.472636
	var opportunity location // Повторное использование типа location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734
	fmt.Println(spirit, opportunity) // выводит: {-14.5684 175.472636} {-1.9462 354.4734}

	// Инициализация структур через композитные литералы Golang
	opportunity = location{lat: -10.5897, long: 58.2157}
	fmt.Println(opportunity) // выводит: {-10.5897 58.2157}
	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)                // выводит: {4.5 135.9}
	spirit = location{-105.548, 54.725} // Композитный литерал в Листинге 4 не уточняет названия полей. Вместо этого значение для каждого поля должно предоставляться в том же порядке, в котором они указаны в определении структуры. Данная форма лучше всего подходит для стабильных типов, у которых только несколько полей.
	fmt.Println(spirit)                 // выводит: {-105.548 54.725}
	curiosity = location{-64.21456, 48.2365}
	fmt.Printf("%v\n", curiosity)  // выводит: {-64.21456 48.2365}
	fmt.Printf("%+v\n", curiosity) // выводит: {lat:-64.21456 long:48.2365} - т.е., знак "+" нужен для вывода названий полей.

	// Копирование структур
	bradbury := location{-85.2135, 234.58}
	curiosity = bradbury                         // Переменная curiosity инициализируется с копией значений, находящихся в bradbury, поэтому изменения происходят независимо.
	curiosity.long += 0.016                      // Направляется на восток к бухте Йеллоунайф (текст взят из учебника)
	fmt.Printf("%+v %+v\n", bradbury, curiosity) // выводит: {lat:-85.2135 long:234.58} {lat:-85.2135 long:234.596}

	// Пример среза структур в Golang
	type location2 struct {
		name string
		lat  float64
		long float64
	}
	locations := []location2{
		{name: "Bradbury Landing", lat: -98.215, long: 37.325},
		{name: "Columbia Memorial Station", lat: -47.1256, long: 91.276},
		{name: "Challenger Memorial Station", lat: -7.8245, long: 208.4698},
	}
	fmt.Printf("%+v\n", locations) // выводит: [{name:Bradbury Landing lat:-98.215 long:37.325} {name:Columbia Memorial Station lat:-47.1256 long:91.276} {name:Challenger Memorial Station lat:-7.8245 long:208.4698}]

	// Кодирование структур в JSON (JavaScript Object Notation)
	type location3 struct {
		Lat, Long float64 // Поля должны начинаться с большой буквы
	}
	curiosity2 := location3{-68.246, 137.2564}
	bytes, err := json.Marshal(curiosity2) // Для работы, пакет json требует экспорта полей.
	exitOnError(err)
	fmt.Println(string(bytes)) // выводит: {"Lat":-68.246,"Long":137.2564}

	// Изменение JSON через теги struct в Golang
	type location4 struct {
		Lat  float64 `json:"latitude"`  // Теги в struct меняют вывод (или можно указать следующим обраом: "json:\"latitude\"")
		Long float64 `json:"longitude"` // Теги в struct меняют вывод (или можно указать следующим обраом: "json:\"longitude\"")
	}
	curiosity3 := location4{-97.5943, 83.2975}
	bytes, err = json.Marshal(curiosity3)
	exitOnError(err)
	fmt.Println(string(bytes)) // выводит: {"latitude":-97.5943,"longitude":83.2975}

	// задача из учебника: Напишите программу для отображения закодированных данных в JSON трех мест высадки марсохода в Листинге 8. JSON должен содержать название каждого места высадки и использовать теги структуры, как показано в Листинге 10. Чтобы сделать вывод приятнее, используйте функцию MarshalIndent из пакета json. - мой вариант решения.
	type location5 struct {
		Name string  `json:"Name of landing site is"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	locations2 := []location5{
		{Name: "Bradbury Landing", Lat: -98.215, Long: 37.325},
		{Name: "Columbia Memorial Station", Lat: -47.1256, Long: 91.276},
		{Name: "Challenger Memorial Station", Lat: -7.8245, Long: 208.4698},
	}
	bytes2, _ := json.MarshalIndent(locations2, "-", "-")
	fmt.Println(string(bytes2))

	// задача из учебника: Напишите программу для отображения закодированных данных в JSON трех мест высадки марсохода в Листинге 8. JSON должен содержать название каждого места высадки и использовать теги структуры, как показано в Листинге 10. Чтобы сделать вывод приятнее, используйте функцию MarshalIndent из пакета json. - вариант решения из учебника.
	type location6 struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	locations3 := []location6{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	bytes, err = json.MarshalIndent(locations3, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
