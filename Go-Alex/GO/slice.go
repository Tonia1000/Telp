package main

import "fmt"

func main() {
	test_slice()

}

func test_slice() {
	s := make([]int, 10, 20)
	s1 := make([]int, 5)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s = append(s, 1)

	// provando que a posição 0 de s2 é a msm que a de s
	s[0] = 25

	s2 := append(s, 11)

	copy(s1, s)

	fmt.Println(s, len(s), cap(s))
	fmt.Println("\n\n", s1, len(s1), cap(s1))
	fmt.Println("\n\n", s2, len(s2), cap(s2))

	// provando que estao usando o msm endereço
	fmt.Println("\n\n", &s[0], &s2[0])

}
