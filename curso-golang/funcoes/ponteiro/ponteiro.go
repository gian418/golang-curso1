package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)        // por valor
	fmt.Println(n) // não altera o n deste main

	inc2(&n)       // por referencia
	fmt.Println(n) // altera o n deste main pq passa a referencia do ponteiro dele

}
