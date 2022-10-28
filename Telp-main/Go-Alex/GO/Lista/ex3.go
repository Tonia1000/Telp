// *VINICIUS* 3 - Criar uma função que converta anos para minutos e segundos.
package main

import "fmt"

func main() {
	conversao()
}

func conversao() {
	an := 0
	fmt.Printf("Digite a quantidade de anos que deseja converter: ")
	fmt.Scan(&an)

	min := an * 525960
	seg := min * 60

	fmt.Printf("Ao converter %d anos, temos %d minutos ou temos %d segundos.\n", seg, an, min)
}
