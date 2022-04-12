// тема: Карта — ассоциативный массив в Golang.
// Карты являются универсальными коллекциями для неструктурированны данных.
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

func main() {
	temperature := map[string]int{"Земля": 15, "Марс": -65}
	earth := temperature["Земля"]
	fmt.Printf("Средняя температура на поверхности Земли составляет %v градусов.\n", earth) // выводит: Средняя температура на поверхности Земли составляет 15 градусов.
	temperature["Земля"] = 16                                                               // небольшое изменение климата
	temperature["Венера"] = 464                                                             // добавление дополнительного значения в карту
	fmt.Println(temperature)                                                                // выводит: map[Венера:464 Земля:16 Марс:-65]
	moon := temperature["Луна"]                                                             // присвоение ключа которого нет в карте
	fmt.Println(moon)                                                                       // выводит: 0
	if moon, ok := temperature["Луна"]; ok {                                                // синтаксис comma, ok - для обозначения разницы между ключом "Луна", которого нет в карте, и тем, что находится в карте с температурой 0.
		fmt.Printf("Средняя температура на поверхности Луны составляет %v градусов.\n", moon) // при наличии ключа значение дополнительной переменной ok будет равно true
	} else { // в противном случае - false
		fmt.Println("Где Луна?")
	}

	// Копируются ли карты в Golang?
	planets := map[string]string{"Земля": "Сектор ZZ9", "Марс": "Сектор ZZ9"}
	planetsMarkII := planets
	fmt.Println(planets) // выводит: Копируются ли карты в Golang?
	planets["Земля"] = "упс"
	fmt.Println(planets)       // выводит: map[Земля:упс Марс:Сектор ZZ9]
	fmt.Println(planetsMarkII) // выводит: map[Земля:упс Марс:Сектор ZZ9]
	delete(planets, "Земля")   // ключ "Земля" удаляется из карты (удаляем элемент из карты)
	fmt.Println(planetsMarkII) // выводит: map[Марс:Сектор ZZ9]	- т.е., изменения в одной карте, затрагивают и другую.

	// Предварительное обозначение карты через make
	temperature_2 := make(map[float64]int, 8) // для карт make принимает один или два параметра. Здесь 8 - это место для количества ключей.
	fmt.Println(temperature_2)                // выводит: map[]

	// Использование карты для подсчета частоты использования элементов
	temperatures := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}
	frequency := make(map[float64]int) // Внимание! в картах два ключа не могут быть одинаковыми.
	for _, t := range temperatures {   // итерирует через срез (индекс, значение)
		frequency[t]++
	}
	fmt.Println(frequency)
	for t, num := range frequency { // итерирует через карту (ключ, значение)
		fmt.Printf("%+.2f встречается %d раз(а)\n", t, num) // Внимание! Go не гарантирует порядок ключей карты, поэтому выводы при различных запусках могут отличаться.
	}

	// Группирование данных с картами и срезами Golang (Вместо определения частоты упоминания температур сгруппируем их вместе с разделением каждой в 10°. Для этого в следующем пример карты группируются в срез температур данной группы.)
	groups := make(map[float64][]float64) // карта с ключами float64 и значениями []float64 (т.е., это срез)
	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // округляет температуры вниз -20, -30, и так далее
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures) // выводит: группированные данные
	}

	// Множества в Golang
	set := make(map[float64]bool) // создание карты с булевыми значениями
	for _, t := range temperatures {
		set[t] = true
	}
	if set[-28.0] { // проверяет, является ли значение -28.0 частью множества set. Внимание! данная запись "if set[-28.0]" - прдполагает условие, что проверяет истинность утверждения, т.е. налогично "if set[-28.0] == true"
		fmt.Println("часть множества") // выводит: "часть множества"
	}
	fmt.Println(set)                       // выводит: map[-33:true -31:true -29:true -28:true -23:true 32:true] - Видно, что карта содержит по одному ключу для каждой температуры, дубликаты удаляются.
	unique := make([]float64, 0, len(set)) // конвертируем температуры обратно в срез для их сортировки.
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique) // используем встроенную функцию sort для сортировки значений.
	fmt.Println(unique)   // выводит: [-33 -31 -29 -28 -23 32]

	// задача из учебника: Напишите функцию для подсчета частоты упоминания слов в строке текста и возвращения карты со словами и числом, указывающем, сколько раз они употребляются. Функция должна конвертировать текст в нижний регистр и обрезать знаки препинания. Используйте пакет strings. Функции, которые пригодятся для выполнения данного задания: Fields, ToLower и Trim. Используйте функцию для подсчета частоты слов следующего отрывка текста и последующего вывода числа употребления каждого слова, что встречается более одного раза. - мой вариант решения (использую 3 функции).
	line := "..   ..BeFore:     After,!   :tomOrrow!!   !,"
	fmt.Printf("%q\n", strings.ToLower(line))     // выводит: "..   ..before:     after,!   :tomorrow!!   !,"
	fmt.Printf("%q\n", strings.Fields(line))      // выводит: [".." "..BeFore:" "After,!" ":tomOrrow!!" "!,"]
	fmt.Printf("%q\n", strings.Trim(line, "!.,")) // выводит: "   ..BeFore:     After,!   :tomOrrow!!   "
	line2 := strings.Fields(line)
	for i := range line2 {
		line2[i] = strings.ToLower(strings.Trim(line2[i], "!,.:"))
	}
	fmt.Printf("%q\n", line2)     // выводит: ["" "before" "after" "tomorrow" ""]
	line3 := make(map[string]int) // объявляем карту с ключами типа string и значением типа int
	for _, i := range line2 {
		line3[i]++
	}
	fmt.Printf("%q\n", line3) // выводит: map["":'\x02' "after":'\x01' "before":'\x01' "tomorrow":'\x01']
	fmt.Println(line3)        // выводит: map[:2 after:1 before:1 tomorrow:1]
	text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."
	output(sliceToMap(textToSlice(text)))

	// задача из учебника: Напишите функцию для подсчета частоты упоминания слов в строке текста и возвращения карты со словами и числом, указывающем, сколько раз они употребляются. Функция должна конвертировать текст в нижний регистр и обрезать знаки препинания. Используйте пакет strings. Функции, которые пригодятся для выполнения данного задания: Fields, ToLower и Trim. Используйте функцию для подсчета частоты слов следующего отрывка текста и последующего вывода числа употребления каждого слова, что встречается более одного раза. - вариант решения из учебника (использует 1 функцию).
	frequency2 := countWords(text)
	for word, count := range frequency2 {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}
func textToSlice(text string) []string { // функция перевода текста в срез, удаления из значений (слов) знаков препинания, перевода в нижний регистр.
	slice := strings.Fields(text)
	for i := range slice {
		slice[i] = strings.ToLower(strings.Trim(slice[i], `!,.-":;)`))
	}
	return slice
}
func sliceToMap(slice []string) map[string]int { // функция перевода среза в карту
	result := make(map[string]int)
	for _, i := range slice {
		result[i]++
	}
	return result
}
func output(result map[string]int) { // функция "красивого" вывода полученной карты
	fmt.Printf("%15s|%15s\n", "ключ", "количество")
	fmt.Println("---------------|---------------")
	for i, num := range result {
		if num > 1 {
			fmt.Printf("%15s|%15d\n", i, num)
			fmt.Println("---------------|---------------")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,"-`)
		frequency[word]++
	}
	return frequency
}
