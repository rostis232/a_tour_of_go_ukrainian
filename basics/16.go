/*
Числові константи
Числові константи є значеннями високої точності.

Нетипова константа приймає тип, необхідний її контексту.

Спробуйте також надрукувати needInt(Big).

(Int може зберігати максимум 64-бітове ціле число, а іноді й менше.)

*/

package main

import "fmt"

const (
	// Створіть величезне число, змістивши 1 біт вліво на 100 місць.
	// Іншими словами, двійкове число, яке дорівнює 1 і 100 нулям.
	Big = 1 << 100
	// Знову зсуньте праворуч на 99 позицій, щоб отримати 1<<1 або 2..
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
