package main

import (
  "flag" // O pacote flag retorna ponteiros para parâmetros
  "fmt"
  "strings"
)

var n = flag.Bool("n", false, "omitir nova linha")
var sep = flag.String("s", " ", "separador")

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }
}