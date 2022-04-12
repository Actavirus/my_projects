// задача из учебника №1: Скучно, когда одинаковая строка кода повторяется из раза в раз.
// Напишите элемент конвейера (горутину), что запоминает предыдущее значение и только отправляет значение на следующий этап конвейера, если оно отличается от того, что пришло ранее.
// Чтобы все упростить, можете предположить, что первая строка никогда не бывает простой.

// задача из учебника №2: Иногда проще оперировать словами, а не предложениями.
// Напишите конвейер, что принимает строки и разделяет их на слова (можете использовать функцию Fields из пакета strings), а также отправляет все слова, одно за другим, на следующую стадию конвейера.
package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	go divideGopher(c1, c2)
	printGopher(c2)
}
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "a bad apple", "goodbye all", "morning sun", "next over space", "a bad cat", "a bad apple", "dog is mine", "reading copybook"} {
		downstream <- v
	}
	close(downstream)
}

// решение к задаче №1 - мой вариант решения
func filterGopher(upstream, downstream chan string) {
	v := ""
	for item := range upstream {
		if item != v {
			v = item
			downstream <- item
		}
	}
	close(downstream)
}

// решение к задаче №2 - мой вариант решения
func divideGopher(upstream, downstream chan string) {
	// sl := make([]string, 10)	// инициализация слайса(среза) необязательна
	for item := range upstream {
		sl := strings.Fields(item)
		for _, word := range sl {
			downstream <- word
		}
	}
	close(downstream)
}
func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
