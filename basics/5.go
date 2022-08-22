/*
Функції продовження

Коли два чи більше параметри підряд, які приймає функція, мають один тип, можна вказати тип тільки останнього параметру.

В цьому прикладі ми скоротили:

	x int, y int

до:

	x, y int
*/

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
