// Go: десериализация JSON с неправильной типизацией, или как обходить ошибки разработчиков API
// (https://habr.com/ru/post/502176/)
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// В качестве пользовательского типа со своей логикой десериализации объявим структуру CustomFloat64 с единственным полем Float64 типа float64:
type CustomFloat64 struct {
	Float64 float64
}

const QUOTES_BYTE = 34

// На входе данный метод принимает слайс байт (data), в котором содержится значение конкретного поля декодируемого json.
// Если преобразовать данную последовательность байт в строку, то мы увидим значение поля именно в том виде, в каком оно записано в json.
func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == QUOTES_BYTE { // проверяем первый байт слайса data, сравнив его с номером символа в UNICODE. Это номер 34.
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64) // необходимо получить строку между кавычками, т. е. слайс байт между первым и последним байтом.
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}
func (cf CustomFloat64) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(cf.Float64)
	return json, err
}

type Target struct {
	Id    int           `json:"id"`
	Price CustomFloat64 `json:"price"` // Укажем тип для поля Price в структуре Target
}

func main() {
	// Т. е. json содержит массив объектов с двумя полями числового типа:
	jsonString := `[{"id":1,"price":2.58},
	{"id":2,"price":"2.58"},
	{"id":3,"price":7.15},
	{"id":4,"price":"7.15"}]`
	targets := []Target{}
	_ = json.Unmarshal([]byte(jsonString), &targets) // В данном коде мы десериализуем поле id в тип int, а поле price в тип float64.
	for _, t := range targets {
		fmt.Println(t.Id, "-", t.Price.Float64)
	}
	jsonStringNew, _ := json.Marshal(targets)
	fmt.Println(string(jsonStringNew))
}
