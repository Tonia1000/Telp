// *ANTONIA* 6 - Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada
package main

import "fmt"

func main() {
	calc()
}

func media() {
	op := ""
	n1 := 0
	n2 := 0
	result := 0
	fmt.Printf("Digite a operação desejada sem acento: ")
	fmt.Scan(&op)

	fmt.Printf("Digite o primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Printf("Digite o segundo numero: ")
	fmt.Scan(&n2)

	if op == "soma" {
		result = n1+n2

	}
	if op == "divisao"{
		result = n1/n2

	}
	if op == "multiplicacao"{
		result == n1*n2

	}

	fmt.Printf("A media das %d notas é igual a %d.\n", n, m)
}

// https://go.dev/play/