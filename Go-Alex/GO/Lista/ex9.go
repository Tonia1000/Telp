// *RAILSON* 9 - Receber três números e faça a multiplicação entre eles
package main

import (
	"fmt"
)

func main() {
	mult()
}

func mult() {
	n1 := 0
	n2 := 0
	n3 := 0
	result := 0

	fmt.Printf("Digite um numero: \n")
	fmt.Scan(&n1)
	fmt.Printf("Digite um numero: \n")
	fmt.Scan(&n2)
	fmt.Printf("Digite um numero: \n")
	fmt.Scan(&n3)

	result = n1 * n2 * n3

	fmt.Printf("A multiplicacao entre os numero %d, %d e %d e igual a %d.\n", n1, n2, n3, result)
}
