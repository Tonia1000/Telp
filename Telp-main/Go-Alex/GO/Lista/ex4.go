// *STHEFANY* 4 - crie uma função que some duas notas
package main

import "fmt"

func main() {
	soma()
}

func soma() {
	n1 := 0
	fmt.Printf("Digite ao primeiro valor: ")
	fmt.Scan(&n1)

	n2 := 0
	fmt.Printf("Digite o segundo valor: ")
	fmt.Scan(&n2)

	soma := n1 + n2

	fmt.Printf("A soma de %d e %d é igual a %d.\n", n1, n2, soma)
}
