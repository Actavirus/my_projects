// тема - оператор Select, который работает как switch, но для каналов.
// обычно select используется для таймеров.
// в данном случае мы не заинтересованы в значениях временных меток, поэтому мы не сохраняем его в переменные.
// также мы можем задать команды, которые выполняются по умолчанию, используя конструкцию default.
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 2", msg2)
	case <-time.After(time.Second): // здесь time.After - создайт канал, по которому посылаем метки времени с заданным интервалом
		fmt.Println("timeout")
	default:
		fmt.Println("nothing ready")
	}
	var input string
	fmt.Println(&input)
}
