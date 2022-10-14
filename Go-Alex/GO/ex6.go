// *ANTONIA* 6 - Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada
package main

import "fmt"

func main() {
	calc()
}

func media() {
	op := 0
	c := ""
	n1 := 0
	n2 := 0
	result := 0
	fmt.Printf("Digite a operação desejada: \n")
	fmt.Printf(" 1. Somar \n 2. Subtrair \n 3. Dividir \n 4. Multiplicar")
	fmt.Scan(&op)

	fmt.Printf("Digite o primeiro numero: ")
	fmt.Scan(&n1)
	fmt.Printf("Digite o segundo numero: ")
	fmt.Scan(&n2)

	if op == 1{
		result = n1+n2
		c == "Somar"

	}
	if op == 2{
		result = n1- 
		c == "Subtrair"

	}
	if op == 3{
		result == n1/n2
		c == "Dividir"

	}
	if op == 4{
		result == n1*n2
		c == "Multiplicar"
		
	}
			
	fmt.Printf("Ao %d, os numeros %d e %d, temos como resultado %d. \n", op, n1, n2, result)
}

// https://go.dev/play/
