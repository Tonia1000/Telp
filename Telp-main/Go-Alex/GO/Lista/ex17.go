// *KARLA* 17 - Escrever uma lista de nomes que retorne em ordem alfabética
package main

import (
	"fmt"
	"sort"
)

func main() {
	alf()
}

func alf() {
	list := []string{"José", "Beatriz", "Vivian", "Maria", "Nicolas", "Adalberto"}

	// for i:=0; i<5; i++{
	// 	p := ""
	// 	fmt.Printf("\n- ")
	// 	fmt.Scan("%s", p)

	// 	list[i] = p
	// }

	fmt.Printf("Lista de nomes original: \n%s\n", list)

	sort.Strings(list)
	fmt.Printf("Em ordem alfabética: \n%s\n", list)

}
