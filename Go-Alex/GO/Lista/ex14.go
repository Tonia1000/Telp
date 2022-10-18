// *GEOVANY* 14 - crie uma função que calcula o IMC
package main

import (
	"fmt"
)

func main() {
	imc()
}

func imc() {
	var a float64 = 0
	var p float64 = 0

	fmt.Printf("Digite sua altura em metros: \n")
	fmt.Scan("%f", &p)
	fmt.Printf("Digite seu peso em kg: \n")
	fmt.Scan("%f", &a)

	imc := p / (a * a)

	fmt.Printf("Seu IMC e %2f.", imc)
}
