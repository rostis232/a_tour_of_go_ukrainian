/*
Додавання до слайсу
Зазвичай до слайсу додаються нові елементи, тому Go надає вбудовану функцію додавання. Документація 
(https://go.dev/pkg/builtin/#append) вбудованого пакета описує append.

	func append(s []T, vs ...T) []T

Перший параметр s функції append — це зріз типу T, а решта — це значення T, які потрібно додати до зрізу.

Отримане значення функції append є слайс, що містить усі елементи вихідного слайса плюс додані значення.

Якщо резервний масив s надто малий, щоб вмістити всі задані значення, буде виділено більший масив. 
Повернений фрагмент вказуватиме на щойно виділений масив.

(Щоб дізнатися більше про зрізи, прочитайте статтю про зрізи: використання та внутрішні функції.)
https://go.dev/blog/go-slices-usage-and-internals
*/
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
