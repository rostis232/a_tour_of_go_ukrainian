/*
Завдання: Рідери
Реалізувати тип Reader, який видає нескінченний потік ASCII-символу 'A'.
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Додати метод Read([]byte) (int, error) до MyReader.

func main() {
	reader.Validate(MyReader{})
}
