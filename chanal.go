// тема - Каналы. каналы обеспечивают возможность общения нескольких горутин друг с другом, чтобы синхронизировать их выполнение.
// программа будет постоянно выводить ping (нажмите enter, чтобы её остановить).
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) { // chan
	for i := 0; ; i++ {
		c <- "ping" // оператор <- (стрелка влево) - используется для отправки и получения сообщений по каналу. Конструкция (c <- "ping") - означает отправку ping.
	}
}
func printer(c chan string) {
	for {
		msg := <-c                  // конструкция (msg := <-c) - означает получение ping и сохранение в переменную msg.
		fmt.Println(msg)            // msg можно опустить и записать эту строку в виде: fmt.Println(<-c)
		time.Sleep(time.Second * 1) // можно указать Millisecond, тогда скорость передачи данных будет быстрее.
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func main() {
	var c chan string = make(chan string) // тип канала представлен ключевым словом chan, за которым следует тип, который будет передаваться по каналу (например, string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input) // если указать Println, то вывод будет не таким, какой мы ожидали.
}
