/*
Порядок оцінки оператора switch
Оператор switch оцінєю вирази зверху вниз, зупиняючись, коли вираз успішний.

(Наприклад,

	switch i {
	case 0:
	case f():
	}

не викликає f, якщо i==0.)

Примітка. Час на ігровому майданчику Go завжди починається з 2009-11-10 23:00:00 UTC, значення якого ми залишаємо на розсуд читача.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
