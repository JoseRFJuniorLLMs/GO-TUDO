package main

import (
	"fmt"
)

func main() {

	for number := 20; number >= 0; number-- {
		fmt.Println("MAIOR", number)
		number += 1
	}
}
