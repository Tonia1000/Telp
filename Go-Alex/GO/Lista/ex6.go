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
	fmt.Printf("Digite a operação desejada: \n")
	fmt.Printf(" 1. Somar \n 2. Subtrair \n 3. Dividir \n 4. Multiplicar\n")
	fmt.Scan(&op)

	fmt.Printf("Digite o primeiro numero: \n")
	fmt.Scan(&n1)
	fmt.Printf("Digite o segundo numero: \n")
	fmt.Scan(&n2)

	if op == 1{
		result = n1+n2

	}else if op == 2{
		result = n1-n2

	}else if op == 3{
		result = n1/n2

	}else if op == 4{
		result = n1*n2
		
	}
			
	fmt.Printf("O resultado da operacao entre os numeros %d e %d e %d. \n", n1, n2, result)
}