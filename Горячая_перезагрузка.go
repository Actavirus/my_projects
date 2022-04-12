// Горячая перезагрузка с использованием Air в golang (https://dev-gang.ru/article/gorjaczaja-perezagruzka-s-ispolzovaniem-air-v-golang-xnsqsln98p/)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Сначала создадим http-сервер на golang, размещенный на порту «5000»:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Golang Dev"))
	})
	PORT := "5000"
	fmt.Println("Server running on port: ", PORT)
	fmt.Println("This file is called: Горячая перезагрузка.go")
	http.ListenAndServe(":"+PORT, nil)
}
