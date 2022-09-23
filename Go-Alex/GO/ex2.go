// *MAURÍCIO* 2 - Criar uma função que converta segundos para minutos e segundos.
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
