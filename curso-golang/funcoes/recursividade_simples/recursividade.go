package main

import (
	"fmt"
)

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)

	//Não deixa passar valor negativo por ser uint
	//resultado2 := fatorial(-4)
	//fmt.Println(resultado2)
}
