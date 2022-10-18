// *STELA* 12 - Crie uma função que classifique o nível de satisfação do cliente usando METRICA NPS
package main

import (
	"fmt"
)

// m = metrica; d = detractors; pas = passives; pro = promoters; c = controle do for(laço)
func main() {
	var m, d, pas, pro int
	c := 1
	feedback(m, d, pas, pro, c)
}

// Recolhendo o feedback do serviço prestado
func feedback(m, d, pas, pro, c int) {
	var n int
	for n = 0; c == 1; n++ {
		fmt.Printf("Qual nível você classifica o nosso atendimento. \nConsiderando '0' como ruim e '10' como excelente.")
		fmt.Scan(&m)
		
		if m > 10 {
			fmt.Printf("Valor inválido! Processo encerrado.")
			n--
		}

		fmt.Printf("\n\nDeseja reavaliar? \n1 - Sim \n0 - Não")
		fmt.Scan(&c)
		
		nps(m, d, pas, pro)
	}
}

// após a avaliação é necessario comparar os niveis de satisfação
func nps(m, d, pas, pro int) {
	if m <= 6 {
		fmt.Println("\n\nObrigado pelo seu feedback, espero que o prossimo contato seja agradavel!")
		d++
	} else if m == 7 || m == 8 {
		fmt.Println("\n\nObrigado pelo feedback, sempre busco a satisfação do cliente!")
		pas++
	} else if m == 9 || m == 10 {
		fmt.Println("\n\nObrigado pelo feedback, espero que nos encontremos novamente em breve!")
		pro++
	}
}