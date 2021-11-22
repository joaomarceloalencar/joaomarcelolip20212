package main
 
import "fmt"

func main() {
  const congelamentoC, ebulicaoC = 0.0, 100.0
  fmt.Printf("%g C = %g F\n", congelamentoC, cParaF(congelamentoC))
  fmt.Printf("%g C = %g F\n", ebulicaoC, cParaF(ebulicaoC))
}

func cParaF(c float64) float64 {
  return (c * 9)/5 + 32
}