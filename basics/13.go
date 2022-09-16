/*
Перетворення типів

Вираз T(v) перетворює значення v на тип T.

Деякі числові перетворення:

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
Або ще простіше:

	i := 42
	f := float64(i)
	u := uint(f)
	
На відміну від C, у Go призначення між елементами різного типу вимагає явного перетворення.
Спробуйте видалити перетворення float64 або uint у прикладі та подивіться, що станеться.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
