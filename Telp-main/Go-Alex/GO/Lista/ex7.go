// *ANA* 7 - Crie uma função que receba 10 números, faça a média e mostre o maior número
package main

import "fmt"

func main() {
	maior()
}

func maior() {
	n := 0
	x := 0

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um numero: ")
		fmt.Scan(&n)
		if n > x {
			x = n
		}

	}

	fmt.Printf("O maior numero digitado foi %d\n", x)
}
