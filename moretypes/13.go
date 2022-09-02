/*
Створення слайсу за допомогою make

Слайси можна створювати за допомогою вбудованої функції make; таким чином ви створюєте масиви динамічного 
розміру.

Функція make виділяє обнулений масив і повертає слайс, який посилається на цей масив:

	a := make([]int, 5) // len(a)=5

Щоб указати ємність, передайте третій аргумент, щоб зробити:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:] // len(b)=4, cap(b)=4
*/
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
