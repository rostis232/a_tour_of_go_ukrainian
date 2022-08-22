/*
Кількість результатів

Функція може повертати будь-яку кількість результатів.

Функція swap повертає дві строки.
*/
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
