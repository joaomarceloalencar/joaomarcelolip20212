package main

import (
	"fmt"
	"strings"
	"bytes"
)

// basename: remove componentes de diretório e um .sufixo
// exemplos: a => a, a.go => a, a/b/c/.go => c, a/b.c.go => b.c
func basename1(s string) string {
	// Descarta a última '/' e tudo que estiver antes
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	} 

	// Preserva tudo que estiver antes do último '.''
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// Exemplo da biblioteca strings para tornar mais fácil
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 se "/" não for encontrado
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

// comma insere vírgulas em uma string que representa um número inteiro decimal não negativo
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]  
}

// intsToString é como fmt.Sprint(values), mas acrescenta vírgulas
func intsToString (values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func  main() {
	diretorio := "/home/jmhal/desenvolvimento/go/basename.go"
	fmt.Println(basename1(diretorio))
	fmt.Println(basename2(diretorio))

	numero := "100000"
	fmt.Println(comma(numero))

	numeros := []int{1, 2, 3}
	fmt.Println(intsToString(numero))
}