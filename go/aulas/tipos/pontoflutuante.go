package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e eˆx = %8.3f\n", x, math.Exp(float64(x)))
	}
}