// тема - Создание и итерация массива в Golang
package main

import "fmt"

func print(shahmat [8][8]rune) { // функция для вывода двумерного массива
	for i := range shahmat {
		fmt.Printf("%c\n", shahmat[i])
		i++
	}
}
func display(board [8][8]rune) { // это функция
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}
func main() {
	// объявление массива и получение доступа к его элементам
	var planets [8]string
	planets[0] = "Меркурий"
	planets[1] = "Венера"
	planets[2] = "Земля"
	earth := planets[2]
	fmt.Println(earth)        // выводит: Земля
	fmt.Println(len(planets)) // выводит: 8 (функция len определяет длину типов, в данном случае возвращается размер массива)

	// Инициализация массивов через композитные литералы в Go
	dwarfs := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Println(len(dwarfs)) // выводит: 5
	planets = [...]string{   // компилятор Go самостоятельно подсчитывает количество элементов
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун", // запятая в самом конце является обязательной
	}
	fmt.Println(planets) // выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Нептун]

	// Итерация через массивы в Go
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
	for i, dwarf := range dwarfs { // использование ключевого слова range
		fmt.Println(i, dwarf)
	}

	// Копирование массивов в Golang
	planetsMarkII := planets   // копирует массив planets
	planets[2] = "упс"         // прокладывает путь для межзвездного шунтирования
	fmt.Println(planets)       // выводит: [Меркурий Венера упс Марс Юпитер Сатурн Уран Нептун]
	fmt.Println(planetsMarkII) // выводит: [Меркурий Венера Земля Марс Юпитер Сатурн Уран Нептун]

	// Массивы из массивов в Golang
	var board1 [8][8]string // массив из восьми массивов с восемью строками (типо шахматная доска)
	board1[0][0] = "r"
	board1[0][7] = "r" // ставит ладью на клетку с координатами [ряд][столбец]
	for column := range board1[1] {
		board1[1][column] = "p"
	}
	for i := range board1 {
		fmt.Println(board1[i])
	}

	// задача: отображение шахматной доски с фигурами - мой вариант решения.

	shahmat := [8][8]rune{
		{'n', 'b', 'r', 'k', 'q', 'r', 'b', 'n'},
		{},
		{},
		{},
		{},
		{},
		{},
		{'N', 'B', 'R', 'K', 'Q', 'R', 'B', 'N'},
	}
	for column := range shahmat[7] {
		shahmat[1][column] = 'p'
		shahmat[6][column] = 'P'
	}
	fmt.Println("=================")
	print(shahmat)
	fmt.Println("=================")

	// задача: отображение шахматной доски с фигурами - решение из учебника.
	var board [8][8]rune
	// черные фигуры
	board[0][0] = 'r'
	board[0][1] = 'n'
	board[0][2] = 'b'
	board[0][3] = 'q'
	board[0][4] = 'k'
	board[0][5] = 'b'
	board[0][6] = 'n'
	board[0][7] = 'r'
	// пешки
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	// белые фигуры
	board[7][0] = 'R'
	board[7][1] = 'N'
	board[7][2] = 'B'
	board[7][3] = 'Q'
	board[7][4] = 'K'
	board[7][5] = 'B'
	board[7][6] = 'N'
	board[7][7] = 'R'
	display(board)
}
