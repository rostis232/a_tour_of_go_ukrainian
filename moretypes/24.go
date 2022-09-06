/*
Значення функцій
Функції - це теж значення. Їх можна передавати так само, як і інші значення.

Значення функції можуть використовуватися як аргументи функції та як значення, що повертаються.
*/

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
