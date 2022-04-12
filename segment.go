// программа работы с сегментами/срезами
package main

import "fmt"

func main() {
	numbers := []float64{19.7, 0, 25.5} // объявляем переменную для сегмента типа int
	words := []string{"a", "b", "c"}    // объявляем переменную для сегмента типа string
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	for i := 0; i < len(words); i++ {
		fmt.Println(i, words[i])
	}
	// сегмент на основе созданного массива
	myArray := [5]string{"a", "b", "c", "d", "e"} // создаём массив
	mySlice := myArray[0:3]                       // создаём сегмент (срез) на основе массива myArray начиная с позиции 0 и завершая ДО позиции 3, т.е.: ("a", "b", "c")
	fmt.Println(mySlice)                          // выводим элементы сегмента
	mySlice = myArray[:4]                         // создаём сегмент (срез) на основе массива myArray начиная с позиции 0 и завершая ДО позиции 4, т.е.: ("a", "b", "c", "d")
	fmt.Println(mySlice)                          // выводим элементы сегмента
	mySlice = myArray[2:]                         // создаём сегмент (срез) на основе массива myArray начиная с позиции 2 и завершая ДО позиции 5, т.е.: ("c", "d", "e")
	fmt.Println(mySlice)                          // выводим элементы сегмента
	// при изменении значения элемента базового массива, значение элемента сегмента на основе этого массива также меняется
	myArray[2] = "F"
	fmt.Println(mySlice)
	// и наоборот: если изменить значение элемента сегмента, то и в базовом массиве значение элемента изменится
	mySlice[2] = "J"
	fmt.Println(myArray)
	// применение функции Append - функция возвращает новый, расширенный сегмент со всеми элементами исходного сегмента и новыми элементами, добавленными в его конец
	mySlice2 := []string{"a", "b"} // создаем сегмент
	fmt.Println(mySlice2, len(mySlice2))
	mySlice2 = append(mySlice2, "c") // возвращаемое значение "append" присваивается той же переменной mySlice2. Элемент "c" присоединяется в конец сегмента
	fmt.Println(mySlice2, len(mySlice2))
	mySlice2 = append(mySlice2, "d", "e") // два элеменат присоединяются в конец сегмента
	fmt.Println(mySlice2, len(mySlice2))
	// ниже приведена прорамма, которая берёт сегмент массива и присоединяет элементы к сегменту.
	myArray3 := [5]string{"a", "b", "c", "d", "e"}
	mySlice3 := myArray3[1:3]
	mySlice3 = append(mySlice3, "x")
	mySlice3 = append(mySlice3, "y", "z")
	for _, letter := range mySlice3 {
		fmt.Println(letter) // результат вывода: b c x y z
	}
	mySlice3[1] = "I" // результат вывода: b I x y z
	fmt.Println(mySlice3)
}
