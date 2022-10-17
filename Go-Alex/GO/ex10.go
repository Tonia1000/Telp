// *CAROLINE* 10 - uma função para dizer se o numero inserido é primo ou n
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
