/*
Карти
Карта відображає ключі на значення.

Нульовим значенням карти є нуль. Нульова карта не має ключів і не може додавати ключі.

Функція make повертає карту заданого типу, ініціалізовану та готову до використання.
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}