// программа для вычисления квадратного корня из введённого числа
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("введено отрицательное число")
	}
	return math.Sqrt(number), nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)   //создаём bufio.Reader для чтения ввода с клавиатуры
	fmt.Print("Укажите число:")           //запрашиваем число
	input, err := reader.ReadString('\n') //прочитать данные, введенные пользователем до нажатия Enter
	if err != nil {                       //если произошла ошибка, программа выводит сообщение и завершается
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)            //удаление символа новой строки \n
	guess, err := strconv.ParseFloat(input, 64) //входная строка преобразуется в целое число типа int (для типа float64 используется функция ParseFloat)

	root, err := squareRoot(guess) //передаём полученные данные в функцию squareRot
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%.3f\n", root)
	}
}
