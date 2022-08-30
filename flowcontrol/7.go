/*
If та else

Змінні, оголошені в короткому операторі if, також доступні в будь-якому з блоків else.

(Обидва виклики pow повертають свої результати до початку виклику fmt.Println у main.)
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
