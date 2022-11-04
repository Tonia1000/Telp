// crie um slice e imprima o primeiro endereço de memoria na tela e seus conteudos em ordem decrescente

package main

import (
	"fmt"
	"sort"
)

func main() {
	mem()

}

func mem() {
	s := make([]int, 10, 20)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(s, len(s), cap(s))

	sort.double(list)
	// imprimir a posição na memoria
	fmt.Println("\n\n", &s[0], &s2[0])

}