// тема - Создаем Шифр Виженера на Golang
// Напишем программу для дешифровки зашифрованного текста
package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 1
	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// зашифрованная буква - ключевая буква
		c = (c-k+26)%26 + 'A' // % - это остаток от деления!!!
		message += string(c)

		// увеличение keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)
	// fmt.Println(keyIndex)
	// fmt.Println(len(keyword))
	// keyIndex %= len(keyword)
	// fmt.Println(keyIndex)
	a := 2 % 6
	fmt.Println(a)
}
