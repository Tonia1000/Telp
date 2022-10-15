// *SERGIO* 8 - Crie uma função que retira as vogais de uma frase
package main

import "fmt"

func main() {
	vogal()
}

func vogal() {
	p := [""]
	l := ""
	n := 0

	fmt.Printf("Quantas letras tem a palavra q deseja digitar?")
	fmt.Scan(&n)

	fmt.Printf("Digite a palavra letra por letra")
	for i := 0; i<n; i++{
		fmt.Scan(&l)

		if  l != "a"{
			if l != "e"{
				if l != "i"{
					if l != "o"{
						if l != "u"{
							p[i] = l

						}
					}
				}
			}
		}

	}

	fmt.Printf("A palavra digitada sem vogais é %s",)
}