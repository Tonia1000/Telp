package main

import "fmt"

func main() {
	test_map()

}

func test_map() {
	nota := make(map[string]float64)

	nota["Alex"] = 3.0
	nota["Mauricio"] = 10.0
	nota["Sthefany"] = 10.0

	fmt.Println(nota)

}
