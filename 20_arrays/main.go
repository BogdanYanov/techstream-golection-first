package main

import "fmt"

const size uint = 3

func main()  {
	// При создании массива, он инициализируется значениями по-умолчанию
	var arr1 [3]int
	fmt.Println("Array of integers:", arr1, "length:", len(arr1))

	// Для определения размера, можно использовать типизированную беззнаковую константу
	var arr2 [2 * size]bool
	fmt.Println("Array of bool:", arr2, "length:", len(arr2))

	var arr3 [4]string
	fmt.Println("Array of strings:", arr3, "length:", len(arr3))

	// Также можно использовать [...] для того чтобы использовалась длина при инициализации
	arr4 := [...]int{1, 2, 3, 4}
	fmt.Println("Array of integers before changes:", arr4, "length:", len(arr4))

	// Изменение значения элемента массива
	arr4[3] = 5
	fmt.Println("Array of integers after changes:", arr4, "length:", len(arr4))

	// Многомерный массив
	var matrix [3][3]int
	matrix[0][0], matrix[1][1], matrix[2][2] = 1, 1, 1
	fmt.Println("Matrix:", matrix, "length:", len(matrix))

	// Можно передавать меньше значений чем длина массива, но не больше
	arr5 := [5]int{1,2}
	fmt.Println("Array 5:", arr5)

	/*
		Индексы в массиве фактически выступают в качестве ключей,
		по которым можно обратиться к соответствующему значению.
		И в принципе мы можем явным образом указать, какому ключу
		какое значение будет соответствовать. При этому числовые ключи
		необязательно располагать в порядке возрастания:
	 */
	colors := [3]string{0: "red", 2: "blue", 1: "green"}
	fmt.Println("Array of colors:", colors)
	return
}
