// *CAROLINE* 10 - uma função para dizer se o numero inserido é primo ou n
package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Printf("Digite um número qualquer: ")
	fmt.Scan(&n)

	if primo(n, 2) == true {
		fmt.Println("\nEste número é primo")
	} else {
		fmt.Println("\nEste número não é primo")
	}
}

func divisivel(n, j int) bool {
	if (n % j) == 0 {
		return true
	}
	return false
}

func primo(n, i int) bool {
	if n == 1 {
		return false
	}
	if n == i {
		return true
	}
	if divisivel(n, i) {
		return false
	}
	return primo(n, i+1)
}