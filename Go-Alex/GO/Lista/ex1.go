// *GABRIEL* 1 - criar um laço que crie com asterisco uma arvore de natal
package main

import (
	"fmt"
)

func main() {
	var matriz [10][7]string
	var f, c int
	f = 4
	c = 2
	i := 1
	count := 0

	matriz[0][3] = "*"
	for i < 15 {
		matriz[i][c] = "*"
		matriz[i][f] = "*"
		if f <= 7 && c >= 0 {
			matriz[i][c] = "*"
			matriz[i][f] = "*"
			if c != 0 {
				c--
			}
			f++
		}
		for f == 7 && c == 0 {
			for j := 0; j < 5; j++ {
				matriz[i][j] = "*"
				matriz[i][j+2] = ""
			}
			f = 4
			c = 2
			count++
		}
		i++
		if count == 2 {
			for i <= 9 {
				matriz[i][2] = "|"
				matriz[i][4] = "|"
				i++
			}
			break
		}
	}

	arv(matriz)

}

func arv(matriz [10][7]string) {
	for i := 0; i < 10; i++ {
		fmt.Println()
		for j := 0; j < 7; j++ {
			fmt.Printf("%s ", matriz[i][j])
		}
	}
	fmt.Println()
}
