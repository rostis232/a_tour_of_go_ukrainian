/*
Цикл For

У Go є лише одна конструкція циклу, цикл for.

Основний цикл for складається з трьох компонентів, розділених крапкою з комою:
	оператор init: виконується перед першою ітерацією
	вираз умови: оцінюється перед кожною ітерацією
	оператор post: виконується в кінці кожної ітерації

Оператор init часто буде коротким оголошенням змінної, а змінні, оголошені там, видимі лише в області дії оператора for.

Цикл припинить ітерацію, коли булева умова оціниться як false.

Примітка. На відміну від інших мов, таких як C, Java або JavaScript, три компоненти оператора for не містять дужок,
а фігурні дужки { } завжди потрібні.
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
