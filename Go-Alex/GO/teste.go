package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("hello")

	// var nome string = "samsung"
	// var num int = 10
	// var flo float64 = 1.90
	// var bol bool = true

	// fmt.Printf("O nome é %s \n", nome)
	// fmt.Printf("valor inteiro: %d\n", num)
	// fmt.Printf("valor em float: %f\n", flo)
	// fmt.Printf("valor booleano: %t\n", bol)

	// var numero1, numero2 int = 30, 10

	// var soma = numero1 + numero2
	// var subtracao = numero1 - numero2
	// var multiplicacao = numero1 * numero2
	// var divisao = numero1 / numero2

	// fmt.Printf("resultado da soma: %d\n", soma)
	// fmt.Printf("resultado da subtracao: %d\n", subtracao)
	// fmt.Printf("resultado da multiplicacao: %d\n", multiplicacao)
	// fmt.Printf("resultado da divisao %d\n", divisao)

	// var numerodig, numero int
	// fmt.Scan(&numerodig, &numero)

	// fmt.Print(numerodig)
	// fmt.Print(numero)

	// numerotype := 10

	// fmt.Printf("%d",numerotype)

	mm()
}

// fazer a soma de 20 numeros do tipo int, e imprimeir o resultado
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

// Fazer a entrada de 15 numeros e devolver quais deles é o maior e qual é o menor
func mm() {

	ma := 0
	me := math.MaxInt
	for j := 0; j < 15; j++ {
		num := 0
		fmt.Printf("Digite um numero: \n")
		fmt.Scan(&num)
		if num > ma {
			ma = num
		}
		if num < me {
			me = num
		}
	}
	fmt.Printf("O maior numero é: %d\n O menor numero é: %d\n", ma, me)

}
