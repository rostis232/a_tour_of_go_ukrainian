/*
Методи - це функції

Запам'ятайте: метод - це просто функція з аргументом-приймачем.

Ось Abs записаний як звичайна функція без зміни функціональності.
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
