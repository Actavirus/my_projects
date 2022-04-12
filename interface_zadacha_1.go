// задача из учебника: Напишите программу, что выводит координаты в формате JSON, расширяя работу, сделанную для предварительной быстрой проверки. JSON вывод должен предоставить все координаты в десятичных градусах (DD), а также в градусах, минутах и секундах - мой вариант решения.
package main

import (
	"encoding/json"
	"fmt"
)

type coordinate6 struct {
	D float64 `json:"degrees"`
	M float64 `json:"minutes"`
	S float64 `json:"seconds"`
	H rune    `json:"hemisphere"`
}

func (c coordinate6) decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'W', 'S', 'w', 's':
		sign = -1.0
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func (c coordinate6) String() string {
	bytes, _ := json.MarshalIndent(c, "", "")
	return fmt.Sprintln("a", string(bytes)) // выводит: {"D":135,"M":54,"S":0,"H":69}
}
func main() {
	coor := coordinate6{135, 54, 0.0, 'E'}
	fmt.Println(coor)
}
