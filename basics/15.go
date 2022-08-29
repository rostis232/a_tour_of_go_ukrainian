/*
Константи
Константи оголошуються як змінні, але з ключовим словом const.

Константи можуть бути символьними, рядковими, логічними або числовими значеннями.

Константи не можуть бути оголошені за допомогою синтаксису :=.
*/

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
