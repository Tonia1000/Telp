// *JOAO* 15 - Faça um programa que mostra na tela os número
// de 1 a 100 e print em uma matriz quais deles sao numero de Ouro ?
package main

import (
	"fmt"
)

func main() {
	ouro()
}

func ouro() {
	var matriz [10][9]int
	n := 1
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			matriz[i][j] = n
			n++
		}
	}

	fmt.Printf("Em uma matriz de 1 a 100: \n")
	for i := 0; i < 10; i++ {
		fmt.Println(matriz[i])
	}

	ou := []int{1, 1}

	for i := 0; ou[i] < 50; i++ {
		ou = append(ou, (ou[i] + ou[i+1]))
	}

	fmt.Printf("\nOs números de ouro no intervalo de 1 a 100 são: %d\n", ou)

}
