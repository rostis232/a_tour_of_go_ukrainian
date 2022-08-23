/*
Змінні з ініціалізаторами

Оголошення var може містити ініціалізатори, по одному для кожної змінної.

Якщо ініціалізатор присутній, тип можна опустити; змінна прийме тип ініціалізатора.
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
