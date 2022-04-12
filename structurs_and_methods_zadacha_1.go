package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// задача из учебника: Используя Листинги 1, 2 и 3, напишите программу, что объявляет location для каждой локации из Таблицы 1. Код должен выводить каждую локацию в десятичных градусах. - моё решение задачи.
	coordSpirit := local{
		location2{14, 34, 6.2, 'S'},
		location2{175, 28, 21.5, 'E'},
	}
	// coordOpportunity := local{
	// 	location2{1, 56, 46.3, 'S'},
	// 	location2{354, 28, 24.2, 'E'},
	// }
	// coordCuriosity := local{
	// 	location2{4, 35, 22.2, 'S'},
	// 	location2{137, 26, 30.1, 'E'},
	// }
	coordInSight := local{
		location2{4, 30, 0.0, 'N'},
		location2{135, 54, 0, 'E'},
	}
	g := perevod(coordInSight.Lat)
	h := perevod(coordInSight.Long)
	spirit2 := table{"Spirit", "Columbia Memorial Station", coordSpirit, coordSpirit}
	opportunity2 := table2{"Opportunity", "Challenger Memorial Station", local{location2{1, 56, 46.3, 'S'}, location2{354, 28, 24.2, 'E'}}}
	curiosity3 := table3{"Curiosity", "Bradbary Landing", location2{4, 35, 22.2, 'S'}, location2{137, 26, 30.1, 'E'}}
	insight := table4{"InSight", "Elysium Planitia", g, h}
	fmt.Printf("%s %s %.2f %.2f\n", spirit2.Module, spirit2.Landing, perevod(coordSpirit.Lat), perevod(coordSpirit.Long)) // Spirit Columbia Memorial Station 14.57 175.47
	fmt.Printf("%s %s %.2f %.2f\n", opportunity2.Module, opportunity2.Landing, perevod(opportunity2.LocalLatLong.Lat), perevod(opportunity2.LocalLatLong.Long))
	fmt.Printf("%s %s %.2f %.2f\n", curiosity3.Module, curiosity3.Landing, perevod(curiosity3.Lat), perevod(curiosity3.Long))
	bytes, _ := json.MarshalIndent(insight, "---", "|")
	fmt.Println(string(bytes))

	// второе решение задачи через использование типа struct в бОльшем формате
	table := []table5{
		{"Spirit", "Columbia Memorial Station", location2{14, 34, 6.2, 'S'}, location2{175, 28, 21.5, 'E'}},
		{"Opportunity", "Challenger Memorial Station", location2{1, 56, 46.3, 'S'}, location2{354, 28, 24.2, 'E'}},
		{"Curiosity", "Bradbary Landing", location2{4, 35, 22.2, 'S'}, location2{137, 26, 30.1, 'E'}},
		{"InSight", "Elysium Planitia", location2{4, 30, 0.0, 'N'}, location2{135, 54, 0, 'E'}},
	}
	fmt.Println(table)

}

type local struct {
	Lat  location2
	Long location2
}
type location2 struct {
	D, M, S float64 // d - градус, m - минута, s - секунда
	H       rune
}
type table struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     local
	Long    local
}
type table2 struct {
	Module       string `json:"The module:"`
	Landing      string `json:"Landing site"`
	LocalLatLong local
}
type table3 struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     location2
	Long    location2
}
type table4 struct {
	Module  string `json:"The module:"`
	Landing string `json:"Landing site"`
	Lat     float64
	Long    float64
}
type table5 struct {
	Module  string
	Landing string
	Lat     location2
	Long    location2
}

// В одной минуте 60 секунд («), в одном градусе 60 минут (‘)
func perevod(a location2) float64 { // это функция.
	b := 1.0
	if a.H == 'S' {
		b = -1.0
	}
	return b * (a.D + a.M/60 + a.S/3600)
}
