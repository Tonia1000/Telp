// *GEOVANY* 14 - crie uma função que calcula o IMC
package main

import (
	"fmt"
)

func main() {
	imc()
}

func imc() {
	var a float64
	var p float64

	fmt.Printf("Digite sua altura em metros: \n")
	fmt.Scanf("%f", &a)
	fmt.Printf("Digite seu peso em kg: \n")
	fmt.Scanf("%f", &p)

	imc := p / (a * a)

	fmt.Printf("Seu IMC e %2f.\n", imc)
}
