// первый способ ввода данных пользователем через клавиатуру
// с использованием дополнительной функции getFloat
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите число : ")
	x, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Вы ввели:", x)
}
func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil { // если при чтении ввода произошла ошибка, она будет возвращена функцией
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil { // также возвращаются ошибки, которые могут возникнуть при преобразовании строки в float64
		return 0, err
	}
	return number, nil
}
