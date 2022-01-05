package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Há espaço para crescer. Estende a fatia. x e z compartilham o array subjacente.
		z = x[:zlen]
	} else {
		// Não há espaço suficiente. Aloca um novo array. x e z não compartilham o array subjacente.
		// Cresce para o dobro, para complexidade linear amortizada.
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Função embutida - copia elementos de uma fatia para outra do mesmo tipo. 
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(y), y)
		x = y 
	}
}