/*
Базовы типи
В Go базовими типами є:

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // аліас для uint8

	rune // аліас для int32
	// представляє кодову точку Юнікоду

	float32 float64

	complex64 complex128

У прикладі показано змінні кількох типів, а також те, що оголошення змінних можуть бути «розбиті» на блоки, як у операторах імпорту.

Типи int, uint та uintptr зазвичай мають ширину 32 біти на 32-бітних системах, та 64 біти на 64-бітних системах.
Якщо вам потрібне ціле значення, ви повинні використовувати int, якщо у вас немає конкретної причини використовувати цілочисельний тип із розміром або без знаку.

*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
