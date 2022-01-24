package main

import (
	"fmt"
)

// Vamos deixar o nome minúsculo mesmo, pois só vamos usá-lo neste arquivo
type tree struct {
	value			int	
	left, right 	*tree // aqui não poderia ter apenas tree, tem que ser um ponteiro
}

// Sort ordena valores in-place.
func Sort(values []int) {
	// Inicialmente, root terá valor nil, padrão para ponteiros não inicializados
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	// [:0] indica a primeira posição da fatia.
	appendValues(values[:0], root)
}

// appendValues concatena os elementos de t a values na ordem 
// e devolve a fatia resultante
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalente a devolver &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func main() {
	numbers := []int{19, 84, 22, 01, 20, 14, 100, 2}
	fmt.Println(numbers)
	Sort(numbers)
	fmt.Println(numbers)
}