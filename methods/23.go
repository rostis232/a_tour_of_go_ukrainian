/*
Вправа: rot13Reader

Поширеним патерном є io.Reader, який обгортає інший io.Reader, модифікуючи потік деяким чином.

Наприклад, функція gzip.NewReader отримує io.Reader (потік стиснених даних) і повертає *gzip.Reader, який 
також реалізує io.Reader (потік розпакованих даних).

Реалізувати rot13Reader, який реалізує io.Reader та читає з io.Reader, модифікуючи потік шляхом застосування 
шифру підстановки rot13 до всіх алфавітних символів.

Вам надається тип rot13Reader. Зробіть його io.Reader, реалізувавши його метод Read.
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
