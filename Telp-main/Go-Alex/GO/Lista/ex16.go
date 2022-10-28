// *PROF ALEX* 16 - Escrever algoritmo em Go que receba a data de nascimento do usuário como entrada,
// e retorne o dia da semana que o usuário nasceu.
package main

import (
	"fmt"
	"time"
)

func main() {
	dia()
}

func dia() {
	var d int
	var m int
	var a int

	fmt.Println("Digite sua data de nascimento:")
	fmt.Println("\nDia: ")
	fmt.Scan(&d)
	fmt.Println("\nMes: ")
	fmt.Scan(&m)
	fmt.Println("\nAno: ")
	fmt.Scan(&a)

	// d = 27
	// m = 8
	// a = 1997

	dn := time.Date(a, time.Month(m), d, 00, 00, 00, 00, time.Local)
	fmt.Println("\n\nO dia da semana que voce nasceu foi ")

	if dn.Weekday().String() == "Monday" {
		fmt.Println("Segunda-feira")
	} else if dn.Weekday().String() == "Tuesday" {
		fmt.Println("Terça-feira")
	} else if dn.Weekday().String() == "Wednesday" {
		fmt.Println("Quarta-feira")
	} else if dn.Weekday().String() == "Thursday" {
		fmt.Println("Quinta-feira")
	} else if dn.Weekday().String() == "Friday" {
		fmt.Println("Sexta-feira")
	} else if dn.Weekday().String() == "Saturday" {
		fmt.Println("Sábado")
	} else if dn.Weekday().String() == "Sunday" {
		fmt.Printf("Domingo")
	}

}
