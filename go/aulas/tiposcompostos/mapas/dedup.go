// Permite simular um conjunto
package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {
	// Se a string já tiver sido lida, o mapa vai apontar para true
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	// Padrão comum em Go, inicializa uma variável antes da condição do if
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}