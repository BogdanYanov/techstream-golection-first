package main

import "fmt"

func main() {
	/*
		Срезы (slice) представляют последовательность элементов одного типа переменной длины.
		В отличие от массивов длина в срезах не иксирована и динамически может меняться, то есть
		можно добавлять новые элементы или удалять уже существующие.
	 */
	var sl1 []int
	fmt.Println("Empty slice:", sl1)
	sl2 := []string{"Monkey", "Cat", "Dog"}
	fmt.Println("Not empty slice:", sl2)

	// Обращение к элементу, как и в массиве, по индексу:
	fmt.Println("Is it a dog?", sl2[2])

	// Добавление элемента в слайс происходит таким образом:
	sl2 = append(sl2, "Elephant")
	fmt.Println("Slice after add:", sl2)

	// Можно добавить несколько элементов указав их через запятую:
	sl1 = append(sl1, 1, 3, 5 ,7)
	fmt.Println("Slice 1:", sl1)

	/*
		Как устроен слайс?
			struct {
				len (length - длина)
				cap (capacity - вместимость, или же длина внутреннего массива)
				*arr (array - ссылка на массив, который содержит наши элементы)
			}
		Если при добавлении элементов мы не вмещаемся в исходный массив,
		т.е. вместимость нам не позволяет это сделать, то создается новый
		массив, с увелечением capacity в два раза.
	 */
	sl3 := []int{100}
	fmt.Println("Slice 3:", sl3, "Length:", len(sl3), "Capacity:", cap(sl3))
	sl3 = append(sl3, 101)
	fmt.Println("Slice 3:", sl3, "Length:", len(sl3), "Capacity:", cap(sl3))
	sl3 = append(sl3, 102, 103, 104)
	fmt.Println("Slice 3:", sl3, "Length:", len(sl3), "Capacity:", cap(sl3))
	sl3 = append(sl3, 105, 106)
	fmt.Println("Slice 3:", sl3, "Length:", len(sl3), "Capacity:", cap(sl3))

	// Добавление слайса в слайс
	pricesAll := []float32{15.90, 22.10, 13.75}
	pricesNew := []float32{45.23, 13.22, 10.99}
	pricesAll = append(pricesAll, pricesNew...)
	fmt.Println("All prices:", pricesAll)

	// Для того, чтобы создать слайс сразу нужной длины:
	sl4 := make([]int, 10)
	fmt.Println("Slice 4:", sl4, "Length:", len(sl4), "Capacity:", cap(sl4))
	/*
		Но мы можем столкнуться с тем, что при добавлении элемента
		в этот слайс, у нас увеличится capacity, а следовательно
		массив будет занимать больше места в памяти.
		Решением будет создать массив нужной длины и нужной capacity.
	 */
	sl4 = append(sl4, 1)
	fmt.Println("Slice 4:", sl4, "Length:", len(sl4), "Capacity:", cap(sl4))
	sl5 := make([]int, 10, 15)
	fmt.Println("Slice 5:", sl5, "Length:", len(sl5), "Capacity:", cap(sl5))
	sl5 = append(sl5, 1)
	fmt.Println("Slice 5:", sl5, "Length:", len(sl5), "Capacity:", cap(sl5))

	// Т.к. внутри слайса ссылка на массив, то при присвоении одного слайса другому она просто копируется
	sl6 := sl2
	sl6[0] = "Mouse"
	fmt.Println("Slice 2:", sl2, "Length:", len(sl2), "Capacity:", cap(sl2))
	fmt.Println("Slice 6:", sl6, "Length:", len(sl6), "Capacity:", cap(sl6))

	// Но при добавлении элемента, будет уже ссылка на другую область памяти, поэтому они становятся независимыми
	sl6 = append(sl6, "Monkey")
	fmt.Println("Slice 2:", sl2, "Length:", len(sl2), "Capacity:", cap(sl2))
	fmt.Println("Slice 6:", sl6, "Length:", len(sl6), "Capacity:", cap(sl6))

	// Неправильная попытка скопировать слайс
	var sl7 []string
	copy(sl7, sl6)
	fmt.Println(sl7)

	// Правильная попытка скопировать слайс
	sl8 := make([]string, len(sl6), len(sl6))
	copy(sl8, sl6)
	fmt.Println(sl8)

	// Можно обращаться к частям слайса
	fmt.Println("С первого по второй:", sl8[:2], "Со второго по пятый:", sl8[1:], "Со третьего по четвертый:", sl8[2:4])

	// Создание слайса из массива
	arr := [...]int{1, 2, 3}
	slice := arr[:]
	arr[2] = 5
	fmt.Println("Slice: ", slice, "Array:", arr)
	return
}
