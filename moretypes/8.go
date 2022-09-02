/*
слайси схожі на посилання на масиви

Фрагмент не зберігає жодних даних, він просто описує розділ основного масиву.

Зміна елементів фрагмента змінює відповідні елементи його основного масиву.

Інші фрагменти, які мають той самий базовий масив, побачать ці зміни.
*/

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
