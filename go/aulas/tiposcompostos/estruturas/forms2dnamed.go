// Exemplo de composição de estruturas
package main

import (
	"fmt"
)

// Campos nomeados
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int // Spokes são os aros da roda.
}

func main() {
	// Campos nomeados
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Printf("%#v\n", w)

	// Inicializando de forma literal com campos nomeados
	nw := Wheel {
		Circle: Circle {
			Center: Point {
				X: 8,
				Y: 8, // Vírgula no último campo é necessária!!
			},
			Radius: 5, // Vírgula no último campo é necessária!!
		},
		Spokes: 20,  // Vírgula no último campo é necessária!!
	}
	fmt.Printf("%#v\n", nw)
}
