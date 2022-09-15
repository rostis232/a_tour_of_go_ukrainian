/*
Помилки

Програми на Go виражають стан помилок значеннями помилок.

Тип error є вбудованим інтерфейсом, подібним до fmt.Stringer:

type error interface {
    Error() рядок
}

(Як і у випадку з fmt.Stringer, пакет fmt шукає інтерфейс помилки при виведенні значень).

Функції часто повертають значення помилки, і викликаючий код повинен обробляти помилки, перевіряючи, чи 
помилка дорівнює nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("не вдалося перетворити число: %v\n", err)
    return
}
fmt.Println("Перетворено ціле число:", i)
Помилка, рівна nil, означає успіх, помилка, відмінна від nil, означає відмову
*/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
