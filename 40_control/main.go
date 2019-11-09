package main

import "fmt"

func main() {
	a := true
	if a {
		fmt.Println("Hello, World!")
	}
	/*
		В Go неявного преобразования не работает, т.е. код вида:
			b := 1
			if b {
				fmt.Println("Something")
			}
		не сработает.
	 */

	b, c := 1, 22
	if (b == 1 && a) || c != 22 {
		fmt.Println("Yeah, it works")
	}

	user := map[string]string {
		"firstName": "Alexey",
		"lastName": "Popov",
	}

	if email, exist := user["email"]; exist {
		// Переменные email и exist будут доступны только в области видимости if`a
		fmt.Println("Email exist:", email)
	} else {
		fmt.Println("Email doesn`t exist")
	}

	// В Go нету конструкций while, do...while, но for обеспечивает тем же самым функционалом, что и while do...while
	for {
		fmt.Println("Infinity loop")
		break
	}

	sl := []int{1, 2, 3, 4 ,5, 6, 7, 8}
	value := 0
	idx := 0

	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value += sl[idx]
		idx++
		fmt.Println("While-style loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("C-style loop, i =", i, "sl[i] =", sl[i])
	}

	// Range - перебирает элементы в различных структурах данных.
	nums := []int{3, 4, 6}
	sum := 0
	// Range на каждой итерации будет возращать индекс (idx), и значение массива по этому индексу (num)
	for idx, num := range nums {
		sum += num
		fmt.Println("Index:", idx, "Number on index:", num, "Summary:", sum)
	}

	testAnswers := map[string]string {
		"a": "apple",
		"b": "banana",
		"c": "orange",
		"d": "strawberry",
	}

	for key, value := range testAnswers{
		fmt.Printf("Key %s -> value %s\n", key, value)
	}

	// Если в range передать просто строку, то мы получим код буквы и ее позицию в слове
	for idx, char := range "Bogdan" {
		fmt.Println(idx, char)
	}

	nums2 := []int{1, 2, 3, 4, 5, 8}
	for _, num := range nums2 {
		switch num {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fallthrough
		case 4:
			fmt.Println("Three of four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Default value")
		}
	}

	// Замена множественным if else
	switch {
	case testAnswers["a"] == "apple":
		fmt.Println("It`s apple")
	case testAnswers["b"] == "banana":
		fmt.Println("It`s banana")
	}

	// Выход из цикла, будучи внутри switch
loopMark:
	for key, value := range testAnswers {
		fmt.Println("Switch in loop; Key ->", key, "Value ->", value)
		switch {
		case key == "c" && value == "orange":
			fmt.Println("Switch break now.")
			break loopMark
		}
	}
	return
}
