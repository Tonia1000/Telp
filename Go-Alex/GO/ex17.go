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

	sort.Strings(list)
	fmt.Printf("Em ordem alfabética: \n %s", list)

	
}