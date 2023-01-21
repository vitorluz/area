package goarea

import "math"

// Pi é um exemplo que sera usado
const Pi = 3.1416

// Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a area de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível, privado nao precisa de comentario
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
