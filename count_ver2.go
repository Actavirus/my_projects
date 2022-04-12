// программа подсчёта голосов из файла с использованием карт
package main

import (
	"fmt"
	"github/headfirstgo/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int) // объявляем карту, у которой ключами являются имена кандидатов, а значениями - счетчики голосов
	for _, line := range lines {
		counts[line]++ // увеличивает счетчик голосов для текущего кандидата
	}
	for name, count := range counts { // обработка каждого ключа и значения в карте
		fmt.Printf("Голосов за %s: %d\n", name, count) // здесь name - вывод ключа (имя кандидата), count - вывод голосов (количество голосов)
	}
}
