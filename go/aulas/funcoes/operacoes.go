package main

import (
	"fmt"
)

func operacoes1(a int, b int) (int, int) {
	return a+b, a - b
}

func operacoes2(a int, b int) (soma int,  diferenca int) {
	soma = a + b
	diferenca = a - b
	return
}

func main() {
	fmt.Println(operacoes1(1,2))
	fmt.Println(operacoes2(1,2))
}