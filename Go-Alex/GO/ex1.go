// *GABRIEL* 1 - criar um laço que crie com asterisco uma arvore de natal
package main

import (
	"fmt"
)

func main() {
	conversao()
}

func conversao() {
	min := 0

	fmt.Printf("Digite a quantidade de minutos que deseja converter: ")
	fmt.Scan(&min)

	conv := min * 60

	fmt.Printf("Ao converter %d minutos, temos %d segundos.", min, conv)
}