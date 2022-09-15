/*
Вправа: Стрінгери

Зробіть так, щоб тип IPAddr реалізував fmt.Stringer для виведення адреси у вигляді чотриьох чисел через
крапки.

Наприклад, IPAddr{1, 2, 3, 4} повинен виводити "1.2.3.4".
*/
package main

import "fmt"

type IPAddr [4]byte

// Додайте метод "String() string" до IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
