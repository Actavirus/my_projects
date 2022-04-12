// тема - Работа со строками в Golang
// программа дешифратор соообщения
package main

import "fmt"

func main() {
	// letter := 'a' // если тип не уточняется, Go присвоит тип rune для символа в одинарных кавычках (символьный тилитерал). Тип rune иначе называется int32
	// letter += 3	// данный код не подойдёт для слов с буквами ‘x’, ‘y’ и ‘z’
	// fmt.Println(letter) // выводит: 97
	// fmt.Printf("\n", letter)   // выводит: %!(EXTRA int32=97)
	// fmt.Printf("%c\n", letter) // выводит: а
	// fmt.Printf("%T\n", letter) // выводит: int32
	// fmt.Printf("%v\n", letter) // выводит: 97

	message1 := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message1); i++ {
		c := message1[i]
		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы, т.е. не включает в расчёт знаки препинания.
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	message2 := "L fdph, L vdz, L frqtxhuhg"
	for i := 0; i < len(message2); i++ {
		c2 := message2[i]
		if c2 >= 'a' && c2 <= 'z' {
			c2 -= 3
			if c2 < 'a' {
				c2 += 26
			}
		} else if c2 >= 'A' && c2 <= 'Z' {
			c2 -= 3
			if c2 < 'A' {
				c2 += 26
			}
		}
		fmt.Printf("%c", c2)
	}
	fmt.Println()
}
