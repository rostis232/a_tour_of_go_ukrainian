/*
Методи та опосередкування покажчика

Порівнюючи попередні дві програми, можна помітити, що функції з покажчиком-аргументом обов'язково 
отримують покажчик:

	var v Vertex
	ScaleFunc(v, 5)  // Помилка компіляції!
	ScaleFunc(&v, 5) // OK

в той час, як методи з покажчиковими приймачами приймають в якості приймача або значення або покажчик 
при їх виклику:

	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK

Для інструкції v.Scale(5), незважаючи на те, що v є значенням, а не покажчиком, автоматично 
викликається метод з приймачем-покажчиком. Тобто, для зручності, Go інтерпретує оператор v.Scale(5) 
як (&v).Scale(5), оскільки метод Scale має приймач-покажчик.
*/
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
