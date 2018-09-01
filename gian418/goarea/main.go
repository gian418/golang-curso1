package goarea

import (
	"math"
)

// Pi Por ser publico precisa de um comentario
const Pi = 3.1416

// Circ Por ser publico precisa de um comentario
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect Por ser publico precisa de um comentario
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é uma função visivel
func _TriangulaEq(base, altura float64) float64 {
	return (base * altura) / 2
}
