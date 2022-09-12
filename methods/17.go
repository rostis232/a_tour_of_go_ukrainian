/*
Стрінгери
Одним з найбільш розповсюджених інтерфейсів є інтерфейс Stringer, що визначається пакетом fmt.

type Stringer interface {
    String() string
}
Stringer - це тип, який може описувати себе як рядок. Пакет fmt (і багато інших) шукає цей інтерфейс для 
виведення значень.
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
