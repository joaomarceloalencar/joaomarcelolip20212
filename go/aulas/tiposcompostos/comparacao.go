package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a 
	// Erro!!! Fatia só pode ser comparada com nil.
	fmt.Println(a == b)
}