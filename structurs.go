// работа с структурами
package main

import (
	"fmt"
	"github/headfirstgo/geo"
	"github/headfirstgo/magazine"
)

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}
type student struct {
	name  string
	grade float64
}

func showInfo(p part) { // объявление одного параметра с типом part
	fmt.Println("Description:", p.description) // обращение к полям параметра
	fmt.Println("Count:", p.count)             // обращение к полям параметра
}

func minimumOrder(description string) part { // здесь psrt - объявляется одно возвращаемое значение типа part
	var p part // объявление переменной p типа part
	p.description = description
	p.count = 100
	return p // функция возвращает тип part
}

func printInfo(s student) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var myStruct struct { // объявление переменной с именем myStruct, которая может хранить структуры, состоящие из...
		number float64 // ...поля float64 с именем number (здесь number - имя поля, float64 - тип поля)...
		word   string  // ...поля string с именем word...
		toggle bool    // ...поля bool с именем toggle.
	}
	fmt.Printf("%#v\n", myStruct) // значение структуры выводится в том виде, в котором оно записывается в коде Go
	// пример присвоения значения структуре
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	// пример вывода значения структуры
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
	// упражнение из учебника
	var animal struct {
		name string
		age  int
	}
	animal.name = "Max"
	animal.age = 5
	fmt.Println(animal.name)
	fmt.Println(animal.age)

	// использование type
	var porsche car // объявление переменной типа car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)
	var bolts part // объявление переменной типа bolts
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)

	// использование определяемый тип part с функцией showInfo
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts) // тип part передаётся функции showInfo

	// здесь функция minimumOrder создаёт значение part с заданным описанием description и заранее определённым значением поля count
	p := minimumOrder("Hex bolts") // вызывает minimumOrder. Для хранения возвращаемого значения part используется короткое определение переменной
	fmt.Println(p.description, p.count)

	// упражнение из учебника
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(student(s))

	// литералы структур
	myCar := car{name: "Corvette", topSpeed: 337}
	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true} // объявление переменной структурного типа методом литералов
	fmt.Println(myCar)
	fmt.Println(subscriber)

	// упражнение из учебника
	location := geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude:", location.Latitude)
	fmt.Println("Longitude:", location.Longitude)

	// использование структурного типа из другого файла
	employee := magazine.Employee{Name: "Joy carr", Salary: 60000}
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	// заполнение полей структуры Addrress внутри структуры Subscriber
	// первый способ
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"} // создание значения Address с заполнением полей
	subscriber = magazine.Subscriber{Name: "Aman Singh"}                                               // создание структуры Subscriber, который будет принадлежать Address
	subscriber.HomeAddress = address                                                                   // инициализирует поле HomeAddress
	fmt.Println(subscriber.HomeAddress)                                                                // выводит данные Address
	fmt.Println(subscriber)                                                                            // выводит данные Subscriber и Address
	//второй способ
	subscriber.HomeAddress.PostalCode = "68222"                    // присваивание значения полю PostalCode через структуру Subscriber
	fmt.Println("Postal code:", subscriber.HomeAddress.PostalCode) // чтение значения поля PostalCode через структуру Subscriber

	//пробуем упрощённую схему присвоения значения
	employee.Name = "Margarete O'nill"
	employee.City = "New York"
	fmt.Println("Name:", employee.Name, "City:", employee.City)

	//задача из учебника
	location2 := geo.Landmark{}
	location2.Name = "The Googleplex"
	location2.Latitude = 37.42
	location2.Longitude = -122.08
	fmt.Println(location2)
}
