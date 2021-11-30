package main

import "fmt"

var numero, Numero = 10, 20
var v int

func fazNada() (int, int) {
	numero1 := 5
	numero2 := 6
	return numero1, numero2
}

func soma(a int, b int) int {
	var temporario int
	temporario = a + b
	return temporario
}

func main() {
    n1, _ := fazNada()
	v = 1.5
	fmt.Printf("%d\n", soma(n1, numero))
    fmt.Printf("%g\n", v)
}