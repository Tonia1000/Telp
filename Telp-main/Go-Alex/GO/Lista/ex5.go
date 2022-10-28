// *ALEXANDRE* 5 - Criar uma função que faça média entre notas
package main

import "fmt"

func main() {
	media()
}

func media() {
	n := 0
	m := 0
	nt := 0
	sn := 0
	fmt.Printf("Digite a quantidade de notas: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Digite uma nota:")
		fmt.Scan(&nt)
		sn = sn + nt

		m = sn / n
	}

	fmt.Printf("A media das %d notas é igual a %d.\n", n, m)
}
