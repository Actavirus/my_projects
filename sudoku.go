// задача: Создание игры Судоку в Golang
package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9 // количество строк и столбцов
	empty         = 0    // необходимо для проверки условия: находится ли в ячейке "0"?
)

// Cell является квадратом сетки Судоку
type Cell struct {
	digit int8 // это цифра для подстановки в ячейку
	fixed bool // "фиксированный"; если в ячейке фиксированное значение не равно "0"
}

// Grid является сеткой Судоку
type Grid [rows][columns]Cell

// Ошибки, что могут возникнуть
var (
	ErrBounds     = errors.New("за пределами")
	ErrDigit      = errors.New("неправильная цифра")
	ErrInRow      = errors.New("эта цифра уже есть в ряду")
	ErrInColumn   = errors.New("эта цифра уже есть в столбце")
	ErrInRegion   = errors.New("в данной части эта цифра уже есть")
	ErrFixedDigit = errors.New("начальные цифры нельзя переписать")
)

// NewSudoku делает новую сетку Судоку
func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for r := 0; r < rows; r++ { // цикл для строк
		for c := 0; c < columns; c++ { // цикл для стобцов
			d := digits[r][c] // присваиваем значение ячейки переменной d
			if d != empty {   // проверяем условие на наличие в ячейке значения "0". Почему бы здесь просто не задать значение "0", вместо empty ?
				grid[r][c].digit = d    // ???
				grid[r][c].fixed = true // ???
			}
		}
	}
	return &grid
}

// Устанавливает цифры на сетку
func (g *Grid) Set(row, column int, digit int8) error {
	switch { // работаем через "переключатель"
	case !inBounds(row, column): // проверка условия на выход из границ сетки Судоку
		return ErrBounds // выводим ошибку если возвращаемое значение "не true", т.е. "false"
	case !validDigit(digit): // проверка условия записываемого значения в ячейку
		return ErrDigit // выводим ошибку если возвращаемое значение "не true", т.е. "false"
	case g.isFixed(row, column): // ??? - возможно этот метод сверяет ячейку с фиксированной ячейкой из функции NewSudoku (grid[r][c].fixed = true)
		return ErrFixedDigit // выводим ошибку если возвращаемое значение "true"
	case g.inRow(row, digit): // проверка на повторяющиеся значения в строке
		return ErrInRow // выводим ошибку если возвращаемое значение "true"
	case g.inColumn(column, digit): // проверка на повторяющиеся значения в столбце
		return ErrInColumn // выводим ошибку если возвращаемое значение "true"
	case g.inRegion(row, column, digit): //
		return ErrInRegion // выводим ошибку если возвращаемое значение "true"
	}

	g[row][column].digit = digit
	return nil
}

// Очищает клетку на сетке Судоку
func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}

	g[row][column].digit = empty
	return nil

}

// Функция для проверки на выход за границы сетки Судоку
func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns { // проверяем вхождение введённых строки и стоолбца в границы сетки Судоку
		return false
	}
	return true
}

// Функция для проверки правильного значения в ячейке
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9 // возвращает true или false в зависимости от выполненного условия
}

// Метод для проверки повторяющихся значений в строке
func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ { // проходим по конкретной строке
		if g[row][c].digit == digit { // ??? - g[row][c].digit
			return true
		}
	}
	return false
}

// Метод для проверки повторяющихся значений в столбце
func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ { // проходим по конкретному столбцу
		if g[r][column].digit == digit { // ??? - g[r][column].digit
			return true
		}
	}
	return false
}

// Метод для проверки повторяющихся значений в границах квадрата 3х3
func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3 // ??? - для чего делить и умножать на одно и то же число?
	for r := startRow; r < startRow+3; r++ {     //
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit { // условие проверки
				return true
			}
		}
	}
	return false
}

// Метод - что он делает??? - возможно этот метод сверяет ячейку с фиксированной ячейкой из функции NewSudoku (grid[r][c].fixed = true)
func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(0, 0, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s {
		fmt.Println(row)
	}
}
