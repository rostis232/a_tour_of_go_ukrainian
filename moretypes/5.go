/*
Структурні літерали

Літерал структури позначає щойно виділене значення структури шляхом переліку значень її полів.

Ви можете перерахувати лише підмножину полів, використовуючи синтаксис Name:. (І порядок іменованих полів
 не має значення.)

Спеціальний префікс & повертає вказівник на значення структури.
*/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

