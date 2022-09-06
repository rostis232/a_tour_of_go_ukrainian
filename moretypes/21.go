/*
Літерали мап продовження
Якщо тип верхнього рівня є просто ім'ям типу, то його можна не включати в елементи літералу.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}