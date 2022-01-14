package main

import (
	"fmt"
)

// Um array de 3 inteiros
var a [3]int = [...]int{1, 2, 3}

func main() {
	// Exibe o primeiro elemento
 	fmt.Println(a[0])

	// Exibe o último elemento, a[2]
	fmt.Println(a[len(a) - 1])

	// Exibe os índices e os elementos
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Exibe apenas os elementos
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

    r := [...]int{9:-1}
	fmt.Println(r)

}