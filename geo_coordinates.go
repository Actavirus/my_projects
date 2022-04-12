// программа-пример использования set и get методов, а так же инкапсуляции (сокрытия данных одной части прграммы от другой части)
package main

import (
	"fmt"
	"github/headfirstgo/coordinates"
	"log"
)

func main() {
	location := coordinates.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(150.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
