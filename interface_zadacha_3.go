// Создание игры-симулятора фермы в Golang - мой варинат решения
package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name    string `json:"Name"`
	Habitat string `json:"Habitat"`
	Eat     string `json:"Eat"`
}

func (a animal) String() string {
	bytes, _ := json.MarshalIndent(a, "", "")
	return fmt.Sprintf(string(bytes))
}
func (a animal) motion() string {
	return fmt.Sprintln("Eating")
}
func main() {
	cat := animal{Name: "Cat", Habitat: "Land", Eat: "Mouse"}
	// fish := animal{Name: "Fish", Habitat: "Water"}
	// bird := animal{Name: "Bird", Habitat: "Air"}
	fmt.Println(cat)
}
