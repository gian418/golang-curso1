package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	switch { // switch true
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(2.1))
}
