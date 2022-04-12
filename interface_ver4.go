// тема: Разбираемся с интерфейсами в Go (https://habr.com/ru/company/vk/blog/463063/)
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type stringer interface {
	String() string
}

// Неважно, каким типом является book или что он делает. Важно лишь, что у него есть метод под названием String(), который возвращает строковое значение.
type book struct { // тип book удовлетворяет интерфейсу stringer, потому что у него есть строковый метод String()
	Title  string
	Author string
}
type Count int // Тип Count тоже удовлетворяет интерфейсу fmt.stringer, потому что у него есть метод с тем же сигнатурным строковым значением String().

func (b book) String() string {
	return fmt.Sprintf("book: %s - %s\n", b.Title, b.Author)
}
func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// А теперь самое важное: Когда вы видите в Go объявление (переменной, параметра функции или поля структуры), имеющее интерфейсный тип, вы можете использовать объект любого типа, пока он удовлетворяет интерфейсу.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

type Customer struct {
	Name string
	Age  int
}

// Реализуем метод WriteJSON, который берёт io.Writer в виде параметра.
// Он отправляет структуру Сustomer в JSON, и если всё отрабатывает
// успешно, то вызывается соответствующий метод Write() из io.Writer.
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}
	_, err = w.Write(js)
	return err
}
func main() {
	// Так что такое интерфейс?
	// Инициализируем объект book и передаём в WriteLog().
	book := book{Title: "Alice in Wonderland", Author: "Lewis Carrol"}
	WriteLog(book)
	// Инициализируем объект Count и передаём в WriteLog().
	count := Count(3)
	WriteLog(count)

	// Чем полезны интерфейсы?
	// Уменьшение количества шаблонного кода
	// Инициализируем структуру Customer.
	c := &Customer{Name: "Alice", Age: 21}
	// Затем с помощью Buffer можем вызвать метод WriteJSON
	var buf bytes.Buffer
	err := c.WriteJSON(buf) // возникает ошибка в этом месте: "cannot use buf (type bytes.Buffer) as type io.Writer in argument to c.WriteJSON: bytes.Buffer does not implement io.Writer (Write method has pointer receiver)"
	if err != nil {
		log.Fatal(err)
	}
	// или воспользоваться файлом.
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}

}
