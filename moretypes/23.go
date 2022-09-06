/*
Вправа: Мапи
Реалізувати функцію WordCount. Вона повинна повертати карту підрахунку кожного "слова" у рядку s. 
Функція wc.Test виконує набір тестів для заданої функції і виводить успіх або невдачу.

Вам може бути корисною функція strings.Fields. https://go.dev/pkg/strings/#Fields
*/
package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
