package main

import (
	"fmt"
)

var months [13]string = [...]string{1: "January", 
									2: "February",
									3: "March",
									4: "April",
									5: "May",
									6: "June",
									7: "July",
									8: "August",
									9: "September",
									10: "October", 
									11: "November",
									12: "December"}

func main() {
	// Imprimindo a string vazia
	fmt.Println(months[0])

	// Definindo duas slices - segundo trimeste e verão no hemisfério norte.
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	// Descobrindo a intersecção de duas fatias
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s aparece em ambos.\n", s)
			}
		}
	}

	// A linha a seguir causa erro!
	fmt.Println(summer[:7])

	// As linhas seguintes referenciam além do tamanho da fatia, mas não da capacidade.
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

}