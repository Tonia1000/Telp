// *LUIZ* 11 - Crie uma função recebe uma temperatura em graus celsius e converta para Fahrenheit.
package main

import (
	"fmt"
)

func main() {
	ctemp()
}

func ctemp() {
	var f float64
	var c float64

	fmt.Printf("Digite a temperatura a ser convertida: ")
	fmt.Scanf("%f", &c)

	f = (c * 1.8) + 32

	fmt.Printf("\n%2f em Celsius equivale a %2f Fahrenheit.\n", c, f)
}
