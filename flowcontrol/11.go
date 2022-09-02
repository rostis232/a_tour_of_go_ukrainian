/*
Switch без умов
Switch без умови те саме, що switch true.

Ця конструкція може бути простим способом написання довгих ланцюжків if-then-else.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
