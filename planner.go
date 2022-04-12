package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Ваша встреча через", days, "дней")
	fmt.Println("Следующая встреча через", days+dates.DaysInWeek, "дней")
}
