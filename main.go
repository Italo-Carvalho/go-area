package goarea

import "math"

/*
   Pi é uma proporção numérica definida pela relação entra
   o perímetro de uma circuferência e seu diâmentro
*/
const Pi = 3.1416

// Circ é reponsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responvel por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
