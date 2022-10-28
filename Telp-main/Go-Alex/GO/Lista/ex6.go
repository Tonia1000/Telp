// *ANTONIA* 6 - Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada
package main

import "fmt"

func main() {
	calc()
}

func calc() {
	op := 0
	n1 := 0
	n2 := 0
	result := 0
	x := ""
	fmt.Printf("Digite a operação desejada: \n")
	fmt.Printf(" 1. Somar \n 2. Subtrair \n 3. Dividir \n 4. Multiplicar\n")
	fmt.Scan(&op)

	fmt.Printf("Digite o primeiro numero: \n")
	fmt.Scan(&n1)
	fmt.Printf("Digite o segundo numero: \n")
	fmt.Scan(&n2)

	if op == 1 {
		result = n1 + n2
		x = "Somar"

	} else if op == 2 {
		result = n1 - n2
		x = "Subtrair"

	} else if op == 3 {
		result = n1 / n2
		x = "Dividir"

	} else if op == 4 {
		result = n1 * n2
		x = "Multiplicar"

	}

	fmt.Printf("Ao %s os numeros %d, %d temos como resultado %d. \n", x, n1, n2, result)
}
