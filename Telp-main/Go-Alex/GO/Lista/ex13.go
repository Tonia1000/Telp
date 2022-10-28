// *DEIVISON* 13 - Criar uma função que converta a moeda Real , para moeda dólar
package main

import (
	"fmt"
)

func main() {
	real_dolar()
}

func real_dolar() {
	var r float64 = 0
	var d float64 = 0

	fmt.Printf("Digite o valor a ser convertido: ")
	fmt.Scanf("%f", &r)

	d = r * 5.28

	fmt.Printf("%2f reais equivale a %2f dolares.\n", r, d)
}
