// программа-календарь для указания даты и присвоения этой дате какого-либо события
package main

import (
	"fmt"
	"github/headfirstgo/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetYear(1959) // этот set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(11) // этот set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(9) // этот set-метод типа Date был повышен до Event
	if err != nil {
		log.Fatal(err)
	}
	err2 := event.SetTitle("Mom's Birthday") // определяется для типа Event
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(event.Day(), event.Month(), event.Year(), event.Title())
}
