/*
Readers (Рідери)
Пакет io визначає інтерфейс io.Reader, який представляє читання кінця потоку даних.

Стандартна бібліотека Go містить багато реалізацій цього інтерфейсу, включаючи файли, мережеві з'єднання, 
компресори, шифри та інші.

Інтерфейс io.Reader має метод Read:

	func (T) Read(b []byte) (n int, err помилка)

Read заповнює заданий байтовий зріз даними і повертає кількість заповнених байт та значення помилки. При 
завершенні потоку повертається помилка io.EOF.

У прикладі створюється strings.Reader і споживає на виході 8 байт за один раз.
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
