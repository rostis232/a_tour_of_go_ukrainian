/*
Перемикачі типів

Перемикач типів - це конструкція, яка допускає декілька послідовних тверджень типу.

Перемикач типів схожий на звичайний оператор switch, але в перемикачі типів в регістрах вказуються типи
(а не значення), і ці значення порівнюються з типом значення, що зберігається в даному значенні інтерфейсу.

	switch v := i.(type) {
	case T:
		// тут v має тип T
	case S
		// тут v має тип S
	default:
		// немає співпадіння; тут v має той самий тип що й i
	}

Оголошення в switch типу має такий самий синтаксис, як і твердження типу i.(T), але конкретний тип T 
замінюється ключовим словом type.

Цей оператор switch перевіряє, чи інтерфейсне значення i має значення типу T або S. У кожному з 
випадків T і S змінна v буде мати тип T або S відповідно і зберігати значення, яке має i. У випадку 
за замовчуванням (коли немає співпадіння) змінна v має той самий інтерфейсний тип і значення, що й i.
*/
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
