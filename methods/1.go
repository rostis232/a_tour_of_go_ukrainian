/*
Методи
У мові Go немає класів. Однак, ви можете визначати методи на типах.

Метод - це функція зі спеціальним аргументом-приймачем.

Отримувач з'являється у власному списку аргументів між ключовим словом func та іменем методу.

У даному прикладі метод Abs має отримувач типу Vertex з іменем v.
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
