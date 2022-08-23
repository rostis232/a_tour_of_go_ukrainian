/*
Змінні

Оператор var оголошує список змінних; як і списках агрументів функцій, тип вказується в кінці.

Оператор var може бути на рівні пакета або функції. Обидва є в цьому прикладі.
*/
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
