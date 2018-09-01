package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numero inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("Byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	s1 := "Olá, meu nome é Gian"
	fmt.Println("O tamanho da String é", len(s1))

	s2 := `Olá,
	meu
	nome
	é
	Gian`
	fmt.Println("O tamanho da String é", len(s2))

	// Não existe tipo char
}
