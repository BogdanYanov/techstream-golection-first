package main

import "fmt"

func main() {
	// Карта (также известна как ассоциативный массив или словарь) — это неупорядоченная коллекция пар вида ключ-значение.
	var keys map[string]int
	fmt.Println("Uninitialized map:", keys)
	/*
		Если мы выполним строчки кода:
			keys["Enabled"] = 1
			keys["Disabled"] = 0
		мы получим ошибку:
			"panic: assignment to entry in nil map"
		Проблема нашей программы в том, что карта должна быть инициализирована перед тем, как будет использована.
		Решение:
	 */
	keys2 := make(map[string]int) // или keys2 := map[string]int{}
	keys2["Enabled"] = 1
	keys2["Disabled"] = 0
	fmt.Println("Initialized map:", keys2)

	// Если обратиться к несуществующему ключу, то вернется значение по-умолчанию
	neutralVal := keys2["Neutral"]
	fmt.Println("neutralVal =", neutralVal)

	// Проверка на то, что значение по даному ключу есть, а не задано по-умолчанию
	myVal, ok := keys2["myVal"]
	fmt.Println("myVal =", myVal, "Is exists?", ok)

	// Можно получить только получение признака существования
	_, isExist := keys2["myVal2"]
	fmt.Println("myVal2 exist?", isExist)

	// Можно создавать многомерную карту
	cars := map[string] map[string]int{
		"BMW": map[string]int {
			"Max speed": 220,
			"Fuel consumption": 8,
		},
		"VAZ": map[string]int {
			"Max speed": 140,
			"Fuel consumption": 12,
		},
	}
	fmt.Println(cars)

	// Удаление значения
	_, isExist = cars["VAZ"]
	fmt.Println("VAZ exists in cars?", isExist)
	delete(cars, "VAZ")
	_, isExist = cars["VAZ"]
	fmt.Println("VAZ exists in cars?", isExist)
	fmt.Println(cars)

	// Проверка присвоения
	cars2 := cars
	cars2["VAZ"] = map[string]int{
		"Max speed": 140,
		"Fuel consumption": 12,
	}
	fmt.Println(cars, cars2)
	// Карты присваиваются по ссылкам, а значит изменения в одной из карт будут отражены и в другой карте.
	return


}
