/*
Вправа: Замикання Фібоначчі
Давайте трохи пограємось з функціями.

Реалізуйте функцію fibonacci, яка повертає функцію (замикання), що повертає послідовні числа 
Фібоначчі (0, 1, 1, 2, 3, 5, ...).
https://en.wikipedia.org/wiki/Fibonacci_number
*/
package main

import "fmt"

// fibonacci це функція, яка повертає
// функцію, яка повертає int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
