/*
Імена, що експортуються

В Go імена експортуються, якщо починаються з великої літери. Наприклад Pizza це ім'я, що експортується, як і Pi,
яке експортується із пакету math.

pizza та pi не починаються з великої літери, тому вони не експортуються.

При імпортуванні пакету, можна звертатись тільки до імен, які експортуються. Будь-яке ім'я, що не експортується, є
недоступним ззовні пактеу.

Запустіть код. Зверніть увагу на повідомлення помилки.

Виправте помилку, переіменувавши math.pi в math.Pi та спробуйте ще раз.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}
