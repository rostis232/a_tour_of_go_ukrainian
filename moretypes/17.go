/*
Range продовження
Ви можете пропустити індекс або значення, призначивши _ замість змінної.

	for i, _ := range pow
	for _, value := range pow

Якщо вам потрібен лише індекс, ви можете опустити другу змінну.

	for i := range pow

*/

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
