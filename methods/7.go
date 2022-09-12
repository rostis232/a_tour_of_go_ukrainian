/*
Методи та вказівники непрямої дії (2)

У зворотному напрямку відбувається те саме.

Функції, що отримують аргумент-значення, повинні приймати значення саме цього типу:

	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // Помилка компіляції!

в той час як методи з отримувачами-значеннями отримують в якості отримувача або значення або покажчик при їх виклику:

	var v Vertex
	fmt.Println(v.Abs()) // ОК
	p := &v
	fmt.Println(p.Abs()) // OK

У цьому випадку виклик методу p.Abs() інтерпретується як (*p).Abs().
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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
