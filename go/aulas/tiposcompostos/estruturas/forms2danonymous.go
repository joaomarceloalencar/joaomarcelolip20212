// Exemplo de composição de estruturas
package main

import (
	"fmt"
)

// Campos anônimos
type Point struct {
	X, Y int
}

type Circle struct {
	Point // Apenas o nome do tipo
	Radius int
}

type Wheel struct {
	Circle // Apenas o nome do tipo
	Spokes int 
}


func main() {
	// Campos anônimos
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Printf("%#v\n", w)

	// Inicializando de forma literal com campos anônimos
	nw := Wheel {
		Circle: Circle {
			Point: Point {
				X: 8,
				Y: 8, // Vírgula no último campo é necessária!!
			},
			Radius: 5, // Vírgula no último campo é necessária!!
		},
		Spokes: 20,  // Vírgula no último campo é necessária!!
	}
	fmt.Printf("%#v\n", nw)
}
