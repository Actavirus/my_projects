// тема - Работа с массивами и срезами в Golang — append() и make()
package main

import (
	"fmt"
	"time"
)

func dump(label string, slice []string) {
	fmt.Printf("%v: длина %v, вместимость %v %v\n", label, len(slice), cap(slice), slice) // здесь: len - длина среза (это инициализированная память элементов, для превышения(добавления) нужно вручную использовать append.); cap - ёмкость среза (это выделенная память под элементы, при превышении размер автоматически увеличивается в два раза.)
}
func terraform(prefix string, worlds ...string) []string { // здесь (...) - для объявления вариативной функции, т.е. параметр worlds является срезом строк, что содержит ноль или более аргументов.
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
func main() {
	// Функция append в Go
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	dwarfs = append(dwarfs, "Оркус")                      // append - добавляет к срезу ещё один элемент
	fmt.Println(dwarfs)                                   // выводит: [Церера Плутон Хаумеа Макемаке Эрида Оркус]
	dwarfs = append(dwarfs, "Салация", "Квавар", "Седна") // append - добавляет к срезу несколько элементов
	fmt.Println(dwarfs)                                   // выводит: [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна]

	// Длина и вместимость среза в Golang
	dump("dwarfs", dwarfs)            // выводит: dwarfs: длина 9, вместимость 10 [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна]
	dump("dwarfs[1:2]", dwarfs[1:2])  // выводит: dwarfs[1:2]: длина 1, вместимость 9 [Плутон]
	dwarfs = append(dwarfs, "Pluto!") // добавялем ещё один элемент
	dump("dwarfs", dwarfs)            // выводит: dwarfs: длина 10, вместимость 10 [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна Pluto!]
	dwarfs = append(dwarfs, "Haski!") // добавялем ещё один элемент
	dump("dwarfs", dwarfs)            // выводит: dwarfs: длина 11, вместимость 20 [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна Pluto! Haski!]

	// Трех-индексный срез в Golang
	terrestrial := dwarfs[0:4:4]                // если уточняется третий индекс (длина 4, вместимость 4)...
	worlds := append(terrestrial, "Церера-2")   // ...добавление элемента "Церера-2" приводит к перемещению нового массива, оставляя массив dwarfs не измененным
	fmt.Println(terrestrial)                    // выводит: [Церера Плутон Хаумеа Макемаке]
	fmt.Println(worlds)                         // выводит: [Церера Плутон Хаумеа Макемаке Церера-2]
	fmt.Println(dwarfs)                         // выводит: [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна Pluto! Haski!]
	terrestrial2 := dwarfs[0:4]                 // если не уточняется третий индекс...
	worlds2 := append(terrestrial2, "Церера-2") // ...добавление элемента "Церера-2" не перемещает новый массив, а вместо этого переписывает "Эрида"
	fmt.Println(terrestrial2)                   // выводит: [Церера Плутон Хаумеа Макемаке]
	fmt.Println(worlds2)                        // выводит: [Церера Плутон Хаумеа Макемаке Церера-2]
	fmt.Println(dwarfs)                         // выводит: [Церера Плутон Хаумеа Макемаке Церера-2 Оркус Салация Квавар Седна Pluto! Haski!]

	// Предварительное выделение срезов через make в Go
	dwarfs2 := make([]string, 0, 10) // make уточняет длину (0) и вместимость (10) среза dwarfs2
	dwarfs2 = append(dwarfs2, "Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида")
	fmt.Printf("len %v cap %v: %v\n", len(dwarfs2), cap(dwarfs2), dwarfs2) // выводит: len 5 cap 10: [Церера Плутон Хаумеа Макемаке Эрида]
	dwarfs2 = append(dwarfs2, "Оркус", "Салация", "Квавар", "Седна", "Pluto!")
	fmt.Printf("len %v cap %v: %v\n", len(dwarfs2), cap(dwarfs2), dwarfs2) //выводит: len 10 cap 10: [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна Pluto!]
	dwarfs2 = append(dwarfs2, "Pluto!")
	fmt.Printf("len %v cap %v: %v\n", len(dwarfs2), cap(dwarfs2), dwarfs2) //выводит: len 11 cap 20: [Церера Плутон Хаумеа Макемаке Эрида Оркус Салация Квавар Седна Pluto! Pluto!]

	// Объявление вариативных функций в Golang (вариативность - переменное число аргументов)
	twoWorlds := terraform("Нью", "Аманда", "Джеймс", "Никита") // передаётся множество аргументов в функцию terraform
	fmt.Println(twoWorlds)                                      // выводит: [Нью Аманда Нью Джеймс Нью Никита]
	twoWorlds2 := []string{"Аманда", "Джеймс", "Никита"}
	newWorlds := terraform("Олд", twoWorlds2...) // для передачи среза просто расширьте срез через многоточие
	fmt.Println(newWorlds)                       // выводит: [Олд Аманда Олд Джеймс Олд Никита]

	// задача из учебника: Напишите программу, что использует цикл для продолжающегося добавления элементов в срез. Каждый раз при изменении вместимости среза выводится новое значение. - моё решение.
	slice_1 := []string{}
	fmt.Printf("len %v cap %v\n", len(slice_1), cap(slice_1)) // выводит: len 0 cap 0
	for i := 0; i <= 10; i++ {
		slice_1 = append(slice_1, dwarfs[i])
		fmt.Printf("len %v cap %v %v\n", len(slice_1), cap(slice_1), slice_1)
		time.Sleep(1 * time.Second)
	}

	// задача из учебника: Напишите программу, что использует цикл для продолжающегося добавления элементов в срез. Каждый раз при изменении вместимости среза выводится новое значение. - решение из учебника.
	s := []string{}
	lastCap := cap(s)

	for i := 0; i < 10000; i++ {
		s = append(s, "An element")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}
