/*
Порожній інтерфейс

Тип інтерфейсу, що визначає нульові методи, називається пустим інтерфейсом:

	interface{}

Пустий інтерфейс може містити значення будь-якого типу. (Кожен тип реалізує як мінімум нуль методів).

Порожні інтерфейси використовуються кодом, який працює зі значеннями невідомого типу. Наприклад, функція 
fmt.Print отримує будь-яку кількість аргументів типу interface{}.
*/
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
