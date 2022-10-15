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

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func filter(data []string, f func(string) bool) []string {

// 	fltd := make([]string, 0)

// 	for _, e := range data {

// 		if f(e) {
// 			fltd = append(fltd, e)
// 		}
// 	}

// 	return fltd
// }

// func main() {

// 	words := []string{"war", "water", "cup", "tree", "storm"}

// 	p := "w"

// 	res := filter(words, func(s string) bool {

// 		return strings.HasPrefix(s, p)
// 	})

// 	fmt.Println(res)

// }
