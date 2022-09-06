/*
Мутація мап
Вставка або оновлення елементу в мапі m:

	m[key] = elem

Отримати елемент:

	elem = m[key]

Видалити елемент:

	delete(m, key)

Перевірити наявність ключа з допомогою присвоєння двох значень:

	elem, ok = m[key]

Якщо key є в m, то ok рівне true. Якщо ні, то ok рівне false.

Якщо ключ відсутній у карті, то elem є нульовим значенням для типу елементу карти.

Зауваження: Якщо elem або ok ще не були оголошені, можна використати коротку форму оголошення:

	elem, ok := m[key]
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
