/*
Defer
Інструкція defer відкладає виконання функції до повернення навколишньої функції.

Аргументи відкладеного виклику оцінюються негайно, але виклик функції не виконується, 
доки не повернеться навколишня функція.
*/
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
