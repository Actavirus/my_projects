// Обработка ошибок в Golang (https://golangify.com/errors)
package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Исправление ошибок в Golang
	// Если возникает ошибка, переменная err не будет равна nil, что заставит программу вывести ошибку и немедленно завершить работу.
	// Ненулевое значение, переданное os.Exit, сообщает операционной системе, что произошла ошибка.
	// Если ReadDir успешно выполнена, files будет назначен к срезу os.FileInfo, предоставляющий информацию о файлах и каталогах по указанному пути.
	files, err := ioutil.ReadDir(".") // В данном случае точка уточняет путь, указывающий текущую директорию.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		fmt.Println(file.IsDir()) // Вы можете использовать file.IsDir() для того, чтобы различить файлы от директорий.
	}

	// ----------------------------------------------------------------------------------------------
	// Элегантная обработка ошибок в Golang
	// Запись данных в файле (использование функции proverbs)
	err = proverbs("proverbs.txt") // вызывает proverbs для создания файла
	if err != nil {                // обрабатывает любые ошибки, отображая ее и затем выходя.
		fmt.Println(err)
		os.Exit(1)
	}

	// ----------------------------------------------------------------------------------------------
	// Применяем defer — отложенные действия в Golang (использование функции proverbs_2)
	err = proverbs_2("proverbs_2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ----------------------------------------------------------------------------------------------
	// Креативная обработка ошибок в Golang (использование метода proverbs_3)
	// Это более простой и чистый способ для записи текстового файла, но смысл не в этом.
	// Такая же техника может использоваться для создания zip-файлов или для совершенно разных задач.
	err = proverbs_3("proverbs_3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ----------------------------------------------------------------------------------------------
	// Новые ошибки в программе на Golang
	// Для демонстрации новых ошибок Листинг 7 создает основу для Судоку, что представляет собой сетку 9 на 9.
	// Каждый квадрат сетки может содержать цифру от 1 до 9. Имплементация использует массив с фиксированным размером, ноль указывает на пустой квадрат.
	// Создает сетку и отображает любую ошибку, возникшую в результате неправильной замены.
	// Фраза «out of bounds» неплоха, но более точное «outside of grid boundaries» может быть лучше. А сообщение «error 37» вообще ни о чем не говорит.
	var g Grid
	err = g.Set(1, 8, 5)
	if err != nil {
		fmt.Printf("Произошла ошибка: %v.\n", err) // На заметку: Для сообщений ошибок часто используются части предложений, чтобы перед отображением к ним можно было добавить дополнительный текст.
		// os.Exit(101)	// Используйте os.Exit для немедленного выхода с полученным статусом.
	}

	// ----------------------------------------------------------------------------------------------
	// Причины каждой ошибки в Go
	// Многие пакеты Go объявляют и экспортируют переменные для ошибок, которые они могут вернуть.
	// Для использования этого с сеткой Судоку следующий листинг объявляет две переменные для ошибок на уровне пакета (смотри за границами main).
	// Если метод Set_2 возвращает ошибку, вызывающая сторона может различить возможные ошибки и обрабатывать определенные ошибки по-разному, как показано в следующем листинге.
	// Вы можете сравнить ошибку, возвращаемую с переменными ошибки, используя == или оператор switch.
	err_2 := g.Set_2(0, 5, 5)
	if err_2 != nil {
		switch err_2 { // На заметку: Конструктор errors.New имплементируется через использование указателя, поэтому оператор switch сравнивает адреса памяти, текст не содержит сообщения об ошибке.
		case ErrBounds:
			fmt.Println(ErrBounds)
		case ErrDigit:
			fmt.Println(ErrDigit)
		default:
			fmt.Println(err_2)
		}
		// os.Exit(102)	// Используйте os.Exit для немедленного выхода с полученным статусом.
	}

	// ----------------------------------------------------------------------------------------------
	// Настраиваемые типы ошибок в Golang
	// Тип error является встроенным интерфейсом, как показано в следующем примере.
	// Любой тип, что имплементирует метод Error() для возвращения строки, неявно удовлетворяет интерфейс.
	// В качестве интерфейса возможно создать новые типы ошибок.
	type error interface { // тип error является встроенным интерфейсом
		Error() string
	}

	// ----------------------------------------------------------------------------------------------
	// Множество ошибок в Golang
	// Вместо возвращения одной ошибки за раз, метод Set_3 может сделать несколько проверок и вернуть все ошибки сразу.
	// Тип SudokuError в Листинге 15 является срезом error (смотри за пределами main).
	// Он удовлетворяет интерфейсу error с методом, что соединяет ошибки вместе в одну строку.
	err_3 := g.Set_3(0, 2, 5)
	if err_3 != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("Ошибок нет")
	}

	// ----------------------------------------------------------------------------------------------
	// Утверждение типа в Go
	// Утверждение типа в Листинге 17 утверждает err для типа SudokuError через код err.(SudokuError).
	// Если так и есть, то ok будет истинным, а err будет SudokuError, давая доступ к срезам ошибок в данном случае.
	// Помните, что отдельные ошибки для SudokuError являются переменными ErrBounds и ErrDigit, что могут сравниваться в случае необходимости.
	err = g.Set_3(0, 0, 5) // передаём методы Set_3 числовые значения
	if err != nil {        // проверяем условие полученного значения на содержание "чего-либо"
		if errs, ok := err.(SudokuError); ok { // err.(SudokuError) - утверждение пытается конвертировать значение err из типа интерфейса error в конкретный тип SudokuError; если
			fmt.Printf("%d ошибка(и) найдена(ы): \n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		// os.Exit(104)	// для завершения выполнения кода
	} else {
		fmt.Println("Ошибок нет")
	}

	// ----------------------------------------------------------------------------------------------
	// задача из учебника: В стандартной библиотеке Go есть функция для парсинга веб адресов (см. golang.org/pkg/net/url/#Parse).
	// Отобразите ошибку, которая возникает, когда url.Parse используется для неправильного веб адреса вроде того, что содержит пробелы: https://a b.com/.
	u, err := url.Parse("https://a b.com/") // "разобрать"
	if err != nil {
		fmt.Println(err)         // выводит: parse "https://a b.com/": invalid character " " in host name
		fmt.Printf("%#v\n", err) // выводит: &url.Error{Op:"parse", URL:"https://a b.com/", Err:" "}
		if e, ok := err.(*url.Error); ok {
			fmt.Println("0p:", e.Op)   // выводит: 0p: parse
			fmt.Println("URL:", e.URL) // выводит: URL: https://a b.com/
			fmt.Println("Err:", e.Err) // выводит: Err: invalid character " " in host name
		}
		// os.Exit(105)
	}
	fmt.Println(u) // чтобы данное сообщение не выводилось, необходимо разкомментировать "os.Exit()"

}
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close() // закрытие файла
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	f.Close() // закрытие файла
	return err
}
func proverbs_2(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close() // Убедиться, что файл правильно закрыт, можно с помощью ключевого слова defer.
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}

// ----------------------------------------------------------------------------------------------
type safeWriter struct {
	w   io.Writer
	err error // место для хранения первой ошибки
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // пропускает запись, если раньше была ошибка
	}
	_, sw.err = fmt.Fprintln(sw.w, s) // записывает строку и затем хранит любую ошибку.
}
func proverbs_3(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	return sw.err // возвращает ошибку в случае её возникновения
}

