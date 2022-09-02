/*
Літерали слайсів

Літерал слайсу схожий на літерал масиву без довжини.

Це літерал масиву:

[3]bool{true, true, false}

І це створює той самий масив, що й вище, а потім будує фрагмент, який посилається на нього:

[]bool{true, true, false}
*/
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
