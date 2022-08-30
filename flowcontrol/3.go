/*
Для Go цикл For також виконує роль циклу while

Ви можете відкинути крапки з комами: цикл for в Go виконує роль циклу while в C.
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