// ----------------------------------------------------------------------------------------------
const rows, columns = 9, 9

type Grid [rows][columns]int8 // Grid является сеткой Судоку
// Пакет errors содержит функцию конструктора, что принимает строку для сообщения об ошибке.
//Используя ее, метод Set в «Листинге 8» может создать и возвратить ошибку "out of bounds".
func (g Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) { // Проверка параметров в начале метода защищает оставшуюся часть метода от неправильного ввода.
		return errors.New("за границами") // "out of bounds"
	}
	g[row][column] = digit
	fmt.Println("Ошибок нет")
	return nil
}

// Функция inBounds помогает убедиться, что row и column находятся в пределах границ сетки. Она не дает методу Set забиться лишними деталями.
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

// ----------------------------------------------------------------------------------------------
var (
	ErrBounds = errors.New("Ошибки в количестве строк и/или столбцов") // На заметку: Принято присваивать сообщения об ошибках переменным, что начинаются со слова Err.
	ErrDigit  = errors.New("Ошибки в указанных числах")                // На заметку: Принято присваивать сообщения об ошибках переменным, что начинаются со слова Err.
)

// По объявлении ErrBounds вы можете изменить метод Set для возвращения его вместо создания новой ошибки, как показано в следующем коде.
func (g Grid) Set_2(row, column int, digit int8) error {
	if !inBounds(row, column) { // что означает знак "!" ? - полагаю, это тот же смысл что и в сочетании !=
		return ErrBounds
	}
	fmt.Println("Ошибок в количестве строк и столбцов нет")
	if !validDigit(digit) {
		return ErrDigit
	}
	fmt.Println("Ошибок в указанных числах нет")
	g[row][column] = digit
	return nil
}

// ----------------------------------------------------------------------------------------------
// задача из учебника: Напишите функцию validDigit и используйте ее, чтобы убедиться, что метод Set_2 принимает только цифры между 1 и 9.
func validDigit(digit int8) bool {
	if digit < 1 || digit > 9 {
		return false
	}
	return true
}

// ----------------------------------------------------------------------------------------------
type SudokuError []error // Тип SudokuError является срезом error
func (se SudokuError) Error() string { // Error возвращает одну или несколько ошибок через запятые.
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // конвертирует ошибки в строки
	}
	return strings.Join(s, ", ") // правило вывода информации (через запятую)
}

// Чтобы использовать SudokuError, метод Set_3 можно модифицировать для валидации границ и цифр, возвращая обе ошибки сразу:
var errs SudokuError

func (g *Grid) Set_3(row, column int, digit int8) error {
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds) // записываем значение ошибки в errs
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit) // записываем значение ошибки в errs
	}
	if len(errs) > 0 { // если в errs что-либо записалось, то...
		return errs // ...выводим значение errs
	}
	g[row][column] = digit
	return nil // возвращает nil, если условия не выполнились.
}
