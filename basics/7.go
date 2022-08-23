/*
Повернення іменованих значень

Значення, що повертаються можуть бути іменованими.
Якщо так, вони розглядаються як змінні, визначені у верхній частині функції.

Ці імена слід використовувати для документування значення повернених значень.

Оператор return без аргументів повертає іменовані значення. Це відоме як «голе» повернення.

Голе повернення має використовуватись тільки в коротких функціях, як у цьому прикладі.
У більших функціях може постраждати читабельність коду.
*/
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
