/*
Коротке оголошення змінних

Усередині функції замість оголошення var із неявним типом можна використовувати короткий оператор присвоєння :=.

Поза функцією кожен оператор починається з ключового слова (var, func тощо), тому конструкція := недоступна.
*/

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}