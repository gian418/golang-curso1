package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{45, 23, 7, 89, 2}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
