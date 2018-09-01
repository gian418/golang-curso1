package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[76123761827] = "Pedro"
	aprovados[61987624876] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])

	delete(aprovados, 12345678978)
	fmt.Println(aprovados)
}
