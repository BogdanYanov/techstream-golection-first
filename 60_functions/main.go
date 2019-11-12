package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	xs := []float64{98, 23, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(sum(xs))
	x, y := f()
	fmt.Printf("x -> %d; y -> %d\n", x, y)
	fmt.Println(add(2, 3))
	fmt.Println(add(1, 2, 3))
	// Можем также передать слайсы
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(sl...))
	x1 := 0
	// Создание функции внутри другой функции
	increment := func() int {
		x1++
		return x1
	}
	fmt.Println(increment())
	fmt.Println(increment())
	// Функцию, использующую переменные, определенные вне этой функции, называют замыканием.

	/*
		Один из способов использования замыкания — функция, возвращающая другую функцию,
		которая при вызове генерирует некую последовательность чисел.
	*/
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(3 / 2)
	goFrom1To10()
	return
}

/*
	Функция начинается с ключевого слова func, за которым следует имя функции.
	Аргументы определяются как (имя тип, имя тип, ...). За параметром следует
	возвращаемый тип. (Аргументы и возвращаемое значение - сигнатура функции)
*/

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Можно также явно указать имя возвращаемого значения
func sum(xs []float64) (total float64) {
	total = 0.0
	for _, v := range xs {
		total += v
	}
	return
}

// Go может возвращать несколько значений из функций
func f() (int, int) {
	return 5, 7
}

// Можно указать переменное количество аргументов в функции
func add(args ...int) (total int) {
	total = 0
	for _, v := range args {
		total += v
	}
	return
}

// Panic, и ее восстановление с помощью recover
func goFrom1To10() {
	for i := 1; i <= 10; i++ {
		task(i)
	}
}

func task(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured:", r)
		}
	}()

	if i == 2 {
		panic("Got 2")
	}

	fmt.Println("I:", i)
}
