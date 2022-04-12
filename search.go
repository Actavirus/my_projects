// программа, которая проверяет совпадение значений из двух разных сегментов
package main

import "fmt"

func main() {
	data := []string{"a", "c", "e", "a", "e"} // подсчёт количества вхождений каждой буквы в сегменте
	counts := make(map[string]int)            // карта для хранения счетчиков
	for _, item := range data {               // обработка каждой буквы
		counts[item]++ // увеличение счетчика для текущей буквы
	}
	letters := []string{"a", "b", "c", "d", "e"} // проверяем, существует ли каждая из этих букв в карте
	for _, letter := range letters {
		count, ok := counts[letter] // получение счетчика для текущей буквы. а также признака её обнаружения
		if !ok {                    // если буква не найдена...
			fmt.Printf("%s: не найден\n", letter) // ...сообщить об этом.
		} else { //  в противном случае буква была найдена...
			fmt.Printf("%s: %d\n", letter, count) // ...тогда выводится буква и счетчик.
		}
	}
}