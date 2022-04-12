// Go: десериализация JSON с неправильной типизацией, или как обходить ошибки разработчиков API
// (https://habr.com/ru/post/502176/)
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CustomBool struct {
	Bool bool
}

func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		cb.Bool = true
		return nil
	case `"false"`, `false`, `"0"`, `0`, `""`:
		cb.Bool = false
		return nil
	default:
		return errors.New("CustomBool: parsing \"" + string(data) + "\":unknown value")
	}
}
func (cb CustomBool) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cb.Bool)
	return json, err
}

type Target struct {
	Id     int        `json:"id"`
	Active CustomBool `json:"active"`
}

func main() {
	jsonString := `[{"id":1,"active":true},
	{"id":2,"active":"true"},
	{"id":3,"active":"1"},
	{"id":4,"active":1},
	{"id":5,"active":false},
	{"id":6,"active":"false"},
	{"id":7,"active":"0"},
	{"id":8,"active":0},
	{"id":9,"active":""}]`
	targets := []Target{}
	_ = json.Unmarshal([]byte(jsonString), &targets)
	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Active.Bool)
	}
	jsonStringNew, _ := json.Marshal(targets)
	fmt.Println(string(jsonStringNew))
}
