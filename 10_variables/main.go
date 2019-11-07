package main

func main() {
	// Целые числа
	var i int = 10 // Платформозависимый тип, 32 или 64 бита
	var autoInt = -10
	/* 	int8  - 1 байт в памяти
		int16 - 2 байта в памяти
		int32 - 4 байта в памяти
		int64 - 8 байт в памяти
	 */
	var bigInt int64 = 1<<32 - 1
	var unsignedInt uint = 100500	// Платформозависимый тип, 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1	// uint8, uint16, uint32, uint64
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// Числа с плавающей точкой
	var p float32 = 3.14 // float = float32, float64
	println("float:", p)

	// Булевые переменные
	var b = true
	println("bool variable:", b)

	// Строки
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	// Бинарные данные
	// 0 1 2 3 4 5 6 7 8 9 a b c d e f
	var rawBinary byte = '\x27'
	var rawBinary2 byte = '\x10'
	println("rawBinary:", rawBinary, rawBinary2)

	// Короткое объявление переменной
	meaningOfLive := 42
	println("Meaning of live is", meaningOfLive)
	/*
		Короткое объявление переменной работает только для
		новых переменных. поэтому для строчки кода:
			world := "Мир"
		будет выдана ошибка:
			"No new variables on left side of :="
	 */

	// Приведение типов
	// При приведении типа float в тип int число всегда будет округлятся до меньшего целого
	println("float to int conversion:", int(p))
	var p2 float32 = 3.51
	println("float to int conversion:", int(p2))
	var p3 float32 = 3.49
	println("float to int conversion:", int(p3))
	var p4 float64 = 3.89
	println("float to int conversion:", int(p4))
	var u1 uint = 17
	var i1 int = 23
	println(int(u1) + i1)
	/*
		При приведении типа int в тип string, мы будем получать
		символ который стоит в таблице Unicode, согласну числу,
		которое приводится.
	 */
	println("int to string conversion:", string(167))

	// Операции со строками
	s1 := "Bogdan"
	s2 := "Yanov"
	fullname := s1 + " " + s2
	println("Name length is:", fullname, "length:", len(fullname))
	/*
		При обращении к отдельному символу строки, нам вернется
		числовое значение данного символа. Для того чтобы получить
		не числовое значение а сам символ, необходимо просто числовое
		значение преобразовать к строке.
	 */
	println("Symbol on third position:", string(fullname[2]))

	escaping := `Hello\r\n
	World`
	println("As-is escaping:", escaping)

	// Значение по-умолчанию
	var defaultInt int			// 0
	var defaultFloat float32	// +0.000000e+000
	var defaultString string	// ""
	var defaultBool bool		// false
	println("Default values:", defaultInt, defaultFloat, defaultString, defaultBool)

	// Объявление нескольких переменных
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 = 12
		m1 = "string"
		m2 = 3.45
	)

	println(m0, m1, m2)
	return
}