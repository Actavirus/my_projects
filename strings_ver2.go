// работа со строками
package main

import "fmt"

func main() {
	fmt.Println(`strings can span multiple lines with the \n escape sequence`) // пример использования обратных кавычек
	fmt.Println(`
	peace be upon you
	upon you be peace`) // другой пример использования обратных кавычек
	fmt.Printf("%v is a %[1]T\n", "literal string")     // выводит: literal string is a string
	fmt.Printf("%v is a %[1]T\n", `raw string literal`) // выводит: raw string literal is a string
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 31
	fmt.Printf("%v%v%v%v\n", pi, alpha, omega, bang) // выводит: 96094096933
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang) // выводит: πάω▼
	// задача из учебника: Напишите программу для вывода каждого байта (символ ASCII) слова «shalom», по одному символу на строку.
	word := "Hellow my sister!"
	for i := 0; i < len(word); i++ { // len - определение длины word (в байтах)
		letter := word[i]
		fmt.Printf("%c - %[1]v\n", letter)
	}
}
