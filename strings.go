// пакет datafile предназначен для чтения данных из файлов
package datafile

import (
	"bufio"
	"os"
)

// GetStrings читает строку из каждой строки данных файла
func GetStrings(fileName string) ([]string, error) { // здесь []string возвращает сегмент строк вместо сегмента float64
	var lines []string // в переменной lines хранится сегмент строк
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()      // вместо преобразования строки файла к типу float64...
		lines = append(lines, line) // ...мы просто добавляем её к сегменту
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil // возвращает сегмент строк
}
