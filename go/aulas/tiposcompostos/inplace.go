package main

import "fmt"

// noempty devolve uma fatia que armazena apenas as strings não vazias
// O array subjacente é modificado durante a chamada.
func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s 
			i++
		}
	}

	return strings[:i]
}

// Versão utilizando append
func noempty2(strings []string) []string {
	out := strings[:0] // fatia de tamanho zero do original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out 
}

// Remover um elemento do meio da fatia
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice) - 1]
}

func main() {

	// Exemplo de noempty - observe que o array é alterado.
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)

	// Exemplo de remove
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s,2))

}