// Простейший веб-сайт на Go (https://codex.so/go-web-server)
// Компилятор попросит вас подтвердить прослушку порта, поэтому соглашаемся. В браузере вводим “http://localhost:8080/” и видим на странице переданное сообщение.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Используем пакет net/http для создания сервера
	http.HandleFunc("/", sayhello) // устанавливает роутер

	// // Функция ListenAndServe принимает два параметра — порт соединения и функцию-обработчик, которая будет выполнена при запуске сервера. В нашем случае, она не задана.
	// err := http.ListenAndServe(":8080", nil) // устанавливает порт веб-сервера
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	// Подключение HTTPS
	// Если вы хотите использовать зашифрованный протокол https, то вам необходимо использовать функцию ListenAndServeTLS.
	// В этой функции помимо порта и функции-обработчика нужно указать путь к сертификату и к приватному ключу:
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAdSrve: ", err)
	}

}

// Любая функция-обработчик принимает два параметра:
// http.ResponseWriter — это структура которая описывает ответ
// *http.Request — указатель на запрос. Из этого параметра можно получать тело запроса, параметры POST, GET или заголовки.
func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}
