package main

import "fmt"

func main() {
	soma()
}

func soma() {
	soma := 0
	for i := 0; i < 20; i++ {
		numero := 0
		fmt.Printf("digite um numero:")
		fmt.Scan(&numero)
		soma = numero + soma

	}

	fmt.Println(soma)
}
