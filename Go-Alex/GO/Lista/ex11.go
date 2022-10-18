// *LUIZ* 11 - Crie uma função recebe uma temperatura em graus celsius e converta para Fahrenheit.
package main

import (
	"fmt"
)

func main() {
	ctemp()
}

func ctemp() {
	var f float64 = 0
	var c float64 = 0

	fmt.Printf("Digite a temperatura a ser convertida: ")
	fmt.Scan("%f", &c)

	f = (c * 1.8) + 32

	fmt.Printf("%2f em Celsius equivale a %2f Fahrenheit.", c, f)
}
