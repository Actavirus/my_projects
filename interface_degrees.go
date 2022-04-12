// тема - Интерфейс error.
// программа получает два значения температуры (факт и максимально допустимую) и в случае превышения фаткической выводит сообщение
package main

import (
	"fmt"
	"log"
)

type OverheatError float64 // определяем тип с базовым типом float64
func (o OverheatError) Error() string { // поддерживает интерфейс error (интерфейс/тип error является "предварительно объявленным идентификатором", как int или string).
	return fmt.Sprintf("Overheating by %0.2f degrees!", o) // температура используется в сообщении об ошибке.
}
func checkTemperature(actual float64, safe float64) error { // здесь error - указывает, что функция возвращает обычное значение ошибки.
	excess := actual - safe
	if excess > 0 { // еси фактическая температура превышает безопасную...
		return OverheatError(excess) // ...возвращается значение OverheatError с превышением безопасной температуры.
	}
	return nil
}
func main() {
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
