package main

import (
	"fmt"
)

func main() {

	soma(3, 6, 9)
	fmt.Println("iniciando soma")
	fmt.Println("RESULTADO", soma)
}

func soma(n1, n2, n3 int) int {

	fmt.Println(soma)

	return n1 + n2 + n3

}
