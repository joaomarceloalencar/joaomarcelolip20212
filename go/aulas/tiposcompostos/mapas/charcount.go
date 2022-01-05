// Charcount conta caracteres Unicode
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// Lembrem que caracteres Unicode são do tipo rune
	counts := make(map[rune]int)

	// Contagem de tamanho das codificações UTF-8 (quantas runas de 1, 2, 3 ou 5 bytes etc)
	var utflen [utf8.UTFMax + 1]int

	// Contagem de caracteres inválidos em UTF-8
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		// Retorna runa, número de bytes que ela ocupa e possível erro
		r, n, err := in.ReadRune() 

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}