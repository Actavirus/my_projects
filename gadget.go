// тема: интерфейс - это набор методов, который должен поддерживаться некоторыми/несколькими значениями.
package main

import (
	"fmt"
	"github/headfirstgo/gadget"
)                       // импортируем пакет. Внимание! определение интерфейса в том пакете, где он используется, обеспечивает большую гибкость
type Player interface { // определяет тип интерфейса
	Play(string) // должен содержать метод Play с параметром string
	Stop()       // также необходим метод Stop
}

func playList(device Player, songs []string) { // здесь []string - квадратные скобки означают возможность нескольких значений; Player - допустимо любое значение, поддреживающее Player.
	for _, song := range songs { // перебираем все песни в цикле
		device.Play(song) // воспроизведение текущей песни
	}
	device.Stop() // плеер останавливается после завершения.
}
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder) // утверждение типа используется для перехода к значению TapeRecorder. Здесь ok - второе возвращаемое значение присваивается перемнной.
	if ok {
		recorder.Record() // вызов метода, определённого только для конкретного типа. Если исходным типом был тип TapeRecorder, для значения вызывается метод Record.
	} else {
		fmt.Println("Player was not a TapeRecorder") // в противном случае выдаётся сообщение об ошибке утверждения типа.
	}
}
func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"} // создаётся сегмент с названиями песен.
	var player Player = gadget.TapePlayer{}                   // обновляем переменную для хранения любого значения, поддерживающего Player.
	playList(player, mixtape)                                 // TypePlayer передаётся playList.
	player = gadget.TapeRecorder{}                            // используем метод TypeRecorder.
	playList(player, mixtape)                                 // TypePlayer передаётся playList.
	TryOut(gadget.TapeRecorder{})                             // функции передаётся значение типа TapeRecorder (которое поддерживает интерфейс Player).
	TryOut(gadget.TapePlayer{})                               // выдаётся паника:"Исходный тип не совпадает с типом в утверждении".
}
