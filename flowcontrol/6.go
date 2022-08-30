/*
If з коротким виразом

Як і for, оператор if може починатися з короткого виразу для виконання перед умовою.

Змінні, оголошені оператором, знаходяться в області видимості лише до кінця if.

(Спробуйте використати v в останньому операторі return.)
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
