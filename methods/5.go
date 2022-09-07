/*
Покажчики та функції

Тут ми бачимо, що методи Abs і Scale переписані як функції.

Знову ж таки, спробуйте забрати * з рядка 27 (в оригіналі - 16). Бачите, чому змінилась поведінка? Що 
ще потрібно було змінити, щоб приклад скомпілювався?

(Якщо ви не впевнені, перейдіть на наступну сторінку).
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

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
