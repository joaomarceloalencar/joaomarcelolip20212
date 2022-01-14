package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool{
	if len(x) != len(y) {
		return false
	}

	// k é de key, chave
	// xv é de x-value, valor de x
	// yv é de y-value, valor de y
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false 
		}
	}
	return true
}

func main() {
	// Alocando o mapa com make
	ages0 := make(map[string]int)
	ages0["alice"] = 31
	ages0["charlie"] = 34
	ages0["bob"] = 29
	ages0["alex"] = 18

	// Alocação alternativa, usando um mapa literal
	ages1 := map[string]int {
		"alice": 31,
		"charlie": 34,
	}
	
	// Misturando alternativas
	ages2 := map[string]int {
		"alice": 31,
	}
	ages2["charlie"] = 34

	// Ordenando as chaves
	var names []string
	for name := range ages0 {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages0[name])
	}

	fmt.Println(equal(ages0, ages1))
	fmt.Println(equal(ages0, ages2))
	fmt.Println(equal(ages1, ages2))
}