package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println("zero(x):", x)
	zeroPtr(&x) // &x вернет *int (указатель на int) потому что x имеет тип int
	fmt.Println("zeroPtr(x):", x)

	// Другой способ получить указатель — использовать встроенную функцию new:
	xPtr := new(int)
	one(xPtr)
	fmt.Println("*xPtr =", *xPtr)
	/*
		Функция new принимает аргументом тип, выделяет для него память и возвращает указатель на эту память.
	 */
	x1, x2 := 1, 2
	fmt.Println("x1 :", x1, "x2:", x2)
	swap(&x1, &x2)
	fmt.Println("x1 :", x1, "x2:", x2)
	return
}

func zero(x int) {
	x = 0
}

/*
	Указатели указывают (прошу прощения за тавтологию) на участок в памяти,
	где хранится значение. Используя указатель (*int) в функции zero,
	мы можем изменить значение оригинальной переменной.
 */
func zeroPtr(xPtr *int) {
	*xPtr = 0 // Храним int 0 в памяти, на которую указывает xPtr
}

func one(xPtr *int) {
	*xPtr = 1
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
