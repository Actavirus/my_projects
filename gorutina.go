// тема - Многопоточность (http://golang-book.ru/chapter-10-concurrency.html).
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Горутины
	// Горутина - это функция, которая может работать параллельно с другими функциями (например, функция main сама по себе является горутиной).
	for i := 0; i < 10; i++ { // цикл для запуска нескольких горутин.
		go f(0) // запуск горутины.
	}
	// С горутиной программа немедленно перейдет к следующей строке (var input string), не дожидаясь, пока вызываемая функция завершится.
	// Вот почему здесь присутствует вызов Scanln - без него программа завершится еще перед тем, как ей удастся вывести числа.
	// Т.е., данный вызов Scanln нужен для того, чтобы понять, что горутина была выполнена (или можно использовать Println)
	var input string
	fmt.Scanln(&input)

	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Каналы - обеспечивают возможность общения нескольких горутин друг с другом, чтобы синхронизировать их выполнение.
	c := make(chan string) // инициализируем канал с
	go pinger(c)
	go printer(c)
	// Давайте добавим ещё одного отправителя сообщений в программу и посмотрим, что будет.
	go ponger(c) // Теперь программа будет выводить на экран то ping, то pong по очереди.
	fmt.Scanln(&input)

	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Направление каналов
	// Мы можем задать направление передачи сообщений в канале, сделав его только отправляющим или принимающим. Например, мы можем изменить функцию pinger_2... (смотри за границами main)
	go pinger_2(c)
	go printer_2(c)
	fmt.Scanln(&input)
	// Существуют и двунаправленные каналы, которые могут быть переданы в функцию, принимающую только принимающие или отправляющие каналы.
	// Но только отправляющие или принимающие каналы не могут быть переданы в функцию, требующую двунаправленного канала!

	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Оператор Select
	// В языке Go есть специальный оператор select который работает как switch, но для каналов.
	// Эта программа выводит «from 1» каждые 2 секунды и «from 2» каждые 3 секунды.
	// Оператор select выбирает первый готовый канал, и получает сообщение из него, или же передает сообщение через него.
	// Когда готовы несколько каналов, получение сообщения происходит из случайно выбранного готового канала.
	// Если же ни один из каналов не готов, оператор блокирует ход программы до тех пор, пока какой-либо из каналов не будет готов к отправке или получению.
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	fmt.Scanln(&input)
	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Обычно select используется для таймеров:
	// time After создаёт канал, по которому посылаем метки времени с заданным интервалом.
	// В данном случае мы не заинтересованы в значениях временных меток, поэтому мы не сохраняем его в переменные.
	// Лучше запускать отдельно от остального кода (при выводе в консоль не понятно как работает код)
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second): // так же можно записать: (time.Millisecond * 500)
				fmt.Println("timeout")
			}
		}
	}()
	fmt.Scanln(&input)
	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Также мы можем задать команды, которые выполняются по умолчанию, используя конструкцию default:
	// Выполняемые по умолчанию команды исполняются сразу же, если все каналы заняты.
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("northing ready") // по факту, "northing ready" выводится постоянно...
			}
		}
	}()
	fmt.Scanln(&input)

	fmt.Println("------------------------------------------------------------------------------------------------------")
	// Буферизированный канал
	// При инициализации канала можно использовать второй параметр:
	c3 := make(chan int, 1) // ...и мы получим буферизированный канал с ёмкостью 1.
	// Обычно каналы работают синхронно - каждая из сторон ждёт, когда другая сможет получить или передать сообщение.
	// Но буферизованный канал работает асинхронно — получение или отправка сообщения не заставляют стороны останавливаться.
	// Но канал теряет пропускную способность, когда он занят, в данном случае, если мы отправим в канал 1 сообщение,
	// то мы не сможем отправить туда ещё одно до тех пор, пока первое не будет получено.

}
func f(n int) { // функция f выводит числа от 0 до 10, ожидая от 0 до 250 мс после каждой операции вывода. Теперь горутины должны выполняться одновременно.
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250)) // добавили небольшую задержку функции с помощью time.Sleep и rand.Intn
		time.Sleep(time.Millisecond * amt)   // Или короткая запись: time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}
func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping" // конструкция означает отправку "ping" в канал с. Когда pinger пытается послать сообщение в канал, он ожидает, пока printer будет готов получить сообщение. Такое поведение называется блокирующим.
	}
}
func printer(c chan string) {
	for {
		msg := <-c                         // конструкция означает получение из канала с и присвоение полученного значения переменной msg (в нашем случае по каналу идёт "ping")
		fmt.Println(msg)                   // или короткая запись: fmt.Println(<- c)
		time.Sleep(500 * time.Millisecond) // небольшая задержка в 1 сек
	}
}
func ponger(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}
func pinger_2(c chan<- string) { // ...и канал c будет только отправлять сообщение. Попытка получить сообщение из канала c вызовет ошибку компилирования.
	for i := 0; i < 10; i++ {
		c <- "ping_2"
	}
}
func printer_2(c <-chan string) { // ...и канал с будет только принимать сообщение.
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond) // небольшая задержка в 1 сек
	}
}
