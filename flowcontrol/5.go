/*
If
В Go оператор if схожий на цикл for; вираз не обов'язково брати в дужки( ),  але фігурні дужки { } обов'язкові.
*/
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}