// программа работы с массивами
package main

import "fmt"

func main() {
	/*	var numbers [3]int
		numbers[0] = 42
		numbers[2] = 108
		var letters = [3]string{"a", "b", "c"}
		fmt.Println(numbers[0])      // 42
		fmt.Println(numbers[1])      // 0
		fmt.Println(numbers[2])      // 108
		fmt.Println(letters[2])      // c
		fmt.Println(letters[0])      // a
		fmt.Println(letters[1])      // b
		fmt.Println(numbers)         // [42 0 108]
		fmt.Println(letters)         // [a b c]
		fmt.Printf("%#v\n", numbers) // [3]int{42, 0, 108}
		fmt.Printf("%#v\n", letters) // [3]string{"a", "b", "c"}
		index := 1
		fmt.Println(index, numbers[index]) // 1 0
		index = 2
		fmt.Println(index, letters[index]) // 2 c
		// перебираем элементы массивов в цикле for
		notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
		for i := 0; i <= 6; i++ {
			fmt.Println(i, notes[i]) // выводит все ноты с указанием индекса
		}
		fmt.Println(len(notes))              // выводим длину массива notes ("7")
		for i := 0; i <= len(notes)-1; i++ { // используем функцию len, полезно её применять когда мы не знаем точную длину массива
			fmt.Println(i, notes[i])
		}
		// применение специального цикла for...range. Используем для безопасного способа обработки элементов массива - ЛУЧШЕ ИСПОЛЬЗОВАТЬ ДАННЫЙ ВИД ЦИКЛА
		for index, note := range notes { //здесь index - переменная для хранения индексов, note - переменная для хранения строк, notes - обработка каждого значения в массиве
			fmt.Println(index, note)
		}
		for _, note := range notes { // здесь используется _ для того, чтобы не выводить индекс элемента массива
			fmt.Println(note)
		}
		for index, _ := range notes { // здесь используется _ для того, чтобы не выводить значение элемента массива
			fmt.Println(index)
		}
	*/
	// найдем среднее арифметическое значений элементов массива
	x := [5]float64{2, 3, 1, 4, 7}
	sum := 0.0 // для начала найдем сумму значений элементов массива
	for _, y := range x {
		sum += y
	}
	sampleCount := float64(len(x)) // здесь получаем длину массива (len) и переводим в тип float64
	fmt.Printf("Среднее арифметическое массива: %0.2f\n", sum/sampleCount)
	// программа, которая выводит индексы и значения всех элементов массива, лежащих в диапазоне от 10 до 20
	a := [6]int{3, 16, -2, 10, 23, 12}
	for b, c := range a {
		if c >= 10 && c <= 20 {
			fmt.Println(b, c)
		}
	}
}
