/*
Defer stack
Відкладені виклики функцій поміщаються в стек. Коли функція повертається, її відкладені виклики
виконуються в порядку "останній прийшов - перший вийшов".

Щоб дізнатися більше про відкладені заяви, прочитайте цю публікацію в блозі.
https://go.dev/blog/defer-panic-and-recover
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
