package main
import (
	"fmt"
	"strings"
)
func main(){
		var pin string
		fmt.Print("Vvedite cifri pin-coda: ")
		fmt.Scan(&pin)
		fmt.Println(pin)
		
		number := strings.Split(pin, "")
		fmt.Println(number)

		for a := range number {
			
		}
}