/*
Масиви

Тип [n]T — це масив із n значень типу T.

Вираз
	var a [10]int
оголошує змінну a як масив із десяти цілих чисел.

Довжина масиву є частиною його типу, тому розмір масиву не можна змінити. Це здається обмеженим, 
але не хвилюйтеся; Go забезпечує зручний спосіб роботи з масивами.
*/

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}