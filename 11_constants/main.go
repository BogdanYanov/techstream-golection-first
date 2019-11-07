package main

const someInt = 1
const typedInt int32 = 17
const fullname = "Bogdan"

/*
	Можем использовать тот же синтаксис объявления констант,
	что и при объявлении переменных
 */

const (
	flagKey1 = 1
	flagKey2 = 2
)

/*
	Iota - является основным инструментом для перечисляемых констант.
	Предварительно объявленный iota идентификатор сбрасывается в 0 каждый раз,
	когда слово const появляется в исходном коде, и увеличивается после каждой спецификации const
 */

const (
	one = iota + 1
	two
	_
	four
)

const (
	five = iota + 5
	six
	seven
)

const (
	_ = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	flagEditText = 1 << iota
	flagReadText
	flagSetLike
)

var userFlags = (1<<2) + (1<<1)

func main() {
	pi := 3.14
	// Тип константы может быть определен во время компиляции
	println(pi + someInt)
	/*
		Но если мы используем типизированную константу, то
		при строчке кода:
			println(pi + typedInt)
		мы получим ошибку:
			"invalid operation: pi + typedInt (mismatched types float64 and int32)
	 */
	// Решение:
	println(pi + float64(typedInt))

	println(one, two, four)
	println(five, six, seven)
	println(KB, MB, GB, TB, PB, EB)
	println((userFlags & flagEditText) == flagEditText)
	println((userFlags & flagSetLike) == flagSetLike)
	println((userFlags & flagReadText) == flagReadText)
	return
}
