/*
Завдання: Помилки
Скопіюйте функцію Sqrt з попередньої вправи (https://go.dev/tour/flowcontrol/8) і змініть її так, щоб вона 
повертала значення помилки.

Функція Sqrt повинна повертати ненульове значення помилки, коли задається від'ємне число, оскільки вона не 
підтримує комплексні числа.

Створіть новий тип

type ErrNegativeSqrt float64

і зробіть його помилковим, присвоївши йому функцію

func (e ErrNegativeSqrt) Error() string

метод таким чином, щоб ErrNegativeSqrt(-2).Error() повертав "cannot Sqrt negative number: -2".

Примітка: Виклик fmt.Sprint(e) всередині методу Error відправить програму в нескінченний цикл. Цього можна 
уникнути, якщо спочатку перетворити e: fmt.Sprint(float64(e)). Чому?

Змініть функцію Sqrt, щоб вона повертала значення ErrNegativeSqrt при отриманні від'ємного числа.
*/

package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
