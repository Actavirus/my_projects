// Условие задачи №1 (https://cups.online/)
// Маша продает куклы ручной работы на маркетплейсе. 
// Для оптимизации стоимости звонков клиентам она хочет собрать информацию по мобильным операторам, 
// которыми пользуются заказчики. Таким образом, зная номер мобильного телефона, 
// нужно определить код оператора связи. 
// Номера представлены в виде 11-значных чисел, где первая цифра - код страны, 
// следующие три цифры – код оператора, а остальные 7 – номер абонента.
package main
import (
	"fmt"
	"os"
	"strconv"
)
func main(){
	var phoneNumber int
	fmt.Print("Введите номер телефона (8ХХХХХХХХХХ): ")
	fmt.Fscan(os.Stdin, &phoneNumber)
	if phoneNumber <= 10000000000 || phoneNumber >= 100000000000 {
		fmt.Println("Введён не правильный номер. Введите в формате 8ХХХХХХХХХХ")
	} else {
		fmt.Printf("Вы ввели следующий номер телефона: %v\n", phoneNumber)
	}
	numberIntToFloat := float64(phoneNumber)
	fmt.Printf("%v type is %T\n", phoneNumber, numberIntToFloat)
	fmt.Println("", (phoneNumber / 10000000)%1000)
	numberIntToString := strconv.Itoa(phoneNumber)
	fmt.Printf("%v type is %T\n", phoneNumber, numberIntToString)
	numberFloatToString := fmt.Sprint(numberIntToFloat)
	fmt.Printf("%v type is %T\n", phoneNumber, numberFloatToString)
}