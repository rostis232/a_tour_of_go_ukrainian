/*
Інтерфейси
Тип інтерфейсу визначається як набір сигнатур методів.

Значення типу інтерфейсу може мати будь-яке значення, що реалізує ці методи.

Примітка: У коді прикладу у рядку 32 (в оригіналі 22) допущено помилку. Vertex (тип-значення) не реалізує метод Abser, тому що 
метод Abs визначений тільки на *Vertex (тип-покажчик).
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat імплементує Abser
	a = &v // *Vertex імплементує Abser

	// В наступному рядку, v є Vertex (не *Vertex)
	// та НЕ імплементує Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
