package main

import (
	"fmt"
)

func main() {
	numeros := [6]int{0, 1, 2, 3, 4}
	numeros[5] = 5
	fmt.Println(numeros)
	fmt.Println("TAMANHO DO ARRAY")
	fmt.Println(len(numeros))
	cliente := map[string]string{
		"nome": "JUNIOR",
	}
	fmt.Println(cliente["nome"])
}
