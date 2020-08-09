package goarea

import "math"

// Pi é uma porporção númerica definida pela relação entre
// o perimetro de uma circunferÊncia e seu diâmetro
const Pi = math.Pi

// Circ é o reponsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é o responsável por calcular a área de um retângolo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
