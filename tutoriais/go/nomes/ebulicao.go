package main

import "fmt"

// Visível no pacote
const ebulicaoC = 100.0

func main () {
  // Visível em main
  var c = ebulicaoC
  var f = (c * 9)/5 + 32

  fmt.Printf("Ponto de ebulição = %g F ou %g C\n", f, c)
}