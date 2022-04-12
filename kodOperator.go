package main
import (
	"fmt"
	"os"
	"errors"
)
func main(){
	var phoneNumber int
	fmt.Print("Введите номер телефона (8ХХХХХХХХХХ): ")
	fmt.Fscan(os.Stdin, &phoneNumber)
	err := proverka(phoneNumber)
	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println("", (phoneNumber / 10000000)%1000)
	}
}
func proverka(p int) error {
	// err := errors.New
	if p <= 10000000000 || p >= 100000000000 {
		return errors.New("Введён не правильный номер. Введите в формате 8ХХХХХХХХХХ")
	}
	return nil
}