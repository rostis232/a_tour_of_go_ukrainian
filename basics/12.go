/*
Нульові значення

Змінні оголошенні без початкових значень отримують нульові значення.

Нульовими значеннями є:

	0 для числових типів
	false для Булеанівського типів
	"" (пуста строка) для строк.
*/
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
