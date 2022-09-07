/*
Методи продовження
Оголошувати метод можна і на неструктурних типах.

У даному прикладі розглядається числовий тип MyFloat з методом Abs.

Оголошувати метод з отримувачем, тип якого визначений в тому самому пакеті, що і метод, можна тільки 
тоді, коли цей отримувач визначений в тому самому пакеті, що і метод. Не можна оголошувати метод з 
отримувачем, тип якого визначений в іншому пакеті (до якого відносяться вбудовані типи типу int).
*/
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}