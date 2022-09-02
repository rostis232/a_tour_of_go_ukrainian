/*
Поінтери (вказівники)
Go має поінтери. Поінтер містить адресу пам'яті значення.

Тип *T є поінтером на значення T. Його нульове значення nil.

	var p *int

Оператор & генерує вказівник на свій операнд.

	i := 42
	p = &i

Оператор * позначає базове значення вказівника.

	fmt.Println(*p) // читання i через покажчик p
	*p = 21 // встановити i через покажчик p

Це відомо як "розіменування".

На відміну від C, Go не має арифметики вказівників.
*/
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
